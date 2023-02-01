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
		fmt.Println(">", word)
		if(word == "STOP"){
			os.Exit(0)
		}
	}
}
