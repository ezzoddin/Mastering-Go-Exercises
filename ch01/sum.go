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

 	sum := 0.0
	for i := 1; i < len(arguments); i++ {
 		temp, _ := strconv.ParseFloat(arguments[i], 64)
 		sum += temp
	}

	fmt.Println(sum)
}