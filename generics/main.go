package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("")
	fmt.Println(Add1(1, 1))
	fmt.Println(Add2(2.2, 2.2))
	fmt.Println(Add3(3, 3))
	fmt.Println(Add4("4", "5"))
	fmt.Println(Add5(5,5))


	dd := Data[int]{ //ต้องมากำหนด type ตรงนี้ด้วย
		Id:   4,
		Name: "Jack",
		Data: 1150,
	}

	fmt.Println(dd)
}

type UserData interface {
	int | float64 | string
}

type Data[T UserData] struct {
	Id   int
	Name string
	Data T // กำหนด ประเถทข้อมูลว่าจะให้เป็นประเภทอะไรจาก UserData
}



func Add1[T int | float64](a T, b T) T {
	return a + b
}

type typeParamiter interface {
	int | float64
}

func Add2[T typeParamiter](a T, b T) T {
	return a + b
}

func Add3[T ~int | float64](a T, b T) T {
	return a + b
}

func Add4[T constraints.Ordered](a T, b T) T {
	return a + b
}

type ParamiterInt int
func Add5[T ParamiterInt](a T, b T) T {
	return a + b
}
