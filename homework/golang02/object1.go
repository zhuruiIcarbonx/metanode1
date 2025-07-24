package golang02

import (
	"fmt"
)

/***

1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。

*/

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

type Circle struct {
	r float32
}

func (r Rectangle) Area() float32 {

	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float32 {

	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float32 {

	return 3.14 * c.r * c.r
}

func (c *Circle) Perimeter() float32 {

	return 2 * 3.14 * c.r
}

func Object1() {

	rectangle := Rectangle{Width: 5, Height: 6}
	circle := Circle{r: 3}
	fmt.Println("rectangle Area is:", rectangle.Area())
	fmt.Println("rectangle Perimeter is:", (&rectangle).Perimeter())
	fmt.Println("circle Area is:", circle.Area())
	fmt.Println("circle Perimeter is:", (&circle).Perimeter())
}

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Rectangle struct {
// 	Width, Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * 3.14 * c.Radius
// }

// rect := Rectangle{Width: 5, Height: 10}
// circle := Circle{Radius: 7}

// shapes := []Shape{rect, circle}

// for _, shape := range shapes {
// 	println("Area:", shape.Area())
// 	println("Perimeter:", shape.Perimeter())
// }
