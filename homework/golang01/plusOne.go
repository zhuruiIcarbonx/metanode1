package golang01

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func PlusOne(digits []int) []int {

	var builder strings.Builder
	var bigInt big.Int
	for _, v := range digits {
		fmt.Println("-------------------------v=", v)
		// sv := string(v)
		// fmt.Println("-------------------------sv=",sv)
		builder.WriteString(strconv.Itoa(v))

	}
	str := builder.String()
	fmt.Println("-------------------------str=", str)
	_, _ = bigInt.SetString(str, 10)
	fmt.Println("-------------------------num=", bigInt.String())
	bigInt.Add(&bigInt, big.NewInt(1))
	fmt.Println("---------", bigInt.String())
	strBack := bigInt.String()
	var arr []int
	for _, v := range strBack {
		nv, _ := strconv.Atoi(string(v))
		arr = append(arr, nv)
	}
	return arr

}
