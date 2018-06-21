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
	myAlg := "(R U R' U')(L B L' B')"
	a.AddMoves(myAlg)
	if len(a.Moves) != 8 {
		t.Fatalf("Moves slice length is incorrect, expected: 8, got: %d", len(a.Moves))
	}

	expected := "R U R' U' L B L' B'"
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
		{"r2'", "r2'"},
	}

	for _, table := range tables {
		r := ReverseMove(table.test)

		if r != table.result {
			t.Errorf("Incorrect reverse, expected: %s, got: %s", table.result, r)
		}
	}
}

func TestReverse(t *testing.T) {
	a := NewAlg("R U R' U' R2")
	r := a.Copy().Reverse()

	if len(a.Moves) != len(r.Moves) {
		t.Errorf("Incorrect reverse, lengths differ. Expected: %d, got: %d", len(a.Moves), len(r.Moves))
	}

	if r.String() != "R2 U R U' R'" {
		t.Errorf("Incorrect reverse, expected: R2 U R U' R', got: %s", r)
	}

	// Second test
	a = NewAlg("r' U (r2 U' r2' U' r2) U r'")
	r = a.Copy().Reverse()
	expected := "r U' r2 U r2' U r2 U' r"
	if r.String() != expected {
		t.Errorf("Incorrect reverse, expected: %s, got: %s", expected, r)
	}
}
