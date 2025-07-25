package gorm1

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/***

题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

**/

type Student struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Age   int    `gorm:"not null"`
	Grade string `gorm:"type:varchar(50);not null"`
}

func initDb() *gorm.DB {
	// 连接到mysql数据库
	db, err := gorm.Open(mysql.Open("root:123456zz@tcp(127.0.0.1:3306)/meta?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Crud() {

	// 连接到mysql数据库
	db := initDb()

	db.AutoMigrate(&Student{})

	// 1.编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	stu := &Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}

	db.Debug().Create(stu)

	fmt.Printf("stu is :%v \n", stu)

	db.Debug().Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", "李四", 19, "三年级")
	db.Debug().Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", "王五", 14, "三年级")

	//2. 查询所有年龄大于 18 岁的学生信息。
	var list []Student

	db.Debug().Where("age > ?", 18).Find(&list)
	fmt.Printf("查询所有年龄大于 18 岁的学生信息:%v \n", list)

	// var list2 []map[string]interface{}
	// db.Raw("SELECT * FROM students WHERE age > ?", 18).Scan(&list2)
	// fmt.Printf("list2 is :%v \n", list2)

	// 3. 更新姓名为 "张三" 的学生年级为 "四年级"
	var zStu Student

	db.Debug().Model(&zStu).Where("name = ?", "张三").Update("grade", "四年级")
	fmt.Printf("张三更新后信息:%v \n", zStu)

	zStu2 := Student{}
	db.Debug().First(&zStu2, "name=?", "张三")
	zStu2.Grade = "五年级"
	db.Debug().Save(&zStu2)
	fmt.Printf("张三2更新后信息:%v \n", zStu2)

	// db.Exec("UPDATE students SET grade = ? WHERE name = ?", "四年级", "张三")

	// 4. 删除年龄小于 15 岁的学生记录

	var listBefore, listAfter []Student
	db.Where("age<?", 15).Find(&listBefore)
	fmt.Printf("listBefore:%v \n", listBefore)

	db.Debug().Exec("DELETE FROM students WHERE age < ?", 15)
	db.Debug().Where("age<?", 15).Find(&listAfter)
	fmt.Printf("listAfter:%v \n", listAfter)

}
