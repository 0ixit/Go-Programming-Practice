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

func (b *batsman) increment_century(temp uint16){
	b.hundred = b.hundred + temp
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

	fmt.Println("before World Cup, Sachin Century is ", sachin.hundred)
	sachin.increment_century(8)
	fmt.Println("After World Cup, Sachin Century is ", sachin.hundred)

}