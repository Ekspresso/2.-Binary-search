// Функция возвращает результат пересечения двух массивов,
// где каждый элемент уникален

package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	var arr [2000]int
	l := 0
	if len(nums1) > len(nums2) {
		l = len(nums1)
	} else {
		l = len(nums2)
	}
	for i := 0; i < l; i++ {
		if i < len(nums1) && !set1[nums1[i]] {
			set1[nums1[i]] = true
		}
		if i < len(nums2) && !set2[nums2[i]] {
			set2[nums2[i]] = true
		}
	}
	l = 0
	for k := range set1 {
		if set2[k] {
			arr[l] = k
			l++
		}
	}
	return (arr[:l])
}
