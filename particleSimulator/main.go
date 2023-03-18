package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	fmt.Println("Particle simulator.")

	fmt.Println("Generating random particles and initializing board.")

	numParticles := 100
	boardWidth := 1000.0
	boardHeight := 1000.0
	particleRadius := 5.0
	diffusionRate := 5.0

	//assumption: all particles are white

	random := true // make true if we want to scatter across board

	initialBoard := InitializeBoard(boardWidth, boardHeight, numParticles, particleRadius, diffusionRate, random)

	fmt.Println("Running simulation in serial.")

	numSteps := 200

	isParallel := true
	start := time.Now()
	boards := UpdateBoards(initialBoard, numSteps, isParallel)

	fmt.Println("Simulation run. Animating system.")
	canvasWidth := 300
	frequency := 10
	images := AnimateSystem(boards, canvasWidth, frequency)

	fmt.Println("Images drawn. Generating GIF.")

	outFileName := "diffusion"
	gifhelper.ImagesToGIF(images, outFileName)
	elapsed := time.Since(start)
	fmt.Printf("Program run for %s when parallel is %t", elapsed, isParallel)
}
