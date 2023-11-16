 package main

import "fmt"

func main() {
	
	var a int64 = 3
	var b int = int(a)  

	var c float64 = 2.2

	var result = float64(b) + c  

	fmt.Println(result)
}