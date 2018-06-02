package data

import (
	"fmt"
	"testing"
)

func TestNewFace(t *testing.T) {
	n := 3
	f := NewFace(n, Colors["white"])

	// Checking number of rows
	if len(f.Colors) != n {
		t.Errorf("Number of rows is incorrect, got: %d, expected: %d.", len(f.Colors), n)
	}

	// Checking that each row has the right number of cols
	for i := 0; i < len(f.Colors); i++ {
		if len(f.Colors[i]) != n {
			t.Errorf("Number of cols is incorrect for line %d, got: %d, expected: %d.", i, len(f.Colors[i]), n)
		}
	}

	// Checking ShortName and Debug values
	tables := []struct {
		i int
		j int
		s string
		d string
	}{
		{0, 0, "W", "W1"},
		{0, 1, "W", "W2"},
		{0, 2, "W", "W3"},
		{1, 0, "W", "W4"},
		{1, 1, "W", "W5"},
		{1, 2, "W", "W6"},
		{2, 0, "W", "W7"},
		{2, 1, "W", "W8"},
		{2, 2, "W", "W9"},
	}

	for _, table := range tables {
		toTest := f.Colors[table.i][table.j].ShortName
		if toTest != table.s {
			t.Errorf("ShortName is incorrect for row=%d, col=%d. Expected: %s, got: %s", table.i, table.j, table.s, toTest)
		}

		toTest = f.Colors[table.i][table.j].Debug
		if toTest != table.d {
			t.Errorf("Debug is incorrect for row=%d, col=%d. Expected: %s, got: %s", table.i, table.j, table.d, toTest)
		}
	}
}

func TestGetCenterColor(t *testing.T) {
	f := NewFace(3, Colors["white"])
	f.Colors[1][1].ShortName = "T"

	c, e := f.GetCenterColor()
	if e != nil {
		t.Errorf("When nb of faces is odd, GetCenterColor shall not fail")
	}
	if c.ShortName != "T" {
		t.Errorf("Center is incorrect, expected: T, got: %s", c.ShortName)
	}

	f = NewFace(2, Colors["white"])
	c, e = f.GetCenterColor()
	if e == nil {
		t.Errorf("When nb of faces is even, GetCenterColor shall fail")
	}
}

func ExampleFaceString() {
	f := NewFace(3, Colors["white"])
	SetDebug(true)
	fmt.Println(f.String())

	f = NewFace(4, Colors["white"])
	SetDebug(false)
	fmt.Println(f.String())
	SetDebug(true)
	fmt.Println(f.String())
	// Output:
	// W1 W2 W3|
	// W4 W5 W6|
	// W7 W8 W9|
	//
	//W W W W|
	//W W W W|
	//W W W W|
	//W W W W|
	//
	//W01 W02 W03 W04|
	//W05 W06 W07 W08|
	//W09 W10 W11 W12|
	//W13 W14 W15 W16|
}

func TestFaceCopy(t *testing.T) {
	f1 := NewFace(3, Colors["white"])
	f2 := f1.Copy()
	f2.Colors[0][0] = Colors["yellow"]
	if f1.Colors[0][0].Name == f2.Colors[0][0].Name {
		t.Errorf("Copy is incorrect")
	}
}

func ExampleFaceFlipVertical() {
	f := NewFace(4, Colors["white"])
	SetDebug(true)
	fmt.Println(f.FlipVertical().String())
	// Output:
	//W04 W03 W02 W01|
	//W08 W07 W06 W05|
	//W12 W11 W10 W09|
	//W16 W15 W14 W13|
}

func ExampleFaceFlipHorizontal() {
	f := NewFace(4, Colors["white"])
	SetDebug(true)
	fmt.Println(f.FlipHorizontal())
	// Output:
	//W13 W14 W15 W16|
	//W09 W10 W11 W12|
	//W05 W06 W07 W08|
	//W01 W02 W03 W04|
}

func ExampleFaceFlip() {
	f := NewFace(4, Colors["white"])
	SetDebug(true)
	fmt.Println(f.Flip())
	// Output:
	//W16 W15 W14 W13|
	//W12 W11 W10 W09|
	//W08 W07 W06 W05|
	//W04 W03 W02 W01|
}

func ExampleFaceReplaceCol() {
	f1 := NewFace(4, Colors["white"])
	f2 := NewFace(4, Colors["yellow"])
	SetDebug(true)
	fmt.Println(f1.ReplaceCol(*f2, 2).String())
	// Output:
	//W01 W02 Y03 W04|
	//W05 W06 Y07 W08|
	//W09 W10 Y11 W12|
	//W13 W14 Y15 W16|
}

func ExampleFaceReplaceRow() {
	f1 := NewFace(4, Colors["white"])
	f2 := NewFace(4, Colors["yellow"])
	SetDebug(true)
	fmt.Println(f1.ReplaceRow(*f2, 2).String())
	// Output:
	//W01 W02 W03 W04|
	//W05 W06 W07 W08|
	//Y09 Y10 Y11 Y12|
	//W13 W14 W15 W16|
}

func ExampleFaceRotateClockwise() {
	f := NewFace(4, Colors["white"])
	SetDebug(true)
	fmt.Println(f.RotateClockwise().String())
	// Output:
	//W13 W09 W05 W01|
	//W14 W10 W06 W02|
	//W15 W11 W07 W03|
	//W16 W12 W08 W04|
}

func ExampleFaceRotateAntiClockwise() {
	f := NewFace(4, Colors["white"])
	SetDebug(true)
	fmt.Println(f.RotateAntiClockwise().String())
	// Output:
	//W04 W08 W12 W16|
	//W03 W07 W11 W15|
	//W02 W06 W10 W14|
	//W01 W05 W09 W13|
}
