// Функция ищет один из пиковых элементов массива
// (элемент, который больше своих соседей)

package peak_element

func findPeakElement(nums []int) int {
	left := 0
	n := len(nums) - 1
	right := n
	for left+1 < right {
		centre := left + (right-left)/2
		if ((centre == 0) || (nums[centre-1] < nums[centre])) && ((centre == n) || (nums[centre+1] < nums[centre])) {
			return (centre)
		} else if (nums[centre-1] < nums[centre]) && (nums[centre+1] > nums[centre]) {
			left = centre
		} else {
			right = centre
		}
	}
	if nums[left] >= nums[right] {
		return (left)
	} else {
		return (right)
	}
}
