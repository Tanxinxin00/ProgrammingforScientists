package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//os.Args[1] is the number of boids to put in the initial sky.
	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if numBoids < 0 {
		panic("Negative number of boids given.")
	}

	//os.Args[2] is the width of the sky.
	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic(err2)
	}

	//os.Args[3] is the initial speed of the boids.
	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic(err3)
	}

	//os.Args[4] is the maximumspeed of the boids.
	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic(err4)
	}

	//os.Args[5] is the number of generations that we want to simulate the sky.
	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic(err5)
	}
	if numGens < 0 {
		panic("Negative number of generations given.")
	}

	//os.Args[6] is how close the boids need to be to interact with each other.
	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic(err6)
	}

	//os.Args[7] is the separation factor used to calculate separation force between the boids.
	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic(err7)
	}

	//os.Args[8] is the alignment factor used to calculate alignment force between the boids.
	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic(err8)
	}

	//os.Args[9] is the cohesion factor used to calculate cohesion force between the boids.
	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9 != nil {
		panic(err9)
	}

	//os.Args[10] is the width of the canvas.
	canvasWidth, err10 := strconv.Atoi(os.Args[10])
	if err10 != nil {
		panic(err10)
	}

	//os.Args[11] is how frequently we want to draw the images.
	imageFrequency, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil {
		panic(err11)
	}

	//os.Args[12] is the timestep between timepoints.
	timeStep, err12 := strconv.ParseFloat(os.Args[12], 64)
	if err12 != nil {
		panic(err12)
	}

	fmt.Println("Command line arguments read successfully.")

	fmt.Println("Simulating system.")

	//using a seed corresponding to the time ensures the random number is different each time.
	rand.Seed(time.Now().Unix())

	//Use the parameters given to generate the first sky.
	InitialSky := InitiateSky(numBoids, initialSpeed, skyWidth, maxBoidSpeed, proximity, separationFactor, alignmentFactor, cohesionFactor)

	timePoints := SimulateBoids(InitialSky, numGens, timeStep)

	fmt.Println("Boids has been simulated!")
	fmt.Println("Ready to draw images.")

	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	//These are for printing the parameters used to generate the filename of the gifs drawn.
	coh := fmt.Sprintf("%.3f", cohesionFactor)
	sep := fmt.Sprintf("%.2f", separationFactor)
	timestep := fmt.Sprintf("%.2f", timeStep)

	filename := fmt.Sprintf("%dproxi_%stime_%scoh_%ssep", int(proximity), timestep, coh, sep)

	gifhelper.ImagesToGIF(images, filename)

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}
