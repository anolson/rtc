package main

import (
	"bytes"
	"io/ioutil"
	"math"

	"github.com/anolson/rtc/canvas"
	"github.com/anolson/rtc/color"
	"github.com/anolson/rtc/tuple"
)

type projectile struct {
	position *tuple.Tuple
	velocity *tuple.Tuple
}

func newProjectile(position, velocity *tuple.Tuple) *projectile {
	return &projectile{
		position: position,
		velocity: velocity,
	}
}

type environment struct {
	gravity *tuple.Tuple
	wind    *tuple.Tuple
}

func newEnvironment(gravity, wind *tuple.Tuple) *environment {
	return &environment{
		gravity: gravity,
		wind:    wind,
	}
}

func (env *environment) tick(proj *projectile) *projectile {
	position := tuple.Add(proj.position, proj.velocity)
	velocity := tuple.Add(proj.velocity, tuple.Add(env.gravity, env.wind))

	return newProjectile(position, velocity)
}

func main() {
	start := tuple.Point(0, 1, 0)
	velocity := tuple.Vector(1, 1.8, 0).Normalize()
	proj := newProjectile(start, tuple.Multiply(velocity, 11.25))

	gravity := tuple.Vector(0, -0.1, 0)
	wind := tuple.Vector(-0.01, 0, 0)
	env := newEnvironment(gravity, wind)

	positions := []*tuple.Tuple{}
	for {
		proj = env.tick(proj)
		if proj.position.Y <= 0 {
			break
		}

		positions = append(positions, proj.position)
	}

	red := color.RGB(0.9, 0.6, 0.75)
	c := canvas.New(900, 650)
	for _, position := range positions {
		x := int(math.Abs(math.Round(position.X)))
		y := c.Height - int(math.Round(position.Y))

		err := c.WritePixel(x, y, red)
		if err != nil {
			panic(err)
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Save(buffer)

	ioutil.WriteFile("example.ppm", buffer.Bytes(), 0644)
}
