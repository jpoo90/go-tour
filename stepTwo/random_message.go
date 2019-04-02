package stepTwo

import (
	"fmt"
	"math"
	"math/rand"
)

func RandomMessage() {
	fmt.Println("A number is", rand.Intn(10))
	fmt.Printf("You have %g problems. \n", math.Sqrt(7))
}

func Add(x int, y int) int {
	return x + y
}

func Swap(one, two string) (string, string) {
	return two, one
}

func NakedReturn(number float32) (half float32) {
	// half is defined in the return
	half = number / 2
	return
}
