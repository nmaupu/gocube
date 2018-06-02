package data

import "fmt"

func ExampleColorString() {
	c := Color{
		Color:     1,
		Name:      "test",
		ShortName: "t",
		Debug:     "t1",
	}

	SetDebug(false)
	fmt.Println(c.String())
	SetDebug(true)
	fmt.Println(c.String())
	// Output:
	// t
	// t1
}
