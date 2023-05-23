package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
}

func reflectType2(x interface{}) {
	typeX := reflect.TypeOf(x)
	valueX := reflect.ValueOf(x)
	fmt.Println(typeX)
	fmt.Println(valueX)
}

func reflectType3(x interface{}) {
	typeX := reflect.TypeOf(x)
	fmt.Println(typeX.Kind()) // struct
	fmt.Println(typeX)        // main.book
}

type book struct {
}

func main() {
	var a int64 = 123
	reflectType(a)
	var b string = "从0到Go语言微服务架构师"
	reflectType2(b)
	var t book
	reflectType3(t)
}
