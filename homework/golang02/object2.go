package golang02

import "fmt"

/***

2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。

*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int32
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee Info:{EmployeeID: %d, Name: %s, Age: %d}\n", e.EmployeeID, e.Name, e.Age)
	// fmt.Printf("Employee ID: %d, Name: %s, Age: %d\n", e.EmployeeID, e.Name, e.Age)
}

func Object2() {

	employee := Employee{
		Person: Person{
			Name: "zhangsan",
			Age:  20,
		},
		EmployeeID: 10001,
	}
	employee.PrintInfo()

}

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Employee struct {
// 	Person
// 	EmployeeID string
// }

// func (e Employee) PrintInfo() {
// 	fmt.Printf("Employee ID: %s, Name: %s, Age: %d\n", e.EmployeeID, e.Name, e.Age)
// }
