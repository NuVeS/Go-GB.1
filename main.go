package main

import (
	"fmt"
	"math"
)

func rectArea() {
	var a, b float64
	fmt.Print("Enter sides: ")
	fmt.Scan(&a, &b)
	fmt.Printf("%f", a*b)
}

func circleLength() {
	var s float64
	fmt.Print("Enter cicrle area: ")
	fmt.Scan(&s)

	R := math.Sqrt(s / math.Pi)
	len := 2 * math.Pi * R
	fmt.Printf("R = %f, L = %f", R, len)
}

func splitNumber() {
	var n int

	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&n)
		if n > 999 || n < 100 {
			fmt.Println("Wrong number")
		} else {
			var a, b, c int
			a = n % 10
			n /= 10

			b = n % 10
			n /= 10

			c = n

			fmt.Printf("1: %d, 10: %d, 100: %d", a, b, c)
			break
		}
	}

}

func main() {
	var mode int
	fmt.Print(`Choose mode:
1. rectArea
2. circleLength
3. splitNumber
>`)
	fmt.Scan(&mode)
	switch mode {
	case 1:
		rectArea()
	case 2:
		circleLength()
	case 3:
		splitNumber()
	default:
		fmt.Print(`Wrong mode`)
	}
}
