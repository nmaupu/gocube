package compute

import (
	"github.com/nmaupu/gocube/data"
	"math/rand"
	"time"
)

var (
	Axis = map[string][][]string{
		"333": {
			{"R", "L"},
			{"U", "D"},
			{"F", "B"},
		},
	}
	Suffixes = map[string][]string{
		"333": {"", "'", "2"},
	}
)

type Generator struct {
	r *rand.Rand
}

func NewGenerator() *Generator {
	g := Generator{}
	seed := rand.NewSource(time.Now().UnixNano())
	g.r = rand.New(seed)
	return &g
}

func (g Generator) generateInt(max int) int {
	r := g.r.Intn(max)
	return r
}

func (g Generator) GenerateAlg(length int) *data.Alg {
	axis := Axis["333"]
	suffixes := Suffixes["333"]

	var (
		lastMoves []int
		finished  bool
	)

	alg := data.Alg{}
	lastAxis := -1
	lenAxis := len(axis)
	lenSuffixes := len(suffixes)
	for i := 0; i < length; i++ {
		finished = false

		for !finished {
			row := g.generateInt(lenAxis)

			if row != lastAxis {
				lastMoves = make([]int, len(axis[row]))
				lastAxis = row
			}

			col := g.generateInt(len(axis[row]))
			suffix := suffixes[g.generateInt(lenSuffixes)]

			if lastMoves[col] == 0 {
				lastMoves[col] = 1
				alg.AddMove(axis[row][col] + suffix)
				finished = true
			}
		}
	}

	return &alg
}
