// Функция ищет минимальный элемент в массиве, отсортированнном
// в порядке неубывания и повёрнутом относительно одного из элементов

package find_min_rot_arr

func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] || len(nums) == 1 {
		return (nums[0])
	} else {
		l, r := 0, len(nums)-1
		for l+1 < r {
			c := l + (r-l)/2
			if nums[c] < nums[c-1] {
				return (nums[c])
			} else if nums[c] > nums[r] {
				l = c
			} else if nums[c] < nums[r] {
				r = c
			} else {
				r--
			}
		}
		if nums[l] <= nums[l+1] {
			return (nums[l])
		}
		if nums[r] <= nums[r-1] {
			return (nums[r])
		}
	}
	return (-1)
}
