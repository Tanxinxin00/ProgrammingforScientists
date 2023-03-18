package main

import (
	"math"
	"math/rand"
	"time"
)

// this is where we will put functions that correspond only to the parallel simulation.
func (b *Board) DiffusionParallel(numProcs int) {
	numParticles := len(b.particles)

	for i := 0; i < numProcs; i++ {
		finished := make(chan bool)
		startIndex := i * numParticles / numProcs
		var endIndex int
		if i < numProcs-1 {
			endIndex = (i + 1) * numParticles / numProcs
		} else {
			endIndex = numParticles
		}
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source)
		go DiffuseOneProc(b.particles[startIndex:endIndex], generator, finished)
	}

}

func DiffuseOneProc(particles []*Particle, generator *(rand.Rand), finished chan bool) {
	for _, p := range particles {
		p.Randstep(generator)
	}
	finished <- true
}

func (p *Particle) Randstep(generator *(rand.Rand)) {
	steplength := p.diffusionRate
	angle := generator.Float64() * 2 * math.Pi
	p.position.x += steplength * math.Cos(angle)
	p.position.y += steplength * math.Sin(angle)
}
