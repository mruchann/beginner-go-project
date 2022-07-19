package main

import (
	"fmt"

	"example.com/Learning-Go/gpa"

	"rsc.io/quote"
)

func main() {
	fmt.Println("My GPA is", gpa.TellMeMyGPA())
	fmt.Println(quote.Go())
}
