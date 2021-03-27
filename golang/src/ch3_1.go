package main

import (
	"fmt"
	"math"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	var big uint64 = 0
	fmt.Println(big<<2, big-1)
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
	f := -5.66
	fmt.Println(int(f))

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	for x := 0; x < 10; x++ {
		fmt.Printf("x = %d e^%d=%8.4f\n", x, x, math.Exp(float64(x)))
	}
}
