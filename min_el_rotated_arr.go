// Функция ищет минимальный элемент в повёрнутом
// отсортированном массиве

package min_rot_arr

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return (nums[0])
	} else {
		l, r := 0, len(nums)
		for l < r {
			c := l + (r-l)/2
			if nums[c] < nums[c-1] {
				return (nums[c])
			} else if nums[c] < nums[0] {
				r = c
			} else {
				l = c + 1
			}
		}
		if (l != len(nums)) && (nums[l] < nums[l-1]) {
			return (nums[l])
		}
	}
	return (-1)
}
