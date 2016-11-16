package main

import (
	"math/rand"
	"time"
)

type Dice struct {
	exemptions []string
}

func NewDice(exemptions []string) *Dice {
	d := &Dice{}
	d.exemptions = exemptions

	return d
}

func (d *Dice) Throw(candidates []string) string {
	var act []string
Loop:
	for _, c := range candidates {
		for _, ex := range d.exemptions {
			if ex == c {
				continue Loop
			}
		}
		act = append(act, c)
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(act))

	return act[i]
}
