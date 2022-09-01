// Функция возвращает результат пересечения двух массивов,
// с точным повторением пересекающихся элементов

package intersection_of_two_arrays

func intersect(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]int)
	set2 := make(map[int]int)
	var arr [2000]int
	l := 0
	if len(nums1) > len(nums2) {
		l = len(nums1)
	} else {
		l = len(nums2)
	}
	for i := 0; i < l; i++ {
		if i < len(nums1) {
			set1[nums1[i]]++
		}
		if i < len(nums2) {
			set2[nums2[i]]++
		}
	}
	l = 0
	for k := range set1 {
		for i := 0; i < findMin(set1[k], set2[k]); i++ {
			arr[l] = k
			l++
		}
	}
	return (arr[:l])
}

func findMin(x, y int) int {
	if x <= y {
		return (x)
	}
	return (y)
}
