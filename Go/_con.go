 package main

import "fmt"

func main() {
	var myInt int = 105
	var myUint uint = uint(myInt)
	var myFloat float64 = float64(myInt)

	fmt.Println(myInt, myUint, myFloat)
}