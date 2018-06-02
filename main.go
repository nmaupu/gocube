package main

import (
	"fmt"
	"github.com/nmaupu/gocube/data"
)

func main() {
	fmt.Println("In development")
	c := data.NewCube(3)
	fmt.Println(c.String())
}
