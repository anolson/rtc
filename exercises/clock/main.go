package main

import (
	"bytes"
	"io/ioutil"
	"math"

	"github.com/anolson/rtc/canvas"
	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/matrix"
	"github.com/anolson/rtc/tuple"
)

func main() {
	// Generate points for the hours on a clock
	twelve := tuple.Point(0, 0, 1)
	hours := []*tuple.Tuple{twelve}
	for i := 1; i <= 12; i++ {
		transform := matrix.RotationY(float64(i) * math.Pi / 6)
		hour := matrix.Transform(transform, twelve)

		hours = append(hours, hour)
	}

	c := canvas.New(400, 400)
	center := tuple.Point(float64(c.Width)/2, float64(c.Height)/2, 0)
	radius := float64(c.Width) * 3 / 8
	white := color.RGB(1.0, 1.0, 1.0)

	// Write each point to the canvas
	for _, hour := range hours {
		x := hour.X*radius + center.X
		y := hour.Z*radius + center.Y

		c.WritePixel(int(x), int(y), white)
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Save(buffer)

	ioutil.WriteFile("clock.ppm", buffer.Bytes(), 0644)
}
