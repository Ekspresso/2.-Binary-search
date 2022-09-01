// Функция ищет начальный и конечный индекс числа
// в массиве, остортированном в порядке не убывания

package search_for_a_range

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	ret := [...]int{-1, -1}
	for (l + 1) < r {
		c := l + (r-l)/2
		if (nums[c] == target) && (nums[c-1] < target) {
			ret[0] = c
			break
		} else if nums[c] < target {
			l = c
		} else {
			r = c
		}
	}
	if (nums[l] == target) && (ret[0] == -1) {
		ret[0] = l
	} else if (nums[r] == target) && (nums[l] != target) && (ret[0] == -1) {
		ret[0] = r
	} else if ret[0] == -1 {
		return (ret[:])
	}
	l, r = 0, len(nums)-1
	for (l + 1) < r {
		c := l + (r-l)/2
		if (nums[c] == target) && (nums[c+1] > target) {
			ret[1] = c
			break
		} else if nums[c] > target {
			r = c
		} else {
			l = c
		}
	}
	if (nums[r] == target) && (ret[1] == -1) {
		ret[1] = r
	} else if (nums[l] == target) && (nums[r] != target) && (ret[1] == -1) {
		ret[1] = l
	}
	return (ret[:])
}
