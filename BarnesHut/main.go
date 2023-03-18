package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"os"
	"time"
)

func main() {
	//read from command line which flag we would use to run this program.
	flag := os.Args[1]

	//this is a timer
	start := time.Now()

	var initialUniverse *Universe
	var canvasWidth int
	var frequency int
	var scalingFactor float64
	var numGens int
	var timestep float64

	rand.Seed(time.Now().UnixNano())

	//simulate a single galaxy
	if flag == "galaxy" {
		fmt.Println("Simulating a galaxy.")
		g0 := InitializeGalaxy(500, 4e21, 5e22, 5e22)
		galaxies := []Galaxy{g0}
		width := 1.0e23
		initialUniverse = InitializeUniverse(galaxies, width)
		numGens = 50000
		timestep = 2e14
		canvasWidth = 1000
		frequency = 1000
		scalingFactor = 1e11

		//simulate two galaxies collide
	} else if flag == "collision" {
		fmt.Println("Simulating two galaxies collide.")

		g0 := InitializeGalaxy(200, 4e21, 6e22, 4e22)
		g1 := InitializeGalaxy(200, 4e21, 4e22, 6e22)
		/*
			var direction OrderedPair
			d := DistanceBetweenStars(g1[len(g1)-1], g0[len(g0)-1])
			direction.x = (g1[len(g1)-1].position.x - g0[len(g0)-1].position.x) / d
			direction.y = (g1[len(g1)-1].position.y - g0[len(g0)-1].position.y) / d

			pushspeed := 150.00
			// for eaver star in the galaxies, add an initial speed to push them towards each other.
			for i := range g0 {
				g0[i].velocity.x += pushspeed * direction.x
				g0[i].velocity.y += pushspeed * direction.y

			}

			for i := range g1 {
				g1[i].velocity.x -= pushspeed * direction.x
				g1[i].velocity.y -= pushspeed * direction.y

			}
		*/
		for i := range g0 {
			g0[i].velocity.x -= 150
			g0[i].velocity.y += 100

		}

		for i := range g1 {
			g1[i].velocity.x += 150
			g1[i].velocity.y -= 100

		}

		galaxies := []Galaxy{g0, g1}
		width := 1.0e23

		initialUniverse = InitializeUniverse(galaxies, width)
		numGens = 300000
		timestep = 2e14
		canvasWidth = 1000
		frequency = 1000
		scalingFactor = 1e11

	} else if flag == "jupiter" {
		initialUniverse = InitializeJupiter()
		canvasWidth = 2000
		frequency = 100
		scalingFactor = 5
		numGens = 20000
		timestep = 10

	} else {
		panic("Wrong flag")
	}

	theta := 0.5

	timePoints := BarnesHut(initialUniverse, numGens, timestep, theta)

	fmt.Println("Simulation run. Now drawing images.")
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	filename := fmt.Sprintf("%s_%dGens_%dtstep_%d_%d", flag, numGens, int(timestep), time.Now().Hour(), time.Now().Minute())
	gifhelper.ImagesToGIF(imageList, filename)
	fmt.Println("GIF drawn.")

	elapsed := time.Since(start)
	fmt.Printf("It took %s", elapsed)
}
