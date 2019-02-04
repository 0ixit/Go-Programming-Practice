package main

import ("fmt"
		"math/rand")


func main(){
	foo()  
}


func foo(){
	  fmt.Println("Random Number Generated between 0 - 99 ", rand.Intn(100))
}
