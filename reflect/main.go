package main

import (
	"fmt"
	"reflect"
)

type Customer struct {
	ID   int
	Name string
}

func main() {
	c1 := Customer{
		ID:   115045,
		Name: "Rosea",
	}

	Relfect_TYpeOf(c1)

	Relfect_ValueOf(c1)
}

func Relfect_TYpeOf(t interface{}) {
	t1 := reflect.TypeOf(t).Name()
	fmt.Println(t1)
}

func Relfect_ValueOf(t interface{}) {
	t1 := reflect.ValueOf(t)
	fmt.Println(t1)
}

// func Relfect_Value(t interface{})  {
// 	t1:=reflect.Value.Type(t)
// 	fmt.Println(t1)
// }

// func Relfect_Value_Interface(t interface{})  {
// 	t1:=reflect.Value.Interface(t)
// 	fmt.Println(t1)
// }

// func Relfect_Value_Kind(t interface{})  {
// 	t1:=reflect.Value.Kind(t)
// 	fmt.Println(t1)
// }

// func Relfect_Value_New(t interface{})  {
// 	t1:=reflect.New(t)
// 	fmt.Println(t1)
// }

// func Relfect_Value_MakeSlice(t interface{})  {
// 	t1:=reflect.MakeSlice(t reflect.Type, len int, cap int)
// 	fmt.Println(t1)
// }

// func Relfect_Value_PtrTo(t interface{})  {
// 	t1:=reflect.PtrTo(t reflect.Type)
// 	fmt.Println(t1)
// }

// func Relfect_Value_Indirect(t interface{})  {
// 	t1:=reflect.Indirect(v reflect.Value)
// 	fmt.Println(t1)
// }

// func Relfect_Value_FieldByName(name string)  {
// 	t1:=reflect.Value.FieldByName(name)
// 	fmt.Println(t1)
// }

// func Relfect_Value_MethodByName(name string)  {
// 	t1:=reflect.Value.MethodByName(name)
// 	fmt.Println(t1)
// }

// func Relfect_Value_Set(name string)  {
// 	t1:=reflect.Value.Set(name)
// 	fmt.Println(t1)
// }
