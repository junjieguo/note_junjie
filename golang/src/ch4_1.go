package main

import (
	"fmt"
	"math"
)

func main() {
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("%d,%d\n", Sunday, Friday)

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	fmt.Printf("%f\n%E\n%G\n", x, y, z)

	a := [3]int{4, 6, 7}
	for i, v := range a {
		fmt.Printf("%d, %d\n", i, v)
	}
	aaa := []int{1, 2, 3, 4}
	bbb := []int{1, 2, 3, 4}
	ccc := []int{1, 23, 4, 5, 6}
	fmt.Printf("dif aaa  bbb %t\n", equal_slice(aaa, bbb))
	fmt.Printf("dif aaa  ccc %t\n", equal_slice(aaa, ccc))

	var xxx, yyy []int
	for i := 0; i < 10; i++ {
		yyy = appendInt(xxx, i)
		fmt.Printf("%d cap=%d\t %v\n", i, cap(yyy), yyy)
		xxx = yyy
	}

	data := []string{"one", "", "three", "four", ""}
	fmt.Printf("data %q\n", data)
	fmt.Printf("nonempey data %q\n", nonempty(data))
	fmt.Printf("data %q\n", data)

	data1 := []string{"one", "good", "goodbye", "", "OK", ""}
	fmt.Printf("data1 %q\n", data1)

	data2 := []string{"one", "one", "two", "three", "three", "three", "four"}
	fmt.Printf("data2 %q\n", data2)
	fmt.Printf("data2 %q\n", removexl(data2))

	data3 := []string{"one"}
	fmt.Printf("data3 %q\n", data3)
	fmt.Printf("data3 %q\n", removexl(data3))

	ages := make(map[string]int)
	ages["sky"] = 31
	ages["dodg"] = 28
	fmt.Printf(" sky age %v\n", ages["sky"])
	ages["bobm"]++

	for name, age := range ages {
		fmt.Printf("%s age is %d\n", name, age)
	}
}

func equal_slice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}

	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func removexl(strings []string) []string {
	i := 1

	if len(strings) <= 1 {
		return strings
	} else {
		for _, vv := range strings {
			if strings[i-1] != vv {
				strings[i] = vv
				i++
			}
		}
	}
	return strings[:i]

}
