package main

import (
	"bytes"
	"io/ioutil"

	"github.com/anolson/rtc/canvas"
	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/primitives"
	"github.com/anolson/rtc/ray"
	"github.com/anolson/rtc/tuple"
)

func main() {
	rayOrigin := tuple.Point(0, 0, -5)

	var wallZ float64 = 10
	var wallSize float64 = 7
	var half float64 = wallSize / 2
	var canvasPixels int = 500
	var pixelSize float64 = wallSize / float64(canvasPixels)

	c := canvas.New(canvasPixels, canvasPixels)
	red := color.RGB(1, 0, 0)
	sphere := primitives.NewSphere()

	for y := 0; y < canvasPixels; y++ {
		for x := 0; x < canvasPixels; x++ {
			worldY := half - pixelSize*float64(y)
			worldX := -half + pixelSize*float64(x)

			position := tuple.Point(worldX, worldY, wallZ)
			direction := tuple.Subtract(position, rayOrigin)

			r := ray.New(rayOrigin, direction.Normalize())
			intersections := sphere.Intersect(r)

			if primitives.Hit(intersections) != nil {
				c.WritePixel(x, y, red)
			}
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Save(buffer)

	ioutil.WriteFile("sphere.ppm", buffer.Bytes(), 0644)
}
