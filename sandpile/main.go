package main

import (
	"fmt"
	// "gifhelper"
	"image/png"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	size, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if size < 0 {
		panic("Negative Board size given.")
	}

	pile, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	if pile < 0 {
		panic("Negative number of piles given.")
	}

	placement := os.Args[3]
	if placement != "central" && placement != "random" {
		panic("Wrong placement pattern given.")
	}

	fmt.Println("Command line arguments read successfully.")

	InitialBoard := InitializeBoard(size, pile, placement)
	InitialBoard2 := CopyBoard(InitialBoard)
	fmt.Println("Initial board initialized successfully.")

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	runtime.GOMAXPROCS(1)
	start1 := time.Now()
	fmt.Println("Serial sandpile simulation is about to start.")
	FinalBoard1 := SandpileSerial(InitialBoard)
	elapsed1 := time.Since(start1)
	fmt.Printf("Serial sandpile simulated successfully,it took %s\n", elapsed1)

	/* making gif(better choose to make gifs when the placment is random)
	images1 := AnimateSandpile(timePoints1, size)
	gifname1 := fmt.Sprintf("%d_%d_%d_%d_%s_Serial", time.Now().Hour(), time.Now().Minute(), size, pile, placement)
	gifhelper.ImagesToGIF(images1, gifname1)
	fmt.Println("Sandpile serial simulation gif produced!")
	*/

	FinalImage1 := DrawBoard(FinalBoard1, size)
	pngname1 := fmt.Sprintf("FinalBoard_%d_%d_%d_%d_%s_Serial.png", time.Now().Hour(), time.Now().Minute(), size, pile, placement)
	outBoard1, _ := os.Create(pngname1)
	err3 := png.Encode(outBoard1, FinalImage1)
	if err3 != nil {
		panic(err3)
	}

	// fmt.Println("FinalBoard_Serial image generated!")
	// start1 := time.Now()
	// finalboard := Update(InitialBoard)
	// elapsed1 := time.Since(start1)
	// fmt.Printf("Serial sandpile simulated successfully,it took %s\n", elapsed1)
	// image := DrawBoard(finalboard, size)
	// pngname1 := fmt.Sprintf("FinalBoard_%d_%d_%d_%d_%s_Serial.png", time.Now().Hour(), time.Now().Minute(), size, pile, placement)
	// outBoard1, _ := os.Create(pngname1)
	// err3 := png.Encode(outBoard1, image)
	// if err3 != nil {
	// 	panic(err3)
	// }

	// fmt.Println("FinalBoard_Serial image generated!")
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	numProcs := 8
	runtime.GOMAXPROCS(numProcs)
	start2 := time.Now()
	fmt.Println("Parallel sandpile simulation is about to start.")
	FinalBoard2 := SandpileParallel(InitialBoard2, numProcs)
	elapsed2 := time.Since(start2)
	fmt.Printf("Parallel sandpile simulated successfully,it took %s\n", elapsed2)

	//making gif(better choose to make gifs when the placment is random)
	// images2 := AnimateSandpile(timePoints2, size)
	// gifname2 := fmt.Sprintf("%d_%d_%d_%d_%s_Parallel", time.Now().Hour(), time.Now().Minute(), size, pile, placement)
	// gifhelper.ImagesToGIF(images2, gifname2)
	// fmt.Println("Sandpile parallel simulation gif produced!")

	FinalImage2 := DrawBoard(FinalBoard2, size)
	pngname2 := fmt.Sprintf("FinalBoard_%d_%d_%d_%d_%s_Parallel.png", time.Now().Hour(), time.Now().Minute(), size, pile, placement)
	outBoard2, _ := os.Create(pngname2)
	err4 := png.Encode(outBoard2, FinalImage2)
	if err4 != nil {
		panic(err4)
	}

	fmt.Println("FinalBoard_Parallel image generated!")

	fmt.Println("Are the boards identical? ", CheckBoard(FinalBoard1, FinalBoard2))

	fmt.Println("Exiting normally")
}
