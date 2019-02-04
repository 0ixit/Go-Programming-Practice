package main

import "fmt"

type batsman struct{
	match uint16
	run uint16
	ball uint16
	hundred uint16
	fifty uint16
	avg float32
}


func (b batsman) strike_rate() float32 {
	return 100*float32(b.run)/float32(b.ball)
}

func main(){
	sachin := batsman{
			match : 367,
			run : 14626,
			ball : 12526,
			hundred : 45,
			fifty : 86,
			avg :0,
	}
	sehwag := batsman{267, 8973, 6142, 15, 46,0}
	sachin.avg = (float32)(sachin.run)/(float32)(sachin.match)
	sehwag.avg = (float32)(sehwag.run)/(float32)(sehwag.match)
	
	fmt.Println("sachin strike rate is ", sachin.strike_rate())
	fmt.Println("sehwag strike rate is ", sehwag.strike_rate())
}