package golang01

import (
	"fmt"
)

func doDetail(intervals [][]int) [][]int {

	//   var mergeMap = make([int]int)
	//   var arr  [][]int
	for i := 0; i < len(intervals); i++ {
		for j := len(intervals) - 1; j >= i+1; j-- {
			fmt.Println("i=", i, "j=", j)
			fmt.Println("intervals[i][0]=", intervals[i][0], "intervals[i][1]=", intervals[i][1])
			fmt.Println("intervals[j][0]=", intervals[j][0], "intervals[j][1]=", intervals[j][1])
			if intervals[j][0] > intervals[i][1] || intervals[i][0] > intervals[j][1] { //没有交集
				fmt.Println("merge===============", 1)
			} else if intervals[j][0] >= intervals[i][0] && intervals[j][1] <= intervals[i][1] { //i包含j  [i0,i1]
				fmt.Println("merge===============", 2)
				if j == 1 {
					intervals = append(intervals[0:1], intervals[j+1:]...)
				} else if j+1 >= len(intervals) {
					intervals = intervals[:j]
				} else {
					intervals = append(intervals[:j], intervals[j+1:]...)
				}

			} else if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] { //j包含i [j0,j1]
				fmt.Println("merge===============", 3)
				intervals[i] = intervals[j]
				if j == 1 {
					intervals = append(intervals[0:1], intervals[j+1:]...)
				} else if j+1 >= len(intervals) {
					intervals = intervals[:j]
				} else {
					intervals = append(intervals[:j], intervals[j+1:]...)
				}

			} else if intervals[i][0] <= intervals[j][1] && intervals[j][1] <= intervals[i][1] { //有交叉，j<i  [j0,i1]
				fmt.Println("merge===============", 4)
				intervals[i] = []int{intervals[j][0], intervals[i][1]}
				if j == 1 {
					intervals = append(intervals[0:1], intervals[j+1:]...)
				} else if j+1 >= len(intervals) {
					intervals = intervals[:j]
				} else {
					intervals = append(intervals[:j], intervals[j+1:]...)
				}

			} else if intervals[j][0] <= intervals[i][1] && intervals[j][1] >= intervals[i][1] { //有交叉，j>i  [i0,j1]
				fmt.Println("merge===============", 5)
				fmt.Println("intervals[i]===============", []int{intervals[i][0], intervals[j][1]})
				intervals[i] = []int{intervals[i][0], intervals[j][1] + 0}
				fmt.Println("intervals===============", intervals)
				if j == 1 {
					intervals = append(intervals[0:1], intervals[j+1:]...)
				} else if j+1 >= len(intervals) {
					intervals = intervals[:j]
				} else {
					intervals = append(intervals[:j], intervals[j+1:]...)
				}

				fmt.Println("intervals===============", intervals)
			}
		}
	}

	return intervals

}

/**
*** letcode执行时需删除所有fmt.Println。避免输出过长
**/
func Merge(intervals [][]int) [][]int {

	fmt.Println("intervals===============", intervals)

	for {
		intervals = doDetail(intervals)
		checkInter := doDetail(intervals)
		if len(intervals) == len(checkInter) {
			return intervals
		}
	}

}
