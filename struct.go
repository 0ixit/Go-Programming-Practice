package main

import "fmt"

type batsman struct{
	match uint16
	run uint16
	hundred uint16
	fifty uint16
	avg float32
}

func main(){
	sachin := batsman{
			match : 367,
			run : 14626,
			hundred : 45,
			fifty : 86,
			avg :0,
	}
	sehwag := batsman{267, 8973, 15, 46,0}
	sachin.avg = (float32)(sachin.run)/(float32)(sachin.match)
	sehwag.avg = (float32)(sehwag.run)/(float32)(sehwag.match)
	fmt.Println("sachin avg is ", sachin.avg)
	fmt.Println("sachin avg is ", sehwag.avg)
}