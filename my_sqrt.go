// Функция вычисляет целую часть корня из числа
package main

import (
	"fmt"
	"os"
	"strconv"
)

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		c := left + (right-left)/2
		if (c*c == x) || (((c+1)*(c+1) > x) && (c*c < x)) {
			return (c)
		} else if c*c > x {
			right = c - 1
		} else if c*c < x {
			left = c + 1
		}
	}
	return (0)
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Error: more than 1 search element")
	} else if len(os.Args) < 2 {
		fmt.Println("Error: the search element is not specified")
	} else {
		n, err := strconv.Atoi(os.Args[1])
		if (err == nil) && (n >= 0) {
			fmt.Println(mySqrt(n))
		}
	}
}
