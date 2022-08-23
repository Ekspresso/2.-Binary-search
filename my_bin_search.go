//Программа принимает из 1 аргумента командной строки
//целое число, которое необходимо найти и выводит его номер
//в заданном массиве
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s [10]int
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: more than 1 search element")
	} else if len(os.Args) < 2 {
		fmt.Println("Error: the search element is not specified")
	} else {
		el, err := strconv.Atoi(os.Args[1])
		if err == nil {
			c := bin_search(el, s[:])
			if c == 1 {
				fmt.Println("Element not found or incorrect array")
			} else {
				fmt.Println(c)
			}
		}
	}
}

//Функция бинарного поиска получает элемент, который надо найти,
//срез массива для поиска в нём и возвращает индекс найденного элемента
func bin_search(el int, arr []int) int {
	first := 0
	n := len(arr)
	if arr[first] < arr[n-1] {
		for i := 0; i <= len(arr)/2; i++ {
			if el < arr[first+(n-first)/2] {
				n = first + (n-first)/2
			} else if el > arr[first+(n-first)/2] {
				first = first + (n-first)/2
			} else {
				return (first + (n-first)/2)
			}
		}
	} else {
		for i := 0; i <= len(arr)/2; i++ {
			if el < arr[first+(n-first)/2] {
				first = first + (n-first)/2
			} else if el > arr[first+(n-first)/2] {
				n = first + (n-first)/2
			} else {
				return (first + (n-first)/2)
			}
		}
	}
	return (1)
}
