package main

import ("fmt")

func add(a, b float64) float64{
	return a + b
}

func main(){
	num1,num2 := 5.4, 8.7
	fmt.Println(add(num1, num2))
}
