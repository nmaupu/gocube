package data

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleAlgString() {
	a := NewAlg("R U R' U'")
	fmt.Println(a)
	// Output:
	// R U R' U'
}

func TestNewAlg(t *testing.T) {
	tables := []struct {
		moves  string
		result []string
	}{
		{"R U R' U'", []string{"R", "U", "R'", "U'"}},
		{"F D F' D'", []string{"F", "D", "F'", "D'"}},
		{"L B L' B'", []string{"L", "B", "L'", "B'"}},
		{"R U R' g", []string{"R", "U", "R'"}},
	}

	for _, table := range tables {
		alg := NewAlg(table.moves)

		// Checking length of slices
		if len(alg.Moves) != len(table.result) {
			t.Fatalf("Incorrect slice, got: %v, want: %v", alg.Moves, table.result)
		}

		// Checking slices
		expected := strings.Join(table.result, " ")
		got := strings.Join(alg.Moves, " ")
		if expected != got {
			t.Errorf("Incorrect slice, expected: %s, got: %s", expected, got)
		}
	}
}

func TestAddMoves(t *testing.T) {
	a := Alg{}
	myAlg := "R U R' U'"
	a.AddMoves(myAlg)
	if len(a.Moves) != 4 {
		t.Fatalf("Moves slice length is incorrect, expected: 4, got: %d", len(a.Moves))
	}

	expected := myAlg
	got := a.String()
	if expected != got {
		t.Fatalf("Moves slice is incorrect, expected: %s, got: %s", expected, got)
	}
}

func ExampleAlgAddMove() {
	a := NewAlg("R U R' U'")
	a.AddMove("F").AddMove("D").AddMove("B")
	fmt.Println(a)
	// Output:
	// R U R' U' F D B
}

func ExampleAlgAddMoveIncorrect() {
	a := NewAlg("R U R' U'")
	a.AddMove("F").AddMove("D").AddMove("g")
	fmt.Println(a)
	// Output:
	// R U R' U' F D
}

func TestReverseMove(t *testing.T) {
	tables := []struct {
		test, result string
	}{
		{"R", "R'"},
		{"R'", "R"},
		{"R2", "R2"},
	}

	for _, table := range tables {
		r := ReverseMove(table.test)

		if r != table.result {
			t.Errorf("Incorrect reverse, expected: %s, got: %s", table.result, r)
		}
	}
}

func ExampleAlgReverse() {
	a := NewAlg("R U R' U' R2")
	fmt.Println(a.Reverse())
	// Output:
	// R2 U R U' R'
}
