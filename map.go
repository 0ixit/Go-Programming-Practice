package main

import "fmt"

func main() {
	 PM_Coutry := make(map[string]string)
	 PM_Coutry["India"] = "Narendra Modi"
	 PM_Coutry["U.S.A"] = "Donald Trump"
	 PM_Coutry["Russia"] = "Vladimir Putin"
	 PM_Coutry["China"] = "Xi jinping"
	 PM_Coutry["Pakistan"] = "Imran Khan"
	 
	 PM_Coutry["fake"] = "fake PM"
	 delete(PM_Coutry, "fake")

	 for k, v := range PM_Coutry {
	 	fmt.Println(k, " : ", v)
	 }
}