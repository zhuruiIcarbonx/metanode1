# metanode1

https://leetcode.cn/problems/single-number/
https://leetcode.cn/problems/palindrome-number/
https://leetcode.cn/problems/valid-parentheses/
https://leetcode.cn/problems/longest-common-prefix/
https://leetcode.cn/problems/plus-one/
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
https://leetcode.cn/problems/merge-intervals/
https://leetcode.cn/problems/two-sum/
112233



------------------------

import(
	"fmt"
	"strings"
)
func longestCommonPrefix(strs []string) string {

	prifix := strs[0]
	for i := len(prifix) - 1; i >= 0; i-- {
        fmt.Println("prifix=",prifix,"i=",i)
		for j := 1; j < len(strs); j++ {
            
            fmt.Println("i=",i,"-----j=",j)
			str := strs[j]
			if len(str) >= len(prifix) && strings.HasPrefix(str, prifix) {
				if j == len(strs)-1 {
					return prifix
				}
				continue
			} else {
				break
			}
		}
        if i==1{
            prifix = prifix[0:1]
        }else if i==0{
            return ""
        }else{
            prifix = prifix[:i]
        }
        
	}
	return ""
}

------------------------


import(
	"fmt"
)


func merge(intervals  [][]int) [][]int {

	//   var mergeMap = make([int]int)
	//   var arr  [][]int
      for i:=0;i<len(intervals);i++{
		for j:=len(intervals)-1;j>=i+1;j--{
            fmt.Println("i=",i,"j=",j)
            fmt.Println("intervals[i][0]=",intervals[i][0],"intervals[i][1]=",intervals[i][1])
            fmt.Println("intervals[j][0]=",intervals[j][0],"intervals[j][1]=",intervals[j][1])
			if intervals[j][0]>intervals[i][1]||intervals[i][0]>intervals[j][1]{//没有交集
                fmt.Println("merge===============",1)
			}else if intervals[j][0]>=intervals[i][0]&& intervals[j][1]<=intervals[i][1]{//i包含j  [i0,i1]
                fmt.Println("merge===============",2)
                if(j==1){
                    intervals = append(intervals[0:1],intervals[j+1:]...)
                }else if j+1>=len(intervals){
                    intervals = intervals[:j]
                }else{
                    intervals = append(intervals[:j-1],intervals[j+1:]...)
                }

			}else if intervals[i][0]>=intervals[j][0]&& intervals[i][1]<=intervals[j][1]{//j包含i [j0,j1]
                fmt.Println("merge===============",3)
				intervals[i] = intervals[j]
                if(j==1){
                    intervals = append(intervals[0:1],intervals[j+1:]...)
                }else if j+1>=len(intervals){
                    intervals = intervals[:j]
                }else{
                    intervals = append(intervals[:j-1],intervals[j+1:]...)
                }

			}else if intervals[i][0]<=intervals[j][1]&&intervals[j][1]<=intervals[i][1]{//有交叉，j<i  [j0,i1]
                fmt.Println("merge===============",4)
				intervals[i] = []int{intervals[j][0],intervals[i][1]}
                if(j==1){
                    intervals = append(intervals[0:1],intervals[j+1:]...)
                }else if j+1>=len(intervals){
                    intervals = intervals[:j]
                }else{
                    intervals = append(intervals[:j-1],intervals[j+1:]...)
                }

			}else if intervals[j][0]<=intervals[i][1]&&intervals[j][1]>=intervals[i][1]{//有交叉，j>i  [i0,j1]
                fmt.Println("merge===============",5)
                fmt.Println("intervals[i]===============",[]int{intervals[i][0],intervals[j][1]})
				intervals[i] = []int{intervals[i][0],intervals[j][1]+0}
                fmt.Println("intervals===============",intervals)
                if(j==1){
                    intervals = append(intervals[0:1],intervals[j+1:]...)
                }else if j+1>=len(intervals){
                    intervals = intervals[:j]
                }else{
                    intervals = append(intervals[:j-1],intervals[j+1:]...)
                }
				
                fmt.Println("intervals===============",intervals)
			}
		}
	  }
	
	  return intervals
	  

}


func mergeIntervals(intervals  [][]int) [][]int {


	fmt.Println("intervals===============",intervals)
	
	for{
		intervals = merge(intervals)
		checkInter := merge(intervals)
		if len(intervals)==len(checkInter){
			return intervals
		}
	}

}


------------------------



import(
	"fmt"
	"strings"
	"strconv"
)



func PlusOne(digits []int) []int {

	var builder strings.Builder
	for _,v := range digits{
        fmt.Println("-------------------------v=",v)
        // sv := string(v)
        // fmt.Println("-------------------------sv=",sv)
		builder.WriteString(strconv.Itoa(v))

	}
	str := builder.String()
    fmt.Println("-------------------------str=",str)
	num,_ := strconv.Atoi(str)
	num += 1
	fmt.Println("---------",num)
	strBack := strconv.Itoa(num)
	var arr []int
	for _,v := range strBack{
		nv,_ := strconv.Atoi(string(v))
        arr = append(arr,nv)
	}
	return arr

}


------------------------

import(
	"fmt"
)



func RemoveDuplicates(nums []int) int {

	var arr []int
	var numMap map[int]int = make(map[int]int)
	for _,v := range nums{
		if _,ok := numMap[v]; ok{
             continue
		}else{
			arr = append(arr,v)
			numMap[v] = 1
		}
	}
    
    fmt.Println("-------------------------arr=",arr)
	return len(arr)

}



------------------------

import(
	"fmt"
)

func twoSum(nums []int,target int) []int {

	for i :=0;i<len(nums);i++{
		for j :=i+1;j<len(nums);j++{
			fmt.Println("i=",i,"j=",j)
			if target == nums[i]+nums[j]{
				fmt.Println("return------i=",i,"j=",j)
				return []int{i,j}
			}
		}
	}
	return []int{0,0}

}



------------------------


