// Функция бинарного поиска заданного элемента target
// в массиве, который перевернули (переставили местами
// левую и правую часть) на элементе с индексом k.
// В начале идёт поиск индекса переворота, затем
// в левой и правой части по отдельности ищется нужный элемент.

package my_bin_search

func search(nums []int, target int) int {
	k := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			k = i + 1
		}
	}
	if k == -1 {
		return (search_in_slice(nums, target))
	} else {
		res1 := search_in_slice(nums[:k], target)
		if res1 == -1 {
			res2 := search_in_slice(nums[k:], target)
			if res2 == -1 {
				return (-1)
			}
			return (k + res2)
		} else {
			return (res1)
		}
	}
}

func search_in_slice(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		centre := left + (right-left)/2
		if target == nums[centre] {
			return (centre)
		} else if target < nums[centre] {
			right = centre - 1
		} else if target > nums[centre] {
			left = centre + 1
		}
	}
	return (-1)
}
