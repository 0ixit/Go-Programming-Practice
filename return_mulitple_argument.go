package main

import ("fmt")


func multipleReturn(s1, s2, s3 string, x int) (string, string, string, in t) {
	return s1, s2, s3, 	x
}


func main(){
	num := 2019
	str1, str2, str3 := "Hello","World!", "It's"
	fmt.Println(multipleReturn(str1, str2, str3, num))
}

