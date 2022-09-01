// Функция ищет в массиве, отсортированном в порядке неубывания
// символ, который больше заданного символа

package find_smallest_letter

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l+1 < r {
		c := l + (r-l)/2
		if (letters[c] > target) && (letters[c-1] <= target) {
			return (letters[c])
		} else if letters[c] > target && letters[c-1] > target {
			r = c
		} else {
			l = c
		}
	}
	if letters[r] > target && letters[r-1] <= target {
		return (letters[r])
	}
	if l == 0 || r == len(letters)-1 {
		return (letters[0])
	}
	if letters[l] > target && letters[l-1] <= target {
		return (letters[l])
	}
	return (48)
}
