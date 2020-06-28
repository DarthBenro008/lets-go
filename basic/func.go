package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//Check how typecasting for each is omitted
func multiply(x, y int) int {
	return x * y
}

//You can return more than one result in a func
func swap(x, y string) (string, string) {
	return y, x
}

//Naked Return : Returns whatever you have defined.
func split(sum int) (x, y int) {
	x = sum * 3
	y = sum - x
	return
}

func main() {
	fmt.Println(add(21, 12))
	fmt.Println(multiply(4, 2))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(18))
}
