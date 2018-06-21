package data

import (
	"strings"
)

type Alg struct {
	Moves []string
}

func NewAlg(moves string) *Alg {
	alg := Alg{}

	alg.AddMoves(moves)

	return &alg
}

func (a Alg) Copy() *Alg {
	return NewAlg(a.String())
}

func (a *Alg) AddMove(m string) *Alg {
	switch m {
	// Simply ignoring not recognized move
	case
		"R", "R'", "R2", "r", "r'", "r2", "r2'", "R2'",
		"L", "L'", "L2", "l", "l'", "l2", "l2'", "L2'",
		"D", "D'", "D2", "d", "d'", "d2", "d2'", "D2'",
		"U", "U'", "U2", "u", "u'", "u2", "u2'", "U2'",
		"B", "B'", "B2", "b", "b'", "b2", "b2'", "B2'",
		"F", "F'", "F2", "f", "f'", "f2", "f2'", "F2'",
		"M", "M'", "M2", "m", "m'", "m2", "m2'", "M2'",
		"S", "S'", "S2", "s", "s'", "s2", "s2'", "S2'",
		"E", "E'", "E2", "e", "e'", "e2", "e2'", "E2'",
		"x", "x'", "x2", "x2'",
		"y", "y'", "y2", "y2'",
		"z", "z'", "z2", "z2'":
		a.Moves = append(a.Moves, m)
	}

	return a
}

func (a *Alg) AddMoves(moves string) *Alg {
	// Replace with space to avoid problems when triggers are stuck together.
	// E.g. (R U R' U')(R U R' U') instead of (R U R' U') (R U R' U')
	m := strings.Replace(moves, "(", " ", -1)
	m = strings.Replace(m, ")", " ", -1)
	mo := strings.Split(m, " ")

	for _, m := range mo {
		a.AddMove(m)
	}

	return a
}

func (a *Alg) String() string {
	return strings.Join(a.Moves, " ")
}

func ReverseMove(m string) string {
	if strings.Contains(m, "2") {
		return m
	} else if strings.Contains(m, "'") {
		return string(m[0])
	} else {
		return m + "'"
	}
}

func (a *Alg) Reverse() *Alg {
	copy := a.Copy()
	a.Moves = make([]string, 0)
	for _, m := range copy.Moves {
		// Prepend to result slice
		a.Moves = append([]string{ReverseMove(m)}, a.Moves...)
	}

	return a
}
