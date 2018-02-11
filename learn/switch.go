package main

import "fmt"

func main() {
	yourAge := 5

	switch yourAge  {
		case 16 : fmt.Println("Go drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go have fun")
	}
}
