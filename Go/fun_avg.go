 package main

import "fmt"

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	x := 10.25
	y := 9.75

	result := avg(x, y)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)
}
