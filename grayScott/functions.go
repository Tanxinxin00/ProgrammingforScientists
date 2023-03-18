//This is Tanxin Qiao's HW0 assignment.
//Due Sep 12, 2022

package main

// Function SimulateGrayScott is the root function for the GrayScott simulation. It takes in all the initial parameters  and initial board needed and
// returns with all the uodated board generated.
func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) []Board {
	boards := make([]Board, numGens+1)
	boards[0] = initialBoard

	for i := 1; i < numGens+1; i++ {
		boards[i] = UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	}

	return boards
}

// Function UpdateBoard takes in a board and the parameters(rates)
// and gives the next generation of board.
func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board {
	rownum := len(currentBoard)
	colnum := len(currentBoard[0])
	newBoard := InitializeBoard(rownum, colnum)

	//range thorough every single cell in the board and retrieve the updated cells for the next board.
	for r := 0; r < rownum; r++ {
		for c := 0; c < colnum; c++ {
			newBoard[r][c] = UpdateCell(currentBoard, r, c, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
		}
	}

	return newBoard
}

// InitializeBoard creates a blank board with given row number and column number.
func InitializeBoard(row, col int) Board {
	newBoard := make(Board, row)

	for r := 0; r < row; r++ {
		newBoard[r] = make([]Cell, col)
	}

	return newBoard
}

// UpdateCell takes in all the rates and updates the assigned cell in the current board to the next generation.
func UpdateCell(currentBoard Board, row, col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
	currentCell := currentBoard[row][col]

	diffusionValues := ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
	reactionValues := ChangeDueToReactions(currentCell, feedRate, killRate)

	//use SumCells to add diffusionValues and reactionValues to the current values of the cell to obtain the new values of the cell.
	return SumCells(currentCell, diffusionValues, reactionValues)
}

// SumCells takes in an uncertain amount of Cell type values, sum the values and gets a new Cell type variable.
func SumCells(cells ...Cell) Cell {
	var sum Cell

	//the cell type inputs form an array that can be ranged over with their indexes.
	for i := range cells {
		sum[0] += cells[i][0]
		sum[1] += cells[i][1]
	}

	return sum
}

// ChangeDueToReactions takes the feedrate and killrate and follows the given formula
// to calculate the changes in concentration of particles due to reactions between the particles.
func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {
	parta := currentCell[0]
	partb := currentCell[1]

	var cellreaction Cell
	cellreaction[0] = feedRate*(1.0-parta) - parta*partb*partb
	cellreaction[1] = parta*partb*partb - killRate*partb

	return cellreaction
}

// CHangeDueToDiffusion takes the location of the current cell, the diffusion rates and the common kernel as input
// and calcutes the changes in  concentration of particles caused by diffusion.
func ChangeDueToDiffusion(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
	var celldiffusionchange Cell

	//First find the MooreNeighborhoods(which is the 3x3 matrix surrounding the cell) of the current cell.
	Moores := MooreNeighborhood(currentBoard, row, col)

	//for each of the particle, a different diffusion rate is taken and used to adjust the kernel
	//to calculate the change caused by diffusion to the cell's neighborhoohs.
	Akernel := SlowKernel(preyDiffusionRate, kernel)
	diffusionchangeA := convolution(Moores, Akernel, 0)
	celldiffusionchange[0] = diffusionchangeA

	Bkernel := SlowKernel(predatorDiffusionRate, kernel)
	diffusionchangeB := convolution(Moores, Bkernel, 1)
	celldiffusionchange[1] = diffusionchangeB

	return celldiffusionchange
}

// SLowKernel takes in a rate and a kernel
// multiplies the rate by every value in the kernel to formulate a new kernel
// (Usually a kernel with smaller values because the rate is often smaller than 1.0)
func SlowKernel(rate float64, kernel [3][3]float64) [3][3]float64 {
	var newker [3][3]float64

	for r := range newker {
		for c := range newker[r] {
			newker[r][c] = kernel[r][c] * rate
		}
	}

	return newker
}

// MooreNeighborhhood takes in the position of a current cell
// and produces a 3x3 matrix composed of the neighbors of the cell with the cell in the centre of the matrix.
func MooreNeighborhood(currentBoard Board, row, col int) Board {
	Moores := make(Board, 3)

	for i := 0; i < 3; i++ {
		Moores[i] = make([]Cell, 3)

		for j := 0; j < 3; j++ {
			//Range over the surrounding cells of the current cell
			//Use InField function to check if the surrounding cells of the current cell is still in the board
			if InField(currentBoard, row-1+i, col-1+j) {
				Moores[i][j] = currentBoard[row-1+i][col-1+j]

				//If the current cell is on the edge of the board,
				// then its neighborhoods that are outside of the board should be deemed as empty cells with no particle at all
			} else {
				Moores[i][j][0] = 0
				Moores[i][j][1] = 0
			}
		}
	}

	return Moores
}

// convolution calculates each of the value in the Mooreneighborhood of a cell and each of the value in the kernel respectively.
// the int variable input uses index to indicate which particle we are processing, with 0 being particle A and 1 being particle B.
func convolution(Moores Board, kernel [3][3]float64, particle int) float64 {
	addconcen := 0.0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			addconcen += Moores[i][j][particle] * kernel[i][j]
		}
	}

	return addconcen
}

// The function Infield takes in the current location of a cell and checks if the cell is in the whole board or not.
func InField(currentBoard Board, row, col int) bool {
	inboard := false
	rownum := len(currentBoard)
	colnum := len(currentBoard[0])

	//inboard is set to be false by default
	//and if the row and column number of the current cell are both in the range of the whole board, make inboard true.
	if row >= 0 && row < rownum {
		if col >= 0 && col < colnum {
			inboard = true
		}
	}

	return inboard
}
