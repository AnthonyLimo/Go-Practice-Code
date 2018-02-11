package main

import "fmt"

func main() {
	numArray := [5]float64 {1,2,3,4,5} //array declaration

	for i, value := range numArray { //for loop declaration
		fmt.Println(value, i)
	}
}
