package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("Rich Boi , OS X")
	case "linux":
		fmt.Print("Linux :)")
	default:
		fmt.Print("Windows?")
	}
}
