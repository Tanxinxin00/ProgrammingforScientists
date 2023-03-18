package main

import (
	"fmt"
	"math"
)

// BarnesHut is our highest level function.
// Input: initial Universe object, a number of generations, and a time interval.
// Output: collection of Universe objects corresponding to updating the system
// over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	timePoints[0] = initialUniverse
	// update each universe given the former universe
	for i := 1; i < numGens+1; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time, theta)
		// This is used to track the progress of update.
		if i%2000 == 0 {
			fmt.Printf("UpdatingUniverse%d", i)
		}
	}

	return timePoints
}

// Update universe takes the pointer of the current universe to update the new one and return the pointer of the new universe.
func UpdateUniverse(currentUniverse *Universe, time, theta float64) *Universe {
	// Use the current universe to generate a quad tree
	CurrentQuadTree := MakeQuadTree(currentUniverse)

	// create a new universe
	var newUniverse Universe
	newUniverse.width = currentUniverse.width
	// create the same number of stars
	numStars := len(currentUniverse.stars)
	newUniverse.stars = make([]*Star, numStars)

	for i := range currentUniverse.stars {
		var newstar Star
		newUniverse.stars[i] = &newstar
		// if the current star is in the universe, update the curent star;else just ignore it.
		if InUniverse(currentUniverse.stars[i], currentUniverse.width) {
			newUniverse.stars[i].mass = currentUniverse.stars[i].mass
			newUniverse.stars[i].radius = currentUniverse.stars[i].radius
			newUniverse.stars[i].red = currentUniverse.stars[i].red
			newUniverse.stars[i].blue = currentUniverse.stars[i].blue
			newUniverse.stars[i].green = currentUniverse.stars[i].green
			newUniverse.stars[i].acceleration = UpdateAcceleration(currentUniverse.stars[i], CurrentQuadTree.root, theta)
			newUniverse.stars[i].position = UpdatePosition(currentUniverse.stars[i], newUniverse.stars[i], time)
			newUniverse.stars[i].velocity = UpdateVelocity(currentUniverse.stars[i], newUniverse.stars[i], time)
			fmt.Printf("star%d", i)
			fmt.Println(newUniverse.stars[i].acceleration, newUniverse.stars[i].position, newUniverse.stars[i].velocity)
		}
	}
	return &newUniverse
}

// function MakeQuadTree generates a quadtree according to the current positions of all the stars in the universe.
func MakeQuadTree(currentUniverse *Universe) *QuadTree {
	var tree QuadTree
	var rootnode Node
	var DummyRootStar Star

	rootnode.star = &DummyRootStar
	tree.root = &rootnode

	// the ultimate root node has the whole universe as its sector/quadrant
	rootnode.sector.x = 0.0
	rootnode.sector.y = currentUniverse.width
	rootnode.sector.width = currentUniverse.width

	// range through all the stars and put the stars in the quadtree.
	for i := range currentUniverse.stars {
		// if the current star is in the universe, put the star in the quadtree, else just ignore it.
		if InUniverse(currentUniverse.stars[i], currentUniverse.width) {
			tree.root.PutTheStarInTheQuadrant(currentUniverse.stars[i])
		}
	}

	// When the quadtree is made, compute mass and center of gravity for all the dummy stars in the quadtree.
	tree.root.ComputeMassAndPositionForDummyStars()

	return &tree
}

// PutTheStarInTheQuadrant is a recursive function that takes a star and puts it in the correct position in the sector of the current node.
func (root *Node) PutTheStarInTheQuadrant(star *Star) {
	// If the children of the node has not been created, create its four children nodes.
	if root.children == nil {
		root.DivideCurrentSector()
	}

	// Find the subquadrant of the current node where the star belongs to
	// loci is an interger corresponding to the index of the children nodes of the rootnode.
	loci := LocateToSub(star, root.sector)

	// If the subquadrant node has not been assigend to a star yet, assign the current star to the node
	if root.children[loci].star == nil {
		root.children[loci].star = star

		// if the subquadrant node has been occupied and the radius of the star is 0, i.e., it is occupied by a dummy star
	} else if root.children[loci].star.radius == 0 {
		// go down a level, put the star in the sector of this children node.
		root.children[loci].PutTheStarInTheQuadrant(star)

		// if the subquadrant node has been occupied by an actual star, create subquadrants of the current node
		// and put the current star and the occupying star into the new subquadrants.
		// Also, reassign a dummy star to the current subquadrant node.
	} else {
		root.children[loci].PutTheStarInTheQuadrant(star)
		root.children[loci].PutTheStarInTheQuadrant(root.children[loci].star)
		var dummystar Star
		dummy := &dummystar
		root.children[loci].star = dummy
	}

}

// DivideCurrentSector divides the current sector of the node into four parts and creates four children nodes representing each of them.
func (root *Node) DivideCurrentSector() {
	root.children = make([]*Node, 4)

	for i := range root.children {
		var node Node
		root.children[i] = &node
		// the sector of the children nodes is a half of the sector of the parent node.
		root.children[i].sector.width = root.sector.width * 0.5
	}

	//define the sector of the NorthWest child
	root.children[0].sector.x = root.sector.x
	root.children[0].sector.y = root.sector.y - root.children[1].sector.width

	//define the sector of the NorthEast child
	root.children[1].sector.x = root.sector.x + root.children[1].sector.width
	root.children[1].sector.y = root.sector.y - root.children[1].sector.width

	//define the sector of the SouthWest child
	root.children[2].sector.x = root.sector.x
	root.children[2].sector.y = root.sector.y

	//define the sector of the SouthEast child
	root.children[3].sector.x = root.sector.x + root.children[3].sector.width
	root.children[3].sector.y = root.sector.y

}

// function LocateToSub takes the current star and a sector as input
// and determines which subquadrant of the sector the star should be in.
// with '0' being the northwest subquadrant;'1' being the northeast;'2' being the southwest;'3' being the southeast
func LocateToSub(star *Star, sec Quadrant) int {
	var loci int
	// initialize the loci as 99 to prevent the loci being wrong(because the default is 0,which is the same as the northwest subquadrant)
	loci = 99

	// if this is true, it is on the east side of the sector
	if star.position.x >= sec.x+sec.width*0.5 {
		// if this is true, it is on the south side of the sector
		if star.position.y >= sec.y-sec.width*0.5 {
			// '3' indicates southeast
			loci = 3
			// if not, then it is on the north side of the sector
		} else {
			// '1' indicates northeast
			loci = 1
		}
		// it is on the west side of the sector
	} else {
		// it is on the south side od the sector
		if star.position.y >= sec.y-sec.width*0.5 {
			// '2' indicates southwest
			loci = 2
			// it is on the north side of the sector
		} else {
			// '0' indicates northwest
			loci = 0
		}
	}

	if loci == 99 {
		panic("Wrong sector! Can't put the current star in this sector")

	}

	return loci
}

// ComputeMassAndPosition is a recursive method that computes the mass and center of gravity for all the dummy stars in the quadtree.
func (root *Node) ComputeMassAndPositionForDummyStars() {
	// for this current node, range through all of its children
	for i := range root.children {
		// if there is no star assigned to the current children node, just ignore it.
		if root.children[i].star != nil {
			// radius is not zero means it is a real star
			if root.children[i].star.radius != 0.0 {
				root.star.mass += root.children[i].star.mass
				root.star.position.x += root.children[i].star.position.x * root.children[i].star.mass
				root.star.position.y += root.children[i].star.position.y * root.children[i].star.mass
				// radius is zero means it is a dummy star.
			} else {
				// if it is a dummy star, compute the mass and position for all its descendents before computing itself.
				root.children[i].ComputeMassAndPositionForDummyStars()
				root.star.mass += root.children[i].star.mass
				root.star.position.x += root.children[i].star.position.x * root.children[i].star.mass
				root.star.position.y += root.children[i].star.position.y * root.children[i].star.mass
			}
		}
	}

	// The above position is the weighted position of a node's descendents combined, divide it by its mass and get the center of gravity.
	root.star.position.x = root.star.position.x / root.star.mass
	root.star.position.y = root.star.position.y / root.star.mass

}

// InUniverse determines if the current star is in the universe or not.
func InUniverse(star *Star, width float64) bool {
	inUniverse := false

	if star.position.x >= 0 && star.position.x <= width {
		if star.position.y >= 0 && star.position.y <= width {
			inUniverse = true
		}

	}

	return inUniverse
}

/*
func InBlackHole(star *Star) bool{

}
*/

// Updateacceleration takes a rootnode(responding to the quadtree) and computes the quadtree's force acting on the star
// and returns the accleration
func UpdateAcceleration(star *Star, root *Node, theta float64) OrderedPair {
	var accel OrderedPair

	// compute net force vector acting on the star
	force := ComputeHeuristicForce(root, star, theta)

	if star.mass == 0 {
		panic("invalid star with 0 mass")
	}

	// now, calculate acceleration (F = ma)
	accel.x = force.x / star.mass
	accel.y = force.y / star.mass

	return accel
}

// ComputeHeuristicForce computes the heuristic force the current node acts on the current star
// theta is a threshold that determines whether to go deeper to compute the more exact force or not.
func ComputeHeuristicForce(root *Node, star *Star, theta float64) OrderedPair {
	var totalforce OrderedPair
	d := DistanceBetweenStars(root.star, star)
	// stat is the measurement used to compare with theta.
	stat := root.sector.width / d

	// if the current node does not have children, i.e., it is a leaf node
	if root.children == nil {
		// if the node has been assigned to a star, directly compute the force that this star acts on our object star.
		if root.star != nil {
			force := ComputeForce(star, root.star, d)
			totalforce.x += force.x
			totalforce.y += force.y
		}
		// if the current node is not a leaf node,i.e., it is a dummy star and the stat is over the theshold,
		// then compute the heuristic force of all its children nodes acting on our object star combined.
	} else if stat > theta {
		// range through all its four children
		for i := range root.children {
			// if the children node is empty, just ignore
			if root.children[i].star != nil {
				force := ComputeHeuristicForce(root.children[i], star, theta)
				totalforce.x += force.x
				totalforce.y += force.y
			}
		}
		// if the current node is a dummy star, and the stat is below the theshold,
		// just compute the force the dummy star acts on our object star.
	} else {
		force := ComputeForce(star, root.star, d)
		totalforce.x += force.x
		totalforce.y += force.y
	}

	return totalforce
}

// ComputeForce
// Input: Two star objects s1 and s2.
// Output: The force due to gravity s2 acts on s1
func ComputeForce(s1, s2 *Star, d float64) OrderedPair {
	var force OrderedPair

	// if it is the same star or dummy star with the same position ,don't compute the force.
	if d == 0 {
		force.x = 0.0
		force.y = 0.0
	} else {
		F := G * s1.mass * s2.mass / (d * d) // magnitude of gravity
		deltaX := s2.position.x - s1.position.x
		deltaY := s2.position.y - s1.position.y

		force.x = F * deltaX / d // deltaX/dist = cos theta
		force.y = F * deltaY / d // deltaY/dist = sin theta
	}

	return force
}

// computes the distance between two stars.
func DistanceBetweenStars(s1, s2 *Star) float64 {
	d2 := math.Pow(s1.position.x-s2.position.x, 2) + math.Pow(s1.position.y-s2.position.y, 2)
	return (math.Sqrt(d2))
}

// Updatevelocity takes the old(current) star, the new star, the timestep as input
// to update the new star's velocity.
func UpdateVelocity(olds, s *Star, time float64) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = olds.velocity.x + s.acceleration.x*time
	vel.y = olds.velocity.y + s.acceleration.y*time

	return vel
}

// UpdatePosition takes the old and the new star, the timestep as input
// and gives us the updated position of the new star.
func UpdatePosition(olds, s *Star, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*s.acceleration.x*time*time + olds.velocity.x*time + olds.position.x
	pos.y = 0.5*s.acceleration.y*time*time + olds.velocity.y*time + olds.position.y

	return pos
}
