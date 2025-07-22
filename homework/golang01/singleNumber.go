package golang01

func SingleNumber(nums []int) int {

	var m map[int]int = make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {

		if v == 1 {
			return k
		}
	}
	return -1
}
