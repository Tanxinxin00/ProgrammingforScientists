package main

import (
	"math/rand"
	"sort"
	"time"
)

// Initialize genrates a board of size*size, with altogether pile number of coins.
// If the placment is central, all the coins will be put on the center square
// if the placement is random, the coins will be randomly distributed over 100 random positions.
func InitializeBoard(size, pile int, place string) *GameBoard {
	//create a new size*size GameBoard
	var board GameBoard
	b := &board
	board = make([][]int, size)
	for r := range board {
		board[r] = make([]int, size)
	}

	if place == "central" {
		center := size / 2
		board[center][center] = pile
	} else {
		rand.Seed(time.Now().UnixNano())

		//create a slice of integers indicating the number of coins distributed on 100 positions.
		randomNumbers := make([]int, 101)
		randomNumbers[0] = 0
		randomNumbers[100] = pile

		//generate 99 random numbers less than the total number of coins that could be used as separators.
		for i := 1; i <= 99; i++ {
			randomNumbers[i] = rand.Intn(pile)
		}
		//sort the numbers in increasing order
		sort.Ints(randomNumbers)

		//for each random position, put the number the coins given by the current separator minus the last separator.
		for i := 1; i <= 100; i++ {
			r := rand.Intn(size)
			c := rand.Intn(size)
			board[r][c] = randomNumbers[i] - randomNumbers[i-1]
		}
	}

	return b
}

func CopyBoard(b1 *GameBoard) *GameBoard {
	var board GameBoard
	b2 := &board
	board = make([][]int, len(*b1))
	for i := range *b1 {
		board[i] = make([]int, len((*b1)[i]))
		copy(board[i], (*b1)[i])

	}
	return b2
}
