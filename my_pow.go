// Функция возводит заданное число в заданную степень

package my_pow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return (1)
	}
	tmp := myPow(x, n/2)
	if (n < 0) && (tmp > 1) && x > 1 {
		if n%2 == 0 {
			return (1 / (tmp * tmp))
		}
		return (1 / (tmp * tmp * x))
	}
	if n%2 == 0 {
		return (tmp * tmp)
	}
	if n < 0 {
		return (tmp * tmp * (1 / x))
	} else {
		return (tmp * tmp * x)
	}
}
