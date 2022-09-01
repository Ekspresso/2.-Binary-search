// Функция определяет, является ли заданное число
// идеальным квадратом другого числа

package is_perfect_square

func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		c := left + (right-left)/2
		if c*c == num {
			return (true)
		} else if c*c > num {
			right = c - 1
		} else if c*c < num {
			left = c + 1
		}
	}
	return (false)
}
