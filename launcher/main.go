package main

import (
	"fmt"

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

func main() {
	gravity := tuple.Vector(0, -0.1, 0)
	wind := tuple.Vector(-0.01, 0, 0)
	env := newEnvironment(gravity, wind)

	position := tuple.Point(0, 1, 0)
	velocity := tuple.Vector(0, 1, 0).Normalize()
	proj := newProjectile(position, tuple.Multiply(velocity, 1))

	tickCount := 0
	for {
		proj = tick(env, proj)
		if proj.position.Y <= 0 {
			break
		}

		tickCount++
		fmt.Printf("Projectile now at - x: %0.2f y: %0.2f\n", proj.position.X, proj.position.Y)
	}

	fmt.Printf("Total ticks: %v\n", tickCount)
}

func tick(env *environment, proj *projectile) *projectile {
	position := tuple.Add(proj.position, proj.velocity)
	velocity := tuple.Add(proj.velocity, tuple.Add(env.gravity, env.wind))

	return newProjectile(position, velocity)
}
