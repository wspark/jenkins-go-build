package main

import "fmt"

func main() {
	var s1 string = "hello world go!!"
	var s2 string = "한글"
	var V1 bool = true
	var V2 bool = false
//	var V3 bool = false

	fmt.Println("s1 = ", len(s1))
	fmt.Println("s2 = ", len(s2))
	fmt.Println(s1 == s2)

	fmt.Println(V1)
	fmt.Println(V2)

}
