package main

import (
	"image"
	"image/color"
)

func AnimateSandpile(timePoints []*GameBoard, size int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		img := DrawBoard(timePoints[i], size)
		images = append(images, img)
	}

	return images
}

func DrawBoard(board *GameBoard, size int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for row := range *board {
		for col := range (*board)[row] {

			if (*board)[row][col] == 0 {
				img.Set(row, col, color.RGBA{0, 0, 0, 255})
			} else if (*board)[row][col] == 1 {
				img.Set(row, col, color.RGBA{85, 85, 85, 255})
			} else if (*board)[row][col] == 2 {
				img.Set(row, col, color.RGBA{170, 170, 170, 255})
			} else if (*board)[row][col] >= 3 {
				img.Set(row, col, color.RGBA{255, 255, 255, 255})
			} else {
				panic("Wrong sandpile!")
			}
		}
	}
	return img
}
