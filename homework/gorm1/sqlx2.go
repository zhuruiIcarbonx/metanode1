package gorm1

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

/***

题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
**/

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func Sqlx2() {

	db, err := sqlx.Open("mysql", "root:123456zz@tcp(127.0.0.1:3306)/meta?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.Exec("drop table if exists books")
	db.Exec("CREATE TABLE IF NOT EXISTS books (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255), author VARCHAR(255), price DECIMAL(10,2))")

	db.Exec("INSERT INTO books (title, author, price) VALUES ('Go Programming', 'John Doe', 45.50), ('Advanced Go', 'Jane Smith', 60.00), ('Database Design', 'Alice Johnson', 75.00)")

	var books []Book
	err = db.Select(&books, "select * from books where price >?", 50)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Books with price > 50:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
