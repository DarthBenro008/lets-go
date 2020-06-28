package main

import "fmt"

//Factored Variable Declarations
var (
	first  = 1
	second = true
	third  = "GO is the best"
)

//intialise variable and then type at last
var c, python bool

//Init variables
var i, j int = 1, 2

//Multi Types - No need of typedef
var hello, mello, kello = false, "Konichuwa!", 4

//Const , cannot use :=
const Pi = 3.141

func main() {
	//ShortHand , epic way to declare variables, only inside a func
	k := 3
	fmt.Println(c, python, i, j)
	fmt.Println(hello, mello, kello)
	fmt.Println(k)
	fmt.Println(first, second, third)
	fmt.Println("PI?", Pi)
}
