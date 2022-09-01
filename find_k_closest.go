// Функция ищет k ближайших по значению элементов
// массива к заданному числу

package find_k_closest_el

import "math"

func findClosestElements(arr []int, k int, x int) []int {
	if (x <= arr[0]) || (len(arr) == 1) {
		return (arr[:k])
	} else if x >= arr[len(arr)-1] {
		return (arr[len(arr)-k:])
	}
	l, r := 0, len(arr)-1
	c := 0
	// Находим элемент массива равный x или 2 элемента между которыми находится x (между c и c+1)
	for l+1 < r {
		c = l + (r-l)/2
		if (arr[c] == x) || ((arr[c] < x) && (arr[c+1] > x)) {
			break
		} else if arr[c] > x {
			r = c
		} else {
			l = c + 1
		}
	}
	if l+1 == r {
		c = l
	}
	//Поиск ближайшего числа к x
	if (arr[c] != x) && (math.Abs(float64(arr[c+1]-x)) < math.Abs(float64(arr[c]-x))) {
		c++
	}
	//Определение левого и правого индекса для вывода k ближайших соседей
	l, r = c, c
	for i := 0; i < k-1; i++ {
		if l == 0 {
			return (arr[:k])
		}
		if r == len(arr)-1 {
			return (arr[len(arr)-k:])
		}
		if math.Abs(float64(arr[l-1]-x)) <= math.Abs(float64(arr[r+1]-x)) {
			l--
		} else {
			r++
		}
	}
	return (arr[l : r+1])
}
