package main

import (
	"fmt"
	"math/rand"
)

// gives a number 0 through 10
func giveMeNumber() int {
	return rand.Intn(10)
}

func main(){
	fmt.Println(rand.Intn(10))
}

