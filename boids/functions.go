package main

import (
	"math"
	"math/rand"
)

// SimulateBoids simulates the motion pattern of a group of flying birds over a series of time snapshots.
// It takes an initalSky, the number of generations that we update the sky and the timestep between each update as input.
// And outputs a slice of pointer of Sky objects representing the situation of birds on each timepoint.
func SimulateBoids(initialSky *Sky, numGens int, timeStep float64) []*Sky {
	timePoints := make([]*Sky, numGens+1)
	timePoints[0] = initialSky

	//range over number of generations and update each sky with the last sky.
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], timeStep)
	}

	return timePoints
}

// UpdateSky takes the pointer of the previous sky, update the next sky based on this sky and outputs the pointer of the new sky.
func UpdateSky(currentSky *Sky, timeStep float64) *Sky {
	var newSky Sky
	newSky.width = currentSky.width
	newSky.proximity = currentSky.proximity
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor
	newSky.separationFactor = currentSky.separationFactor

	//range over each boid and update their acceleration, velocity and position on the next timepoint.
	newSky.boids = make([]Boid, len(currentSky.boids))
	for i := range currentSky.boids {
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, currentSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(currentSky.boids[i], newSky.boids[i], timeStep, currentSky.maxBoidSpeed)
		newSky.boids[i].position = UpdatePosition(currentSky.boids[i], newSky.boids[i], timeStep, currentSky.width)
	}
	return &newSky
}

//This is the original copySky function that I wrote before learning about pointers.
//I just keep it in case I need it somehow.
/*
	func copySky(currentSky Sky) Sky {
		var newSky Sky
		newSky.width = currentSky.width
		newSky.proximity = currentSky.proximity
		newSky.maxBoidSpeed = currentSky.maxBoidSpeed
		newSky.alignmentFactor = currentSky.alignmentFactor
		newSky.cohesionFactor = currentSky.cohesionFactor
		newSky.separationFactor = currentSky.separationFactor

		numBoids := len(currentSky.boids)
		newSky.boids = make([]Boid, numBoids)
		for i := range currentSky.boids {
			newSky.boids[i].acceleration = currentSky.boids[i].acceleration
			newSky.boids[i].position = currentSky.boids[i].position
			newSky.boids[i].velocity = currentSky.boids[i].velocity
		}

		return newSky
	}
*/

// UpdateAcceleration takes teh pointer of the currentsky and the current boid we are updating as input.
// Decide which boids are close enough to the current boid and calculate their forces on the current boid to get the acceleration.
func UpdateAcceleration(currentSky *Sky, b Boid) OrderedPair {
	neighbors := FindNearbyBoids(currentSky, b)
	force := ComputeForce(neighbors, b, currentSky.separationFactor, currentSky.alignmentFactor, currentSky.cohesionFactor)

	var accel OrderedPair

	numNeis := len(neighbors)
	//If a boid has no neighbor, then no force is acting on it, the acceleration should be zero.
	if numNeis != 0 {
		accel.x = force.x / float64(numNeis)
		accel.y = force.y / float64(numNeis)
	}

	return accel
}

// FindNearbyBoids takes the current sky and the current boid as input and
// gives us a slice of the neighbor boids of the current boid.
func FindNearbyBoids(currentSky *Sky, b Boid) []Boid {
	var neighbors []Boid

	//range over all the boids in the sky and calculate the distances to the current boid.
	for i := range currentSky.boids {
		dist := DistBetweenBoids(currentSky.boids[i], b)

		//If the distance of a boid to the current boid is in the force-acting range, then append it to the neighbor list.
		//Since the current boid itself is part of the iteration, we need to remove it.(its distance to itself is zero)
		if dist <= currentSky.proximity && dist > 0 {
			neighbors = append(neighbors, currentSky.boids[i])
		}
	}

	return neighbors
}

// DistBetweenBoids takes two boids as input and gives us the distance between them.
func DistBetweenBoids(a, b Boid) float64 {
	d2 := (a.position.x-b.position.x)*(a.position.x-b.position.x) + (a.position.y-b.position.y)*(a.position.y-b.position.y)
	d := math.Sqrt(d2)

	return d
}

// ComputeForce computes the separation, alignment and cohesion force acted on the current boid by each of its neighbors and add them together to acquire the total force.
func ComputeForce(neighbors []Boid, b Boid, Sepfactor, Alnfactor, Cohfactor float64) OrderedPair {
	var forces OrderedPair

	// range over all of the current boid's neighbors and compute the three forces respectively.
	for i := range neighbors {
		d := DistBetweenBoids(neighbors[i], b)

		sepforce := SeparationForce(neighbors[i], b, d, Sepfactor)
		alnforce := AlignmentForce(neighbors[i], b, d, Alnfactor)
		cohforce := CohesionForce(neighbors[i], b, d, Cohfactor)

		forces.x += sepforce.x + alnforce.x + cohforce.x
		forces.y += sepforce.y + alnforce.y + cohforce.y
	}

	return forces

}

// The separationforce takes two boids, the distance between them and teh separation factor as input
// and gives us the separationforce according to the formula. The same applies to the other two forces.
func SeparationForce(neighbor, b Boid, dist, Sepfactor float64) OrderedPair {
	var SepForce OrderedPair

	SepForce.x = Sepfactor * (b.position.x - neighbor.position.x) / dist / dist
	SepForce.y = Sepfactor * (b.position.y - neighbor.position.y) / dist / dist

	return SepForce
}

func AlignmentForce(neighbor, b Boid, dist, Alnfactor float64) OrderedPair {
	var AlnForce OrderedPair

	AlnForce.x = Alnfactor * neighbor.velocity.x / dist
	AlnForce.y = Alnfactor * neighbor.velocity.y / dist

	return AlnForce
}

func CohesionForce(neighbor, b Boid, dist, Cohfactor float64) OrderedPair {
	var CohForce OrderedPair

	CohForce.x = Cohfactor * (neighbor.position.x - b.position.x) / dist
	CohForce.y = Cohfactor * (neighbor.position.y - b.position.y) / dist

	return CohForce
}

// Updatevelocity takes the old(current) boid, the new boid, the timestep and the maximum speed as input
// to update the new boid's velocity.
func UpdateVelocity(oldb, b Boid, time, maxSpeed float64) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = oldb.velocity.x + b.acceleration.x*time
	vel.y = oldb.velocity.y + b.acceleration.y*time

	//check if the speed of the boid exceeds the maximum speed for a boid.
	absoluteVel := math.Sqrt(vel.x*vel.x + vel.y*vel.y)
	// if the speed exceeds maximum, then keep the current direction of the speed but limit the absolute speed to maximum speed.
	if absoluteVel > maxSpeed {
		vel.x = maxSpeed * vel.x / absoluteVel
		vel.y = maxSpeed * vel.y / absoluteVel
	}

	return vel
}

// UpdatePosition takes the old and the new boid, the timestep and the width of the sky as input
// and gives us the updated position of the new boid.
func UpdatePosition(oldb, b Boid, time, width float64) OrderedPair {
	var pos OrderedPair
	pos.x = 0.5*b.acceleration.x*time*time + oldb.velocity.x*time + oldb.position.x
	pos.y = 0.5*b.acceleration.y*time*time + oldb.velocity.y*time + oldb.position.y

	//If the boids run out of the picture, it will enter on the other side(right out left in; bottom out top in)
	if pos.x > width {
		pos.x = math.Mod(pos.x, width)
	} else if pos.x < 0 {
		pos.x = width + math.Mod(pos.x, width)
	}

	if pos.y > width {
		pos.y = math.Mod(pos.y, width)
	} else if pos.y < 0 {
		pos.y = width + math.Mod(pos.y, width)
	}

	return pos
}

// InitiateSky takes the number of boids, initialSpeed of the boids, the other parameters for creating the sky as input
// and initiates the initialsky and returns the pointer to this sky.
func InitiateSky(numBoids int, initialSpeed, skyWidth, maxSpeed, proxi, separation, alignment, cohesion float64) *Sky {
	var InitialSky Sky
	InitialSky.width = skyWidth
	InitialSky.boids = make([]Boid, numBoids)
	InitialSky.cohesionFactor = cohesion
	InitialSky.maxBoidSpeed = maxSpeed
	InitialSky.proximity = proxi
	InitialSky.alignmentFactor = alignment
	InitialSky.separationFactor = separation

	//Check if the given initialspeed exceeds the maximum
	if initialSpeed > maxSpeed {
		initialSpeed = maxSpeed
	}

	//For each of the boids, randomly generate a position and a direction of its initialspeed.
	for i := range InitialSky.boids {
		//the x and y position of the boids range from [0.0,skywidth]
		InitialSky.boids[i].position.x = rand.Float64() * skyWidth
		InitialSky.boids[i].position.y = rand.Float64() * skyWidth
		InitialSky.boids[i].velocity.x = randomDirection().x * initialSpeed
		InitialSky.boids[i].velocity.y = randomDirection().y * initialSpeed
		InitialSky.boids[i].acceleration.x = 0.0
		InitialSky.boids[i].acceleration.y = 0.0
	}

	return &InitialSky
}

// randomDirection generates a random direction with unit length.
func randomDirection() OrderedPair {
	var di OrderedPair
	di.x = rand.Float64() - 0.5
	di.y = rand.Float64() - 0.5
	d := math.Sqrt(di.x*di.x + di.y*di.y)
	di.x = di.x / d
	di.y = di.y / d

	return di
}
