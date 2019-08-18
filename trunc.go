package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Printf("Please enter a decimal:")
	input1 := ""
	fmt.Scanln(&input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {

	}
	fmt.Println(math.Trunc(num1))
	fmt.Println("Enter another decimal: ")
	input2 := ""
	fmt.Scanln(&input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {

	}
	fmt.Println(math.Trunc(num2))

}
