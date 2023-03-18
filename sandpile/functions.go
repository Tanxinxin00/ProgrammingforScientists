package main

import (
	"fmt"
	"sync"
)

// SanspileSerial takes the initial gameboard as input, updates the board in a serial manner and returns thr final stablized board.
func SandpileSerial(board *GameBoard) *GameBoard {

	//while the board is still updatable, update the board.
	for board.IsUpdatable() {
		board.UpdateBoard()
	}

	return board
}

// UpdateBoard is a method of the board that consecutes the topple operations while ranging through all the squares once.
func (b *GameBoard) UpdateBoard() {
	rowNum := len(*b)
	colNum := len((*b)[0])

	//range through all the squares on the board
	for r := range *b {
		for c := range (*b)[r] {
			//If the current number of coins is equal or greater than 4
			//reduce the coins in the current square by 4 and add 1 to the left, right, top, bottom neighbor squares.
			if (*b)[r][c] >= 4 {
				(*b)[r][c] -= 4
				//only the square positions on the board will be added
				//some of the coins may just fall off the board
				if r-1 >= 0 {
					(*b)[r-1][c] += 1
				}
				if r+1 < rowNum {
					(*b)[r+1][c] += 1
				}
				if c-1 >= 0 {
					(*b)[r][c-1] += 1
				}
				if c+1 < colNum {
					(*b)[r][c+1] += 1
				}
			}
		}
	}

}

// IsUpdatable is a mothod of the gameboard that determines whether the currentboard can still be updated or not.
func (b *GameBoard) IsUpdatable() bool {

	//range through all the squares in the board, as long as there is one square with more than 4 coins, the board is updatable.
	for r := range *b {
		for c := range (*b)[r] {
			if (*b)[r][c] >= 4 {
				return true
			}
		}
	}

	return false
}

// SanpileParallel takes a gameboard and the number of processors as input, updates the board in a parallel manner and returns the final stablized board.
func SandpileParallel(b *GameBoard, numProcs int) *GameBoard {

	//while the board is still updatable, update the board with multi processors.
	for b.IsUpdatable() {
		b.SandpileMultiprocs(numProcs)
	}

	return b
}

func (b *GameBoard) PrintBoard() {
	for r := range *b {
		fmt.Println((*b)[r])
	}
}

// waitgroup is used to track the progress of all the goroutines.
var wg sync.WaitGroup

func (b *GameBoard) SandpileMultiprocs(numProcs int) {
	rowNum := len(*b)
	//ch is a channel of a slice, in which slices of 2 intergers indicating the position of the square is stored.
	ch := make(chan [][2]int)

	//for each processor, assign several rows of the board to execute the topple process.
	for i := 0; i < numProcs; i++ {
		startIndex := i * rowNum / numProcs
		endIndex := (i + 1) * rowNum / numProcs

		if i < numProcs-1 {
			//For each processing goroutine, add 1 to the waitgroup.
			wg.Add(1)
			go SandpileOneProc((*b)[startIndex:endIndex], startIndex, ch)
		} else {
			//for the last processor, assign the remainder of the board.
			wg.Add(1)
			go SandpileOneProc((*b)[startIndex:], startIndex, ch)
		}

	}

	// Another way of making sure all the goroutines are finished before dealing with thew fallen coins.
	// But it is slower so is abandoned.
	// var DiffusionSquares [][2]int
	// for i := 0; i < numProcs; i++ {
	// 	diff := <-ch
	// 	DiffusionSquares = append(DiffusionSquares, diff...)
	// }

	// for i := range DiffusionSquares {
	// 	if DiffusionSquares[i][0] >= 0 && DiffusionSquares[i][0] < rowNum {
	// 		newBoard[DiffusionSquares[i][0]][DiffusionSquares[i][1]] += 1
	// 	}
	// }

	//wait for all the waitgroups to be finished to go on with dealing with the fallen coins.
	wg.Wait()
	//receive the slices from the channels
	for i := 0; i < numProcs; i++ {
		diff := <-ch
		//range thorugh each slice
		for j := range diff {
			//add the fallen coins to the indicated positions when still in board.
			if diff[j][0] >= 0 && diff[j][0] < rowNum {
				(*b)[diff[j][0]][diff[j][1]] += 1
			}
		}
	}

}

// SandpileOneProc takes a subslice of the gameboard, the startrow index of the subslice with regard to the original board and a channel as input.
// Range through all the squares in the subboard, do the topple operations while available and send the positions of the other parts of the board where coins need to be added as a slice to the channel.
func SandpileOneProc(b GameBoard, Sindex int, ch chan [][2]int) {
	//CoinPositions is a slice of size2 interger slices that stores the positions of the board where the coins need to be added.
	var CoinPositions [][2]int

	rowNum := len((b))
	colNum := len((b)[0])

	for r := range b {
		for c := range b[r] {
			if b[r][c] >= 4 {
				b[r][c] -= 4
				//if the top of the current square is still in this subboard, add 1 to it.
				if r-1 > 0 {
					b[r-1][c] += 1
				} else { //if the top exceeds the current subboard, add the position to the slice
					CoinPositions = append(CoinPositions, [2]int{r + Sindex - 1, c})
				}

				//if the bottom of the current square is still in this subboard, add 1 to it.
				if r+1 < rowNum {
					b[r+1][c] += 1
				} else { //if the bottom exceeds the current subboard, add the position to the slice
					CoinPositions = append(CoinPositions, [2]int{r + Sindex + 1, c})
				}

				//if the left of the current square is still in this subboard, add 1 to it.
				if c-1 >= 0 {
					b[r][c-1] += 1
				} //else just let it fall off the board

				//if the right of the current square is still in this subboard, add 1 to it.
				if c+1 < colNum {
					b[r][c+1] += 1
				} //else just let it fall off the board
			}
		}
	}
	//When this goroutine is finished, reduce 1 from the waitgroup.
	wg.Done()
	//pass the position slice to the channel
	ch <- CoinPositions

}

// CheckBoard is used to check if two boards are identical.
func CheckBoard(b1, b2 *GameBoard) bool {
	//If the rownumber of the two boards are different, the boards are different.
	if len(*b1) != len(*b2) {
		return false
	} else {
		for i := range *b1 {
			//If the column number of the two boards are different, the boards are different.
			if len((*b1)[i]) != len((*b2)[i]) {
				return false
			} else {
				//range through each element of the boards, if any of the value is different, the boards are different.
				for j := range (*b1)[i] {
					if (*b1)[i][j] != (*b2)[i][j] {
						fmt.Printf("The %d row %d col is different", i, j)
						fmt.Printf("In board1 it is %d;In board2 it is%d.", (*b1)[i][j], (*b2)[i][j])
						return false
					}
				}
			}
		}
	}
	return true
}
