package main

import "fmt"

//Defer is stacked.
//Defer waits for the surrounding func to return
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
