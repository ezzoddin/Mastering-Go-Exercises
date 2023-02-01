package main

import (
	"fmt"
	"os"
	"strconv"
 )

func main() { 
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	var sum float64
	count := len(arguments)
	for i := 1; i < count; i++ {
 		number, _ := strconv.ParseFloat(arguments[i], 64)
 		sum += number
	}

	average := sum / float64(count-1) 
	fmt.Println(average)
}