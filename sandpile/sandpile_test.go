package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSandpileSerial(t *testing.T) {
	//declare test type
	type test struct {
		inputboard  *GameBoard
		outputboard *GameBoard
	}

	inputDirectory := "tests/SandpileSerial/input/"
	outputDirectory := "tests/SandpileSerial/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].inputboard = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].outputboard = ReadBoardFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := SandpileSerial(test.inputboard)
		//check if the board given by the function is exactly idential to the one given in outputfile
		if !BoardistheSame(outcome, test.outputboard) {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Println("Correct!")
		}
	}

}
func TestUpdateBoard(t *testing.T) {
	//declare test type
	type test struct {
		inputboard  *GameBoard
		outputboard *GameBoard
	}

	inputDirectory := "tests/UpdateBoard/input/"
	outputDirectory := "tests/UpdateBoard/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].inputboard = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].outputboard = ReadBoardFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		test.inputboard.UpdateBoard()

		//check if the updated inputboard is exactly identical to the outputboard given in outputfile.
		if !BoardistheSame(test.inputboard, test.outputboard) {
			t.Errorf("Error! For input test dataset %d", i)
		} else {
			fmt.Println("Correct!")
		}
	}

}

func TestIsUpdatable(t *testing.T) {
	//declare test type
	type test struct {
		board  *GameBoard
		answer bool
	}

	inputDirectory := "tests/IsUpdatable/input/"
	outputDirectory := "tests/IsUpdatable/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].board = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadBoolFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := test.board.IsUpdatable()
		//check if the function gives the correct judgement reagarding the specific board
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d", i)
		} else {
			fmt.Println("Correct! When the frequency map is", test.board, "the richness is", test.answer)
		}
	}

}

func TestSandpileMultiprocs(t *testing.T) {
	//declare test type
	type test struct {
		inputboard  *GameBoard
		numProcs    int
		outputboard *GameBoard
	}

	inputDirectory := "tests/SandpileMultiprocs/input/"
	inputDirectory2 := "tests/SandpileMultiprocs/input2/"
	outputDirectory := "tests/SandpileMultiprocs/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].inputboard = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].numProcs = ReadIntegerFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputboard = ReadBoardFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		test.inputboard.SandpileMultiprocs(test.numProcs)
		//check if the updated inputboard is identical to the one given by outputfile
		if !BoardistheSame(test.inputboard, test.outputboard) {
			t.Errorf("Error! For input test dataset %d", i)
		} else {
			fmt.Println("Correct!")
		}
	}

}
func TestSandpileParallel(t *testing.T) {
	//declare test type
	type test struct {
		inputboard  *GameBoard
		numProcs    int
		outputboard *GameBoard
	}

	inputDirectory := "tests/SandpileParallel/input/"
	inputDirectory2 := "tests/SandpileParallel/input2/"
	outputDirectory := "tests/SandpileParallel/output/"

	//assert that files are non-empty and have the same length
	inputFiles := ReadFilesFromDirectory(inputDirectory)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].inputboard = ReadBoardFromFile(inputDirectory, inputFiles[i])
		tests[i].numProcs = ReadIntegerFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputboard = ReadBoardFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := SandpileParallel(test.inputboard, test.numProcs)
		//check if the outcome and the outputboard is exactly identical
		if !BoardistheSame(outcome, test.outputboard) {
			t.Errorf("Error! For input test dataset %d", i)
		} else {
			fmt.Println("Correct!")
		}
	}

}

func TestInitializeBoard(t *testing.T) {
	//declare test type
	type test struct {
		size        int
		pile        int
		place       string
		outputboard *GameBoard
	}

	inputDirectory := "tests/InitialBoard/input/"
	outputDirectory := "tests/InitialBoard/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)
	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].size, tests[i].pile, tests[i].place = ReadInitializationFromFile(inputDirectory, inputFiles[i])
		tests[i].outputboard = ReadBoardFromFile(outputDirectory, outputFiles[i])
		//note that some of the outputboard (when the placement is random) will not be used in testing, pseudo boards could be put in those outputfiles.
	}

	for i, test := range tests {
		//if the placement is "central",check if the outcomeboard is exactly identical to the outputboard
		if test.place == "central" {
			outcome := InitializeBoard(test.size, test.pile, test.place)
			if !BoardistheSame(outcome, test.outputboard) {
				t.Errorf("Error! For input test dataset %d", i)
			} else {
				fmt.Println("Correct!")
			}

			//if the placement is random, simpley check if the outcome board has the correct size and number of coins
			//since the positions where the coins will be put on will be different.
		} else if test.place == "random" {
			outcome := InitializeBoard(test.size, test.pile, test.place)
			if len(*outcome) != test.size {
				t.Errorf("wrong size")
			}
			//sum is the total number of coins on the board
			sum := 0
			for r := range *outcome {
				if len((*outcome)[r]) != test.size {
					t.Errorf("wrong size")
				}
				for c := range (*outcome)[r] {
					sum += (*outcome)[r][c]
				}
			}

			if sum != test.pile {
				t.Errorf("wrong number of coins")
			}
		}
	}

}

func ReadBoardFromFile(directory string, inputFile os.FileInfo) *GameBoard {
	fileName := inputFile.Name() //grab file name

	//read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	board := make(GameBoard, len(inputLines))
	b := &board
	//each inputLine is a row in the board
	for row, inputLine := range inputLines {

		//read out the current line
		currentLine := strings.Split(inputLine, " ")
		//currentLine has several strings correponding to the elements of a row in the board
		board[row] = make([]int, len(currentLine))

		//convert string into integer as the board value.
		for col, str := range currentLine {
			board[row][col], err = strconv.Atoi(str)
		}
		if err != nil {
			panic(err)
		}

	}

	return b
}

func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in input directory.")
	}
	if length1 == 0 {
		panic("No files present in output directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}

func ReadBoolFromFile(directory string, inputFile os.FileInfo) bool {
	fileName := inputFile.Name() //grab file name

	//ead in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	if string(fileContents) == "true" {
		return true
	} else if string(fileContents) == "false" {
		return false
	} else {
		panic("Wrong boolean given")
	}

}

func BoardistheSame(a, b *GameBoard) bool {
	if len(*a) != len(*b) {
		return false
	}

	for i := range *a {
		if len((*a)[i]) != len((*b)[i]) {
			return false
		}
		for j := range (*a)[i] {
			if (*a)[i][j] != (*b)[i][j] {
				return false
			}
		}
	}

	return true
}

func ReadIntegerFromFile(directory string, file os.FileInfo) int {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.Atoi(outputLines[0])

	if err != nil {
		panic(err)
	}

	return answer
}

func ReadInitializationFromFile(directory string, file os.FileInfo) (int, int, string) {
	fileName := file.Name() //grab file name

	//read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//the first string in the inputline is the size
	size, err := strconv.Atoi(inputLines[0])
	if err != nil {
		panic(err)
	}

	//the second string in the inputline is the pilenumber
	pile, err := strconv.Atoi(inputLines[1])
	if err != nil {
		panic(err)
	}

	//the third string in the inputline is the placement pattern
	place := inputLines[2]

	return size, pile, place

}
