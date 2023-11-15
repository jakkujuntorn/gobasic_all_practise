package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("")

	// *********  Check struct *************
	// TestCheckType_Struct() // check struct

	//*****************  Checr type validble ****************
	// s1 := []any{14, "SS", 10.2}
	// for _, v := range s1 {
	// 	CheckType_Valiable(v) // check หลายแบบ
	// }

	// ss := 77
	// CheckType_Valiable2(ss) // check แบบเดียวตามเงื่อนไข

	//******************* Error ***************************
	customErr := ErrorCustomer{
		Message: "",
		Code:    1150,
	}
	newError := errors.New("Error golang")

	_ = customErr
	_ = newError

	//  CheckType_Error1(customErr)
	// CheckType_Error1(newError)

	// CheckType_Error2(newError)// error อยู่


	p1 := Promotion{
		ID:   1,
		Name: "Book",
		DiscountPrice: 50,
	}
	u1 := User{
		Name:     "jack",
		Lastname: "Five",
	}
	c1 := Company{
		Name:     "jack",
		Lastname: "Five",
	}
_=p1
_=u1
_=c1

	// CheckType_Struct1(c1)
	CheckType_Struct2(p1)

	// CheckType_Struct3(u1)
	// CheckType_Struct3(p1)
	// CheckType_Struct3(c1)


	// fmt.Println(reflect.TypeOf(u1).Name())
	

}

type Promotion struct {
	ID            int
	Name          string
	DiscountPrice int
}

type User struct {
	Name     string
	Lastname string
}

type Company struct {
	Name     string
	Lastname string
}

type ErrorCustomer struct {
	Message string
	Code    int
	error
}

func TestCheckType_Struct() {

	p1 := Promotion{
		ID:   1,
		Name: "Book",
	}
	u1 := User{
		Name:     "jack",
		Lastname: "Five",
	}
	c1 := Company{
		Name:     "jack",
		Lastname: "Five",
	}

	loop := []any{p1, u1, c1}
	_ = loop

	for _, v := range loop {
		CheckType_Struct1(v)
	}

	// for _, v := range loop {
	// 	CheckType_Struct2(v)
	// }
}

func CheckType_Struct1(s any) { // check ได้หลายแบบตาม case ต่างๆ
	t := reflect.TypeOf(s).Name()
fmt.Println(t)
	// t2 := reflect.TypeOf(s).Kind() // คือะไร // => struct
	// fmt.Println(t2)

	switch t {
	case "Promotion":
		fmt.Println("Promotion")
	case "User":
		fmt.Println("User")
	case "Company":
		fmt.Println("Company")
	default:
		fmt.Println("Orther Type")
	}
}

func CheckType_Struct2(s any) { // check แบบเดียว
	// fmt.Println(s.(Promotion)) // print ไม่ได้

	c, ok := s.(Promotion) // 
	if !ok {
		fmt.Println("Cast Error")
	} else {
		fmt.Println(c)
	}
}

func CheckType_Struct3(s any) { // check หลายแบบด้วย type
	// fmt.Println(s.(type))// print ไม่ได้

	switch s.(type) {

	case Promotion:
		fmt.Println("Promotion")
	case User:
		fmt.Println("User")
	case Company:
		fmt.Println("Company")
	default:
		fmt.Println("Orther Type")

	}
}

func CheckType_Valiable(s any) { // check ได้หลายแบบตาม case ต่างๆ

	// t1 := reflect.TypeOf(s).Name() // คือะไร // => ประเภทข้อมูล int | string |float64
	// fmt.Println(t1)

	// t2 := reflect.TypeOf(s).Kind() // คือะไร // => ประเภทข้อมูล int | string |float64
	// fmt.Println(t2)

	switch s.(type) {
	case int:
		fmt.Println("Int")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Orther Type")
	}
}

func CheckType_Valiable2(text any) { // check ได้แบบเดียว ตาม
	t, ok := text.(int)
	fmt.Println(ok)

	if ok {
		fmt.Printf("Value: %v", t)
	} else {
		fmt.Println("Check Error")
	}
}

func CheckType_Error1(err error) { // check error หลายแบบตาม case
	// t1 := reflect.TypeOf(err).Name() // คือะไร // => ประเภทข้อมูล
	// fmt.Println(t1)

	// t2 := reflect.TypeOf(err).Kind() // คือะไร // => ประเภทข้อมูล int | string |float64
	// fmt.Println(t2)

	switch e := err.(type) {
	case ErrorCustomer:
		fmt.Printf("ErrorCustomer: [Message: %v,Code: %v]", e.Message, e.Code)
	case error:
		fmt.Printf("Error Go: %v", e.Error())
	}
}

func CheckType_Error2(err error) { // check error แบบเดียวตามเงื่อนไข
	e1 := errors.New("")
	e2 := ErrorCustomer{}
	_ = e1
	_ = e2

	// if errors.Is(err, e2) {
	// 	fmt.Println("ErrorCustomer")
	// }

	// if errors.As(err, error) {
	// 	fmt.Println("Error Go lang")
	// }

}
