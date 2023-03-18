package main

import (
	"canvas"
	"image"
)

func AnimateSystem(timePoints []*Sky, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { //only draw if current index of sky is divisible by the frequency parameter
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Sky
func DrawToCanvas(s *Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the boidss and draw them.
	for _, b := range s.boids {
		//The birds have a yellow color.
		c.SetFillColor(canvas.MakeColor(180, 135, 45))

		//The birds are represented by circles.
		cx := (b.position.x / s.width) * float64(canvasWidth)
		cy := (b.position.y / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, 5.0)

		c.Fill()
	}

	return c.GetImage()
}
