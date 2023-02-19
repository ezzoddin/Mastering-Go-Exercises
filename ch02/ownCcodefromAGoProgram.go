package main

// #include <stdio.h>
// void callMyName(){
//    printf("Hello my name is Mobina :) !\n");
// }
import "C"

import "fmt"

func main() {
	fmt.Println("Bellow line call c code from a go program")
	C.callMyName()
 }