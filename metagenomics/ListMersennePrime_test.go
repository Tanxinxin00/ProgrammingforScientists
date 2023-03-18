package main

/*
import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

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

func TestListMersennePrimes(t *testing.T) {
	type test struct {
		integer int
		answer  []int
	}

	inputDirectory := "tests/ListMersennePrime/input/"
	outputDirectory := "tests/ListMersennePrime/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//first, retrieve all the input and output files and keep each pair of files in a test.
	tests := make([]test, len(inputFiles))
	for i := range inputFiles {
		tests[i].integer = ReadIntegerFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadListFromFile(outputDirectory, outputFiles[i])

	}
	// range through all the test to check if the outcome given by the function ListMersennePrimes is the same as the correct answer given by the outputfiles.
	for i, test := range tests {
		outcome := ListMersennePrimes(test.integer)
		if ListIsSame(outcome, test.answer) {
			fmt.Println("Correct! When the integer is", test.integer, "the MersennePrimes are", test.answer)
		} else {
			t.Errorf("Error! For input test dataset %d, your code gives %d, and the correct answer is %d", i, outcome, test.answer)
		}
	}

}

// function LIstIsSame takes two list of integers as input and range through all the values in the list to see if the tao lists are identical.
func ListIsSame(lista, listb []int) bool {
	same := true

	//first compare if the two lists have the same length
	// if they don't, then they are not identical lists and don't need to range through them.
	if len(lista) == len(listb) {

		//range through the lists to check if each value is the same
		for i := range lista {

			//once there's an unmatching element, don't need to go through the rest of the loop.
			if lista[i] != listb[i] {
				same = false
				break
			}
		}
	} else {
		same = false
	}

	return same
}

// function ReadListFromFile reads a single line from file and obtains a list of integers.
func ReadListFromFile(directory string, outputFile os.FileInfo) []int {
	fileName := outputFile.Name()

	//first read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//since the answer only contains a single list, just use the first line.
	//split the line into elements in a list with " " and convert each string element into integer.
	answerlist := strings.Split(outputLines[0], " ")
	answer := make([]int, len(answerlist))

	for i := range answerlist {
		answer[i], err = strconv.Atoi(answerlist[i])

		if err != nil {
			panic(err)
		}

	}

	return answer
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
*/
