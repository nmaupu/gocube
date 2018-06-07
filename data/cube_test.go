package data

import (
	"fmt"
)

// Backup resetted cube
//         |W1 W2 W3|
//         |W4 W5 W6|
//         |W7 W8 W9|
// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
// O4 O5 O6|G4 G5 G6|R4 R5 R6|B4 B5 B6|
// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
//         |Y1 Y2 Y3|
//         |Y4 Y5 Y6|
//         |Y7 Y8 Y9|
//
//      |W W W|
//      |W W W|
//      |W W W|
// O O O|G G G|R R R|B B B|
// O O O|G G G|R R R|B B B|
// O O O|G G G|R R R|B B B|
//      |Y Y Y|
//      |Y Y Y|
//      |Y Y Y|

func ExampleCubeString3() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c)
	SetDebug(true)
	fmt.Println(c)
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
	c := NewCube(4, 0)
	SetDebug(false)
	fmt.Println(c)
	SetDebug(true)
	fmt.Println(c)
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.R())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Rp())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.L())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Lp())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.D())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Dp())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.U())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Up())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.F())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Fp())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.B())
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
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Bp())
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

func ExampleCubeM() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.M())
	// Output:
	//         |W1 B8 W3|
	//         |W4 B5 W6|
	//         |W7 B2 W9|
	// O1 O2 O3|G1 W2 G3|R1 R2 R3|B1 Y8 B3|
	// O4 O5 O6|G4 W5 G6|R4 R5 R6|B4 Y5 B6|
	// O7 O8 O9|G7 W8 G9|R7 R8 R9|B7 Y2 B9|
	//         |Y1 G2 Y3|
	//         |Y4 G5 Y6|
	//         |Y7 G8 Y9|
}

func ExampleCubeMp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Mp())
	// Output:
	//         |W1 G2 W3|
	//         |W4 G5 W6|
	//         |W7 G8 W9|
	// O1 O2 O3|G1 Y2 G3|R1 R2 R3|B1 W8 B3|
	// O4 O5 O6|G4 Y5 G6|R4 R5 R6|B4 W5 B6|
	// O7 O8 O9|G7 Y8 G9|R7 R8 R9|B7 W2 B9|
	//         |Y1 B8 Y3|
	//         |Y4 B5 Y6|
	//         |Y7 B2 Y9|
}

func ExampleCubeM2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.M2())
	// Output:
	//         |W1 Y2 W3|
	//         |W4 Y5 W6|
	//         |W7 Y8 W9|
	// O1 O2 O3|G1 B8 G3|R1 R2 R3|B1 G8 B3|
	// O4 O5 O6|G4 B5 G6|R4 R5 R6|B4 G5 B6|
	// O7 O8 O9|G7 B2 G9|R7 R8 R9|B7 G2 B9|
	//         |Y1 W2 Y3|
	//         |Y4 W5 Y6|
	//         |Y7 W8 Y9|
}

func ExampleCubeS() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.S())
	// Output:
	//         |W1 W2 W3|
	//         |O8 O5 O2|
	//         |W7 W8 W9|
	// O1 Y4 O3|G1 G2 G3|R1 W4 R3|B1 B2 B3|
	// O4 Y5 O6|G4 G5 G6|R4 W5 R6|B4 B5 B6|
	// O7 Y6 O9|G7 G8 G9|R7 W6 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |R8 R5 R2|
	//         |Y7 Y8 Y9|
}

func ExampleCubeSp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Sp())
	// Output:
	//         |W1 W2 W3|
	//         |R2 R5 R8|
	//         |W7 W8 W9|
	// O1 W6 O3|G1 G2 G3|R1 Y6 R3|B1 B2 B3|
	// O4 W5 O6|G4 G5 G6|R4 Y5 R6|B4 B5 B6|
	// O7 W4 O9|G7 G8 G9|R7 Y4 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |O2 O5 O8|
	//         |Y7 Y8 Y9|
}

func ExampleCubeS2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.S2())
	// Output:
	//         |W1 W2 W3|
	//         |Y6 Y5 Y4|
	//         |W7 W8 W9|
	// O1 R8 O3|G1 G2 G3|R1 O8 R3|B1 B2 B3|
	// O4 R5 O6|G4 G5 G6|R4 O5 R6|B4 B5 B6|
	// O7 R2 O9|G7 G8 G9|R7 O2 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |W6 W5 W4|
	//         |Y7 Y8 Y9|
}

func ExampleCubeE() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.E())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// B4 B5 B6|O4 O5 O6|G4 G5 G6|R4 R5 R6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeEp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Ep())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// G4 G5 G6|R4 R5 R6|B4 B5 B6|O4 O5 O6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeE2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.E2())
	// Output:
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
	// O1 O2 O3|G1 G2 G3|R1 R2 R3|B1 B2 B3|
	// R4 R5 R6|B4 B5 B6|O4 O5 O6|G4 G5 G6|
	// O7 O8 O9|G7 G8 G9|R7 R8 R9|B7 B8 B9|
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
}

func ExampleCubeX() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.X())
	// Output:
	//         |G1 G2 G3|
	//         |G4 G5 G6|
	//         |G7 G8 G9|
	// O3 O6 O9|Y1 Y2 Y3|R7 R4 R1|W9 W8 W7|
	// O2 O5 O8|Y4 Y5 Y6|R8 R5 R2|W6 W5 W4|
	// O1 O4 O7|Y7 Y8 Y9|R9 R6 R3|W3 W2 W1|
	//         |B9 B8 B7|
	//         |B6 B5 B4|
	//         |B3 B2 B1|
}

func ExampleCubeXp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Xp())
	// Output:
	//         |B9 B8 B7|
	//         |B6 B5 B4|
	//         |B3 B2 B1|
	// O7 O4 O1|W1 W2 W3|R3 R6 R9|Y9 Y8 Y7|
	// O8 O5 O2|W4 W5 W6|R2 R5 R8|Y6 Y5 Y4|
	// O9 O6 O3|W7 W8 W9|R1 R4 R7|Y3 Y2 Y1|
	//         |G1 G2 G3|
	//         |G4 G5 G6|
	//         |G7 G8 G9|
}

func ExampleCubeX2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.X2())
	// Output:
	//         |Y1 Y2 Y3|
	//         |Y4 Y5 Y6|
	//         |Y7 Y8 Y9|
	// O9 O8 O7|B9 B8 B7|R9 R8 R7|G9 G8 G7|
	// O6 O5 O4|B6 B5 B4|R6 R5 R4|G6 G5 G4|
	// O3 O2 O1|B3 B2 B1|R3 R2 R1|G3 G2 G1|
	//         |W1 W2 W3|
	//         |W4 W5 W6|
	//         |W7 W8 W9|
}

func ExampleCubeY() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Y())
	// Output:
	//         |W7 W4 W1|
	//         |W8 W5 W2|
	//         |W9 W6 W3|
	// G1 G2 G3|R1 R2 R3|B1 B2 B3|O1 O2 O3|
	// G4 G5 G6|R4 R5 R6|B4 B5 B6|O4 O5 O6|
	// G7 G8 G9|R7 R8 R9|B7 B8 B9|O7 O8 O9|
	//         |Y3 Y6 Y9|
	//         |Y2 Y5 Y8|
	//         |Y1 Y4 Y7|
}

func ExampleCubeYp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Yp())
	// Output:
	//         |W3 W6 W9|
	//         |W2 W5 W8|
	//         |W1 W4 W7|
	// B1 B2 B3|O1 O2 O3|G1 G2 G3|R1 R2 R3|
	// B4 B5 B6|O4 O5 O6|G4 G5 G6|R4 R5 R6|
	// B7 B8 B9|O7 O8 O9|G7 G8 G9|R7 R8 R9|
	//         |Y7 Y4 Y1|
	//         |Y8 Y5 Y2|
	//         |Y9 Y6 Y3|
}
func ExampleCubeY2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Y2())
	// Output:
	//         |W9 W8 W7|
	//         |W6 W5 W4|
	//         |W3 W2 W1|
	// R1 R2 R3|B1 B2 B3|O1 O2 O3|G1 G2 G3|
	// R4 R5 R6|B4 B5 B6|O4 O5 O6|G4 G5 G6|
	// R7 R8 R9|B7 B8 B9|O7 O8 O9|G7 G8 G9|
	//         |Y9 Y8 Y7|
	//         |Y6 Y5 Y4|
	//         |Y3 Y2 Y1|
}

func ExampleCubeZ() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Z())
	// Output:
	//         |O7 O4 O1|
	//         |O8 O5 O2|
	//         |O9 O6 O3|
	// Y7 Y4 Y1|G7 G4 G1|W7 W4 W1|B3 B6 B9|
	// Y8 Y5 Y2|G8 G5 G2|W8 W5 W2|B2 B5 B8|
	// Y9 Y6 Y3|G9 G6 G3|W9 W6 W3|B1 B4 B7|
	//         |R7 R4 R1|
	//         |R8 R5 R2|
	//         |R9 R6 R3|
}

func ExampleCubeZp() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Zp())
	// Output:
	//         |R3 R6 R9|
	//         |R2 R5 R8|
	//         |R1 R4 R7|
	// W3 W6 W9|G3 G6 G9|Y3 Y6 Y9|B7 B4 B1|
	// W2 W5 W8|G2 G5 G8|Y2 Y5 Y8|B8 B5 B2|
	// W1 W4 W7|G1 G4 G7|Y1 Y4 Y7|B9 B6 B3|
	//         |O3 O6 O9|
	//         |O2 O5 O8|
	//         |O1 O4 O7|
}

func ExampleCubeZ2() {
	c := NewCube(3, 0)
	SetDebug(true)
	fmt.Println(c.Z2())
	// Output:
	//         |Y9 Y8 Y7|
	//         |Y6 Y5 Y4|
	//         |Y3 Y2 Y1|
	// R9 R8 R7|G9 G8 G7|O9 O8 O7|B9 B8 B7|
	// R6 R5 R4|G6 G5 G4|O6 O5 O4|B6 B5 B4|
	// R3 R2 R1|G3 G2 G1|O3 O2 O1|B3 B2 B1|
	//         |W9 W8 W7|
	//         |W6 W5 W4|
	//         |W3 W2 W1|
}

// No need to test using debug mode for wide moves
// because those moves are made using
// other basic moves that are already well tested
func ExampleCubeRw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Rw())
	// Output:
	//      |W G G|
	//      |W G G|
	//      |W G G|
	// O O O|G Y Y|R R R|W W B|
	// O O O|G Y Y|R R R|W W B|
	// O O O|G Y Y|R R R|W W B|
	//      |Y B B|
	//      |Y B B|
	//      |Y B B|
}

func ExampleCubeRwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Rwp())
	// Output:
	//      |W B B|
	//      |W B B|
	//      |W B B|
	// O O O|G W W|R R R|Y Y B|
	// O O O|G W W|R R R|Y Y B|
	// O O O|G W W|R R R|Y Y B|
	//      |Y G G|
	//      |Y G G|
	//      |Y G G|
}

func ExampleCubeRw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Rw2())
	// Output:
	//      |W Y Y|
	//      |W Y Y|
	//      |W Y Y|
	// O O O|G B B|R R R|G G B|
	// O O O|G B B|R R R|G G B|
	// O O O|G B B|R R R|G G B|
	//      |Y W W|
	//      |Y W W|
	//      |Y W W|
}

func ExampleCubeLw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Lw())
	// Output:
	//      |B B W|
	//      |B B W|
	//      |B B W|
	// O O O|W W G|R R R|B Y Y|
	// O O O|W W G|R R R|B Y Y|
	// O O O|W W G|R R R|B Y Y|
	//      |G G Y|
	//      |G G Y|
	//      |G G Y|
}

func ExampleCubeLwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Lwp())
	// Output:
	//      |G G W|
	//      |G G W|
	//      |G G W|
	// O O O|Y Y G|R R R|B W W|
	// O O O|Y Y G|R R R|B W W|
	// O O O|Y Y G|R R R|B W W|
	//      |B B Y|
	//      |B B Y|
	//      |B B Y|
}

func ExampleCubeLw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Lw2())
	// Output:
	//      |Y Y W|
	//      |Y Y W|
	//      |Y Y W|
	// O O O|B B G|R R R|B G G|
	// O O O|B B G|R R R|B G G|
	// O O O|B B G|R R R|B G G|
	//      |W W Y|
	//      |W W Y|
	//      |W W Y|
}

func ExampleCubeUw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Uw())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// G G G|R R R|B B B|O O O|
	// G G G|R R R|B B B|O O O|
	// O O O|G G G|R R R|B B B|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeUwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Uwp())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// B B B|O O O|G G G|R R R|
	// B B B|O O O|G G G|R R R|
	// O O O|G G G|R R R|B B B|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeUw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Uw2())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// R R R|B B B|O O O|G G G|
	// R R R|B B B|O O O|G G G|
	// O O O|G G G|R R R|B B B|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeDw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Dw())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// O O O|G G G|R R R|B B B|
	// B B B|O O O|G G G|R R R|
	// B B B|O O O|G G G|R R R|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeDwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Dwp())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// O O O|G G G|R R R|B B B|
	// G G G|R R R|B B B|O O O|
	// G G G|R R R|B B B|O O O|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeDw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Dw2())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// O O O|G G G|R R R|B B B|
	// R R R|B B B|O O O|G G G|
	// R R R|B B B|O O O|G G G|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeFw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Fw())
	// Output:
	//      |W W W|
	//      |O O O|
	//      |O O O|
	// O Y Y|G G G|W W R|B B B|
	// O Y Y|G G G|W W R|B B B|
	// O Y Y|G G G|W W R|B B B|
	//      |R R R|
	//      |R R R|
	//      |Y Y Y|
}

func ExampleCubeFwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Fwp())
	// Output:
	//      |W W W|
	//      |R R R|
	//      |R R R|
	// O W W|G G G|Y Y R|B B B|
	// O W W|G G G|Y Y R|B B B|
	// O W W|G G G|Y Y R|B B B|
	//      |O O O|
	//      |O O O|
	//      |Y Y Y|
}

func ExampleCubeFw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Fw2())
	// Output:
	//      |W W W|
	//      |Y Y Y|
	//      |Y Y Y|
	// O R R|G G G|O O R|B B B|
	// O R R|G G G|O O R|B B B|
	// O R R|G G G|O O R|B B B|
	//      |W W W|
	//      |W W W|
	//      |Y Y Y|
}

func ExampleCubeBw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Bw())
	// Output:
	//      |R R R|
	//      |R R R|
	//      |W W W|
	// W W O|G G G|R Y Y|B B B|
	// W W O|G G G|R Y Y|B B B|
	// W W O|G G G|R Y Y|B B B|
	//      |Y Y Y|
	//      |O O O|
	//      |O O O|
}

func ExampleCubeBwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Bwp())
	// Output:
	//      |O O O|
	//      |O O O|
	//      |W W W|
	// Y Y O|G G G|R W W|B B B|
	// Y Y O|G G G|R W W|B B B|
	// Y Y O|G G G|R W W|B B B|
	//      |Y Y Y|
	//      |R R R|
	//      |R R R|
}

func ExampleCubeBw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Bw2())
	// Output:
	//      |Y Y Y|
	//      |Y Y Y|
	//      |W W W|
	// R R O|G G G|R O O|B B B|
	// R R O|G G G|R O O|B B B|
	// R R O|G G G|R O O|B B B|
	//      |Y Y Y|
	//      |W W W|
	//      |W W W|
}

func ExampleCubeMw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Mw())
	// Output:
	//      |B W B|
	//      |B W B|
	//      |B W B|
	// O O O|W G W|R R R|Y B Y|
	// O O O|W G W|R R R|Y B Y|
	// O O O|W G W|R R R|Y B Y|
	//      |G Y G|
	//      |G Y G|
	//      |G Y G|
}

func ExampleCubeMwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Mwp())
	// Output:
	//      |G W G|
	//      |G W G|
	//      |G W G|
	// O O O|Y G Y|R R R|W B W|
	// O O O|Y G Y|R R R|W B W|
	// O O O|Y G Y|R R R|W B W|
	//      |B Y B|
	//      |B Y B|
	//      |B Y B|
}

func ExampleCubeMw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Mw2())
	// Output:
	//      |Y W Y|
	//      |Y W Y|
	//      |Y W Y|
	// O O O|B G B|R R R|G B G|
	// O O O|B G B|R R R|G B G|
	// O O O|B G B|R R R|G B G|
	//      |W Y W|
	//      |W Y W|
	//      |W Y W|
}

func ExampleCubeSw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Sw())
	// Output:
	//      |O O O|
	//      |W W W|
	//      |O O O|
	// Y O Y|G G G|W R W|B B B|
	// Y O Y|G G G|W R W|B B B|
	// Y O Y|G G G|W R W|B B B|
	//      |R R R|
	//      |Y Y Y|
	//      |R R R|
}

func ExampleCubeSwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Swp())
	// Output:
	//      |R R R|
	//      |W W W|
	//      |R R R|
	// W O W|G G G|Y R Y|B B B|
	// W O W|G G G|Y R Y|B B B|
	// W O W|G G G|Y R Y|B B B|
	//      |O O O|
	//      |Y Y Y|
	//      |O O O|
}

func ExampleCubeSw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Sw2())
	// Output:
	//      |Y Y Y|
	//      |W W W|
	//      |Y Y Y|
	// R O R|G G G|O R O|B B B|
	// R O R|G G G|O R O|B B B|
	// R O R|G G G|O R O|B B B|
	//      |W W W|
	//      |Y Y Y|
	//      |W W W|
}

func ExampleCubeEw() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Ew())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// B B B|O O O|G G G|R R R|
	// O O O|G G G|R R R|B B B|
	// B B B|O O O|G G G|R R R|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeEwp() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Ewp())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// G G G|R R R|B B B|O O O|
	// O O O|G G G|R R R|B B B|
	// G G G|R R R|B B B|O O O|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeEw2() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.Ew2())
	// Output:
	//      |W W W|
	//      |W W W|
	//      |W W W|
	// R R R|B B B|O O O|G G G|
	// O O O|G G G|R R R|B B B|
	// R R R|B B B|O O O|G G G|
	//      |Y Y Y|
	//      |Y Y Y|
	//      |Y Y Y|
}

func ExampleCubeScramble() {
	c := NewCube(3, 0)
	SetDebug(false)
	fmt.Println(c.U2().F2().Rp().B2().D2().B2().Rp().F2().
		U2().R2().D().Fp().Dp().B().Fp().Rp().Dp().
		U2().Fp().L2())
	// Output:
	//      |R G Y|
	//      |W W Y|
	//      |Y O O|
	// W B B|R W B|Y R G|O O B|
	// R O W|G G Y|G R R|W B G|
	// R Y W|O Y W|G B G|W B G|
	//      |B O O|
	//      |B Y O|
	//      |Y R R|
}

func ExampleCubeExecute() {
	c := NewCube(3, 0)
	SetDebug(false)

	a := NewAlg("U2 F2 R' B2 D2 B2 R' F2 U2 R2 D F' D' B F' R' D' U2 F' L2")

	fmt.Println(c.Execute(a))
	// Output:
	//      |R G Y|
	//      |W W Y|
	//      |Y O O|
	// W B B|R W B|Y R G|O O B|
	// R O W|G G Y|G R R|W B G|
	// R Y W|O Y W|G B G|W B G|
	//      |B O O|
	//      |B Y O|
	//      |Y R R|
}
