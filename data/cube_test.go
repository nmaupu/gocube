package data

import (
	"fmt"
)

func ExampleCubeString3() {
	c := NewCube(3)
	SetDebug(false)
	fmt.Println(c.String())
	SetDebug(true)
	fmt.Println(c.String())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// O O O|G G G|R R R|B B B|
	// O O O|G G G|R R R|B B B|
	// O O O|G G G|R R R|B B B|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
	//
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeString4() {
	c := NewCube(4)
	SetDebug(false)
	fmt.Println(c.String())
	SetDebug(true)
	fmt.Println(c.String())
	// Output:
	//        |W W W W|
	//        |W W W W|
	//        |W W W W|
	//        |W W W W|
	// O O O O|G G G G|R R R R|B B B B|
	// O O O O|G G G G|R R R R|B B B B|
	// O O O O|G G G G|R R R R|B B B B|
	// O O O O|G G G G|R R R R|B B B B|
	//        |Y Y Y Y|
	//        |Y Y Y Y|
	//        |Y Y Y Y|
	//        |Y Y Y Y|
	//
	//                |W01 W02 W03 W04|
	//                |W05 W06 W07 W08|
	//                |W09 W10 W11 W12|
	//                |W13 W14 W15 W16|
	// O01 O02 O03 O04|G01 G02 G03 G04|R01 R02 R03 R04|B01 B02 B03 B04|
	// O05 O06 O07 O08|G05 G06 G07 G08|R05 R06 R07 R08|B05 B06 B07 B08|
	// O09 O10 O11 O12|G09 G10 G11 G12|R09 R10 R11 R12|B09 B10 B11 B12|
	// O13 O14 O15 O16|G13 G14 G15 G16|R13 R14 R15 R16|B13 B14 B15 B16|
	//                |Y01 Y02 Y03 Y04|
	//                |Y05 Y06 Y07 Y08|
	//                |Y09 Y10 Y11 Y12|
	//                |Y13 Y14 Y15 Y16|
}

func ExampleCubeR() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.R().String())
	// Output:
	//         |W1 W2 G3|
	//         |W4 W5 G6|
	//         |W7 W8 G9|
	// O1 O2 O3|G1 G2 Y3|R7 R4 R1|W9 B2 B3|
	// O4 O5 O6|G4 G5 Y6|R8 R5 R2|W6 B5 B6|
	// O7 O8 O9|G7 G8 Y9|R9 R6 R3|W3 B8 B9|
	//         |Y1 Y2 B7|
	//         |Y4 Y5 B4|
	//         |Y7 Y8 B1|
}

func ExampleCubeRp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Rp().String())
	// Output:
	//         |W1 W2 B7|
	//         |W4 W5 B4|
	//         |W7 W8 B1|
	// O1 O2 O3|G1 G2 W3|R3 R6 R9|Y9 B2 B3|
	// O4 O5 O6|G4 G5 W6|R2 R5 R8|Y6 B5 B6|
	// O7 O8 O9|G7 G8 W9|R1 R4 R7|Y3 B8 B9|
	//         |Y1 Y2 G3|
	//         |Y4 Y5 G6|
	//         |Y7 Y8 G9|
}

func ExampleCubeL() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.L().String())
	// Output:
	//         |B9 W2 W3|
	//         |B6 W5 W6|
	//         |B3 W8 W9|
	// O7 O4 O1|W1 G2 G3|R1 R2 R3|B1 B2 Y7|
	// O8 O5 O2|W4 G5 G6|R4 R5 R6|B4 B5 Y4|
	// O9 O6 O3|W7 G8 G9|R7 R8 R9|B7 B8 Y1|
	//         |G1 Y2 Y3|
	//         |G4 Y5 Y6|
	//         |G7 Y8 Y9|
}

func ExampleCubeLp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Lp().String())
	// Output:
	//         |G1 W2 W3|
	//         |G4 W5 W6|
	//         |G7 W8 W9|
	// O3 O6 O9|Y1 G2 G3|R1 R2 R3|B1 B2 W7|
	// O2 O5 O8|Y4 G5 G6|R4 R5 R6|B4 B5 W4|
	// O1 O4 O7|Y7 G8 G9|R7 R8 R9|B7 B8 W1|
	//         |B9 Y2 Y3|
	//         |B6 Y5 Y6|
	//         |B3 Y8 Y9|
}

func ExampleCubeD() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.D().String())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
	// B7 B8 B9|O7 O8 O9|G7 G8 G9|R7 R8 R9|
	//         |Y7 Y4 Y1|
	//         |Y8 Y5 Y2|
	//         |Y9 Y6 Y3|
}

func ExampleCubeDp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Dp().String())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
	// G7 G8 G9|R7 R8 R9|B7 B8 B9|O7 O8 O9|
	//         |Y3 Y6 Y9|
	//         |Y2 Y5 Y8|
	//         |Y1 Y4 Y7|
}

func ExampleCubeU() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.U().String())
	// Output:
	//         |W7 W4 W1|
	//         |W8 W5 W2|
	//         |W9 W6 W3|
	// G1 G2 G3|R1 R2 R3|B1 B2 B3|O1 O2 O3|
	// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeUp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Up().String())
	// Output:
	//         |W3 W6 W9|
	//         |W2 W5 W8|
	//         |W1 W4 W7|
	// B1 B2 B3|O1 O2 O3|G1 G2 G3|R1 R2 R3|
	// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeF() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.F().String())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |O9 O6 O3|
	// O1 O2 Y1|G7 G4 G1|W7 R2 R3|B1 B2 B3|
	// O4 O5 Y2|G8 G5 G2|W8 R5 R6|B4 B5 B6|
	// O7 O8 Y3|G9 G6 G3|W9 R8 R9|B7 B8 B9|
	//         |R7 R4 R1|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeFp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Fp().String())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |R1 R4 R7|
	// O1 O2 W9|G3 G6 G9|Y3 R2 R3|B1 B2 B3|
	// O4 O5 W8|G2 G5 G8|Y2 R5 R6|B4 B5 B6|
	// O7 O8 W7|G1 G4 G7|Y1 R8 R9|B7 B8 B9|
	//         |O3 O6 O9|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeB() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.B().String())
	// Output:
	//         |R3 R6 R9|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// W3 O2 O3|G1 G2 G3|R1 R2 Y9|B7 B4 B1|
	// W2 O5 O6|G4 G5 G6|R4 R5 Y8|B8 B5 B2|
	// W1 O8 O9|G7 G8 G9|R7 R8 Y7|B9 B6 B3|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |O1 O4 O7|
}

func ExampleCubeBp() {
	c := NewCube(3)
	SetDebug(true)
	fmt.Println(c.Bp().String())
	// Output:
	//         |O7 O4 O1|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// Y7 O2 O3|G1 G2 G3|R1 R2 W1|B3 B6 B9|
	// Y8 O5 O6|G4 G5 G6|R4 R5 W2|B2 B5 B8|
	// Y9 O8 O9|G7 G8 G9|R7 R8 W3|B1 B4 B7|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |R9 R6 R3|
}

func ExampleCubeScramble() {
	c := NewCube(3)
	SetDebug(false)
	fmt.Println(c.U2().F2().Rp().B2().D2().B2().Rp().F2().
		U2().R2().D().Fp().Dp().B().Fp().Rp().Dp().
		U2().Fp().String())
	// Output:
	//      |B G Y|
	//      |B W Y|
	//      |Y O O|
	// W Y R|G W B|Y R G|O O O|
	// W O R|G G Y|G R R|W B G|
	// B B W|B Y W|G B G|W B R|
	//      |R O O|
	//      |W Y O|
	//      |Y R R|
}
