package main

import "fmt"

func standardForLoop() {
	sum := 0
	//Similar to Java
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

func whileLoop() {
	sum := 1
	for sum < 100 {
		sum++
	}
	fmt.Println(sum)
}

func infinity() {

	for sum := 0; ; {
		sum++
		fmt.Println(sum, "GO is the best!")
	}
}

func main() {
	standardForLoop()
	whileLoop()
	// infinity()

}
