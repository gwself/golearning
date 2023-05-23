package book

import "fmt"

func init() {
	fmt.Println("我是包的初始化")
}

type Book struct {
	/**
	大写为public
	*/
	Title  string
	Author string
	/**
	小写为private
	*/
	price float64
}

/*
	方法大写开头为public  小写开头为private
*/
func (receiver Book) ShowBookInfo() {
	fmt.Println(receiver.Title, "的作者是", receiver.Author, "该书价格是：", receiver.price)
}

/**
小写开头为private 方法
*/
func (receiver Book) name() {
	fmt.Println("this is a private method", receiver)
}

func (receiver *Book) SetPrice(newPrice float64) {
	receiver.price = newPrice
}

const (
	A = iota // 0
	B        // 1
	C        // 2
	D        // 3
)

func (receiver Book) GetLevel() int {
	if receiver.price < 10 {
		return D
	} else if receiver.price < 20 {
		return C
	} else if receiver.price < 30 {
		return B
	} else {
		return A
	}
}
