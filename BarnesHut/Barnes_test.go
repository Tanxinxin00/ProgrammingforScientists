package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMakeQuadTree(t *testing.T) {
	//first, declare our test type
	type test struct {
		input  *Universe
		output *QuadTree
	}

	inputDirectory := "tests/MakeQuadTree/input/"
	outputDirectory := "tests/MakeQuadTree/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	tests := make([]test, len(inputFiles))
	for i := range inputFiles {
		tests[i].input = ReadUniverseFromDirectory(inputDirectory, inputFiles[i])
		tests[i].output = ReadQuadTreeFromDirectory(outputDirectory, outputFiles[i])
	}
	for i, test := range tests {
		outcome := MakeQuadTree(test.input)

		if QuadTreeistheSame(outcome, test.output) != true {
			fmt.Println("Your code gives", outcome, "The answer is", test.output)

			t.Errorf("Error! For input test dataset %d failed", i)
		} else {
			fmt.Println("Correct!")
		}
	}

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

func ReadUniverseFromDirectory(directory string, inputFile os.FileInfo) *Universe {
	var inputuniverse Universe
	u := &inputuniverse

	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
	starNum := len(inputLines) - 2

	u.stars = make([]*Star, starNum)

	contentIndex := 0
	starIndex := 0
	for _, inputLine := range inputLines {
		if inputLine == "-" {
			contentIndex += 1
			continue
		}
		if contentIndex == 0 {
			u.width, err = strconv.ParseFloat(inputLine, 64)
		} else {
			var s Star
			u.stars[starIndex] = &s
			currentLine := strings.Split(inputLine, " ")

			s.position.x, err = strconv.ParseFloat(currentLine[0], 64)
			if err != nil {
				panic(err)
			}
			s.position.y, err = strconv.ParseFloat(currentLine[1], 64)
			if err != nil {
				panic(err)
			}
			s.velocity.x, err = strconv.ParseFloat(currentLine[2], 64)
			if err != nil {
				panic(err)
			}
			s.velocity.y, err = strconv.ParseFloat(currentLine[3], 64)
			if err != nil {
				panic(err)
			}
			s.acceleration.x, err = strconv.ParseFloat(currentLine[4], 64)
			if err != nil {
				panic(err)
			}
			s.acceleration.y, err = strconv.ParseFloat(currentLine[5], 64)
			if err != nil {
				panic(err)
			}
			s.mass, err = strconv.ParseFloat(currentLine[6], 64)
			if err != nil {
				panic(err)
			}
			s.radius, err = strconv.ParseFloat(currentLine[7], 64)
			if err != nil {
				panic(err)
			}
			starIndex += 1

		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(u.stars[0], u.stars[1], u.stars[2], u.stars[3])
	return u

}

func ReadQuadTreeFromDirectory(directory string, outputFile os.FileInfo) *QuadTree {
	var quad QuadTree
	q := &quad
	fileName := outputFile.Name() //grab file name
	var rtnode Node
	q.root = &rtnode
	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	ParentNodes := make([]*Node, len(inputLines))
	CurrentParentIndex := 0
	for row, inputLine := range inputLines {
		if row == 0 {
			var star Star
			s := &star
			currentLine := strings.Split(inputLine, " ")
			s.position.x, err = strconv.ParseFloat(currentLine[0], 64)
			if err != nil {
				panic(err)
			}
			s.position.y, err = strconv.ParseFloat(currentLine[1], 64)
			if err != nil {
				panic(err)
			}
			s.velocity.x, err = strconv.ParseFloat(currentLine[2], 64)
			if err != nil {
				panic(err)
			}
			s.velocity.y, err = strconv.ParseFloat(currentLine[3], 64)
			if err != nil {
				panic(err)
			}
			s.acceleration.x, err = strconv.ParseFloat(currentLine[4], 64)
			if err != nil {
				panic(err)
			}
			s.acceleration.y, err = strconv.ParseFloat(currentLine[5], 64)
			if err != nil {
				panic(err)
			}

			s.mass, err = strconv.ParseFloat(currentLine[6], 64)
			if err != nil {
				panic(err)
			}
			s.radius, err = strconv.ParseFloat(currentLine[7], 64)
			if err != nil {
				panic(err)
			}
			q.root.star = s
			ParentNodes[CurrentParentIndex] = q.root
			CurrentParentIndex -= 1

			if err != nil {
				panic(err)
			}
			continue
		}

		if inputLine == "EnteringChildrenNodes" {
			CurrentParentIndex += 1
			continue
		}

		if inputLine == "LeavingChildrenNodes" {
			CurrentParentIndex -= 1
			continue
		}

		if ParentNodes[CurrentParentIndex].children == nil {

			var node1, node2, node3, node4 Node
			v1 := &node1
			v2 := &node2
			v3 := &node3
			v4 := &node4
			ParentNodes[CurrentParentIndex].children = []*Node{v1, v2, v3, v4}
		}

		var star Star
		s := &star
		currentLine := strings.Split(inputLine, " ")
		s.position.x, err = strconv.ParseFloat(currentLine[0], 64)
		if err != nil {
			panic(err)
		}
		s.position.y, err = strconv.ParseFloat(currentLine[1], 64)
		if err != nil {
			panic(err)
		}
		s.velocity.x, err = strconv.ParseFloat(currentLine[2], 64)
		if err != nil {
			panic(err)
		}
		s.velocity.y, err = strconv.ParseFloat(currentLine[3], 64)
		if err != nil {
			panic(err)
		}
		s.acceleration.x, err = strconv.ParseFloat(currentLine[4], 64)
		if err != nil {
			panic(err)
		}
		s.acceleration.y, err = strconv.ParseFloat(currentLine[5], 64)
		if err != nil {
			panic(err)
		}
		s.mass, err = strconv.ParseFloat(currentLine[6], 64)
		if err != nil {
			panic(err)
		}
		s.radius, err = strconv.ParseFloat(currentLine[7], 64)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}

		for i := range ParentNodes[CurrentParentIndex].children {
			if ParentNodes[CurrentParentIndex].children[i].star == nil {
				ParentNodes[CurrentParentIndex].children[i].star = s
				if inputLines[row+1] == "EnteringChildrenNodes" {
					ParentNodes[CurrentParentIndex+1] = ParentNodes[CurrentParentIndex].children[i]

				}
				break
			}
		}
	}

	return q
}

func QuadTreeistheSame(q1, q2 *QuadTree) bool {
	same := true
	same = CompareRootNodes(q1.root, q2.root)
	fmt.Println(q1.root.star, q1.root.children[0].star, q1.root.children[1].star, q1.root.children[2].star, q1.root.children[3].star)
	return same
}

func CompareRootNodes(root1, root2 *Node) bool {
	same := true
	if !StarsAreSame(root1.star, root2.star) {
		same = false
	} else {
		for i := range root1.children {
			if root1.children[i].star == nil {
				continue
			}
			if !CompareRootNodes(root1.children[i], root2.children[i]) {
				same = false
				break
			}
		}
	}
	return same
}

func StarsAreSame(s1, s2 *Star) bool {
	same := true
	if s1.position.x != s2.position.x {
		same = false
	} else if s1.position.y != s2.position.y {
		same = false
	} else if s1.mass != s2.mass {
		same = false
	} else if s1.radius != s2.radius {
		same = false
	}
	return same
}
