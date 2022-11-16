// package main

// import "fmt"


// type Numbers struct{
// 	num1 int
// 	num2 int
// }

// // type NumbersInterface interface{
// // 	sum() int
// // 	multiply() int
// // 	division() float64
// // 	substract() int
// // } 

// // func (n Numbers) sum() int{
// // 	return n.num1 + n.num2
// // }


// // func (n Numbers) multiply() int{
// // 	return n.num1 * n.num2
// // }


// // func (n Numbers) division() float64{
// // 	return float64(n.num1) / float64(n.num2)}


// // func (n Numbers) substract() int{
// // 	return n.num1 - n.num2
// // }

// // func main(){
// // 	var i NumbersInterface
// // 	numbers := Numbers{12, 6}
// // 	i = numbers
// // 	fmt.Printf("sum of numbers: %d\n", i.sum())
// // 	fmt.Printf("multiply of numbers: %d\n", i.multiply())
// // 	fmt.Printf("division oi = numbersf numbers: %f\n", i.division())
// // 	fmt.Printf("substract of numbers: %d\n", i.substract())

// // }
// type Square interface{
// 	SquareArea() int
// 	SquarePerimeter() int
// }


// func (n Numbers)  SquareArea() int{
// 	return n.num1 * n.num2
// }


// func (n Numbers)  SquarePerimeter() int{
// 	return (n.num1 + n.num2)*2
// }
// func main(){
// var a Square
// numbers := Numbers{1056, 994}
// a = numbers
// fmt.Printf("Площадь квадрата: %d\n", a.SquareArea())
// fmt.Printf("Периметр квадрата: %d\n", a.SquarePerimeter())
// Type("hi")
// Type(12.3)
// }


// func Type(value interface{}){ //пустой интерфейс
// 	switch value.(type){
// 	case int:
// 		fmt.Println("Type is a int")
// 	case string:
// 		fmt.Println("Type is a string")
// 	case float64:
// 		fmt.Println("Type is a float")
// 	default:    //иначе
// 		fmt.Println("Type isn't int and strng")
// 	}
// }


