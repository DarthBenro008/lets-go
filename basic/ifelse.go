package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func shortHandif(x, i string) string {
	if d := "Hola"; x == d {
		return i
	}
	return "null"
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(shortHandif("Hola", "Zola"))
}
