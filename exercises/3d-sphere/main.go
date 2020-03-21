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
	half := wallSize / 2
	canvasPixels := 500
	pixelSize := wallSize / float64(canvasPixels)

	c := canvas.New(canvasPixels, canvasPixels)

	material := primitives.DefaultMaterial()
	material.Color = color.RGB(1, 0.2, 1)

	sphere := primitives.NewSphere()
	sphere.Material = material

	lightPosition := tuple.Point(-10, 10, -10)
	lightColor := color.RGB(1, 1, 1)
	light := primitives.NewPointLight(lightPosition, lightColor)

	for y := 0; y < canvasPixels; y++ {
		for x := 0; x < canvasPixels; x++ {
			worldY := half - pixelSize*float64(y)
			worldX := -half + pixelSize*float64(x)

			position := tuple.Point(worldX, worldY, wallZ)
			direction := tuple.Subtract(position, rayOrigin)

			r := ray.New(rayOrigin, direction.Normalize())
			intersections := sphere.Intersect(r)

			if hit := primitives.Hit(intersections); hit != nil {
				point := r.Position(hit.T)
				normal := sphere.NormalAt(point)
				eye := tuple.Negate(r.Direction)

				result := material.Lighting(light, point, eye, normal)

				c.WritePixel(x, y, result)
			}
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Save(buffer)

	ioutil.WriteFile("sphere.ppm", buffer.Bytes(), 0644)
}
