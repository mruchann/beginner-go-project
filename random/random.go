package main

import (
	"fmt"
	"math/rand"
	"time"
)

// gives a number 0 through 10
func giveMeNumber() int {
	return rand.Intn(10)
}

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println(giveMeNumber())

}

