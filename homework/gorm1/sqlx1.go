package gorm1

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

/***

题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，
employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
**/

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func Sqlx1() {

	db, err := sqlx.Open("mysql", "root:123456zz@tcp(127.0.0.1:3306)/meta?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()
	// db.AutoMigrate(&Employee{})

	db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "张三", "技术部", 8000)
	db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "李四", "技术部", 9000)
	db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "王五", "市场部", 7000)
	db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)", "赵六", "技术部", 9500)

	// 查询所有部门为 "技术部" 的员工信息
	var employees []Employee
	err = db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		fmt.Println("Error querying employees:", err)
		return
	}
	fmt.Println("Employees in 技术部:", employees)

	// 查询工资最高的员工信息
	var highestPaidEmployee Employee
	err = db.Get(&highestPaidEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")

	if err != nil {
		fmt.Println("Error querying highest paid employee:", err)
		return
	}
	fmt.Println("Highest paid employee:", highestPaidEmployee)
}
