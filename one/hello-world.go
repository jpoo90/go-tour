package main

import (
	"fmt"

	"github.com/jpoo90/go-tour/stepTwo"
)

// Declaring variables
// Variables without assignation get a:
// 0 if number , false if bool, "" empty string
var uno, dos, tres bool

// type is infered from assigment
var i, j = 1, 3

// we need to the casting of types manually, otherwise it will throw
var flo float64 = float64(i*j) / 2

// Constants are declared with const, and can't use the := notation
const Euler = 2.71
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Println("hello World")

	stepTwo.RandomMessage()

	fmt.Printf("%d \n", stepTwo.Add(2, 8))

	// this variable assigment is only valid inside functions
	a, b := stepTwo.Swap("Pablo", "Juan")
	fmt.Println(a, b)

	half := stepTwo.NakedReturn(5)
	fmt.Printf("%f \n", half)

	fmt.Println(uno, dos, tres, i, j)
	fmt.Printf("Uno's type %T \n", uno)
	fmt.Printf("%f \n", flo)

}
