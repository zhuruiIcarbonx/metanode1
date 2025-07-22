package lesson16

import (
	"fmt"
	"strconv"
)

func StrToNum() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	//当使用 ParseUint 方法把字符串转换成数字时，或者使用 FormatUint 方法把数字转换成字符串时，
	// 都需要提供第二个参数 base，这个参数表示的是数字的进制，即标识字符串输出或输入的数字进制。

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	str2 := strconv.FormatUint(ui64, 10)
	fmt.Printf("uint64转换为字符串: %s \n", str2)
}

func DoInterface() {
	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}

	// var i interface{} = 3
	switch v := i.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}

	//test3
	var aa Supplier = &DigitSupplier{value: 1}
	fmt.Println(aa.Get())

	b, ok := (aa).(*DigitSupplier)
	fmt.Println(b, ok)

}

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}
