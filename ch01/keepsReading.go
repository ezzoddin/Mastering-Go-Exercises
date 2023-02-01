package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(">", scanner.Text())
		if(word == "STOP"){
			os.Exit(0)
		}
	}
}
