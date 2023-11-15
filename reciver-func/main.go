package main

import (
	"encoding/json"
	"fmt"
	"math"

)

func main() {

	// ss.ServiceLayer()

	// ChangeJson()
	dd := data{
		Name: "Jack",
		Age:  30,
	}
	// fmt.Println(dd)
	fmt.Println("-----------------")

	dataGet := dd.GetData()
	_=dataGet
	// fmt.Println(dataGet)
	fmt.Println("-----------------")

	// dd.ChangData()
	aa := data{
		Name: "Russy",
		Age:  39,
	}
	dd.setData(aa)
	// fmt.Println(dd)


	r2:=data{
		Name: "JAck",
		Age: 33,
	}
	r3:=data{
		Name: "Rose",
		Age: 47,
	}
	r4:=data{
		Name: "Tron",
		Age: 39,
	}
	// r1:= NewData(r2).GetData()
	r1:= NewData(r2).GetData()
	fmt.Println(r1)
	fmt.Println("-----------------")

	r1.ChangData(r4)
	fmt.Println(r1)
fmt.Println("-----------------")

	r1.setData(r3).GetData()
	fmt.Println(r1)
	
	// dddd:= r1.GetData()
	// fmt.Println(dddd)
}

type data struct {
	Name string
	Age  int
}

func NewData(dd data) *data {
	return &data{
		Name: dd.Name,
		Age: dd.Age,
	}
}

func (d *data) GetData() data {

	return *d
}

func (d *data) ChangData(da data) {
	d.setData(da)
}

func (d *data) setData(da data) *data {
	d.Name = da.Name
	d.Age = da.Age
	return d
}

func ChangeJson() {
	mySt := MystructType{
		User1: "JAck",
		// User2: "Russy",
	}

	encod := MysteuctEncoder{Mystruct: mySt}
	bb, _ := encod.Encoder()
	fmt.Println(mySt)
	fmt.Println(bb)              // output []byte
	fmt.Println(encod.Encoder()) // output []byte and nil ถ้าไม่มี error

}

type MysteuctEncoder struct {
	Mystruct MystructType
}

func (m *MysteuctEncoder) Encoder() ([]byte, error) {
	return json.Marshal(m.Mystruct)
}

type MystructType struct {
	User1 string
	// User2 string
}

type circle struct {
	redius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.redius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.redius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape :", c)
// 	fmt.Println("Area :", c.area())
// 	fmt.Println("Perimeter :", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape :", r)
// 	fmt.Println("Area :", r.area())
// 	fmt.Println("Perimeter :", r.perimeter())
// }

type shap interface {
	area() float64
	perimeter() float64
}

func print(s shap) {
	fmt.Println("Shape :", s)
	fmt.Println("Area :", s.area())
	fmt.Println("Perimeter :", s.perimeter())
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.redius, 3)
}

func Reciver() {
	c1 := circle{redius: 5.}
	r1 := rectangle{width: 3., height: 2.1}

	// fmt.Println("******* Revicer ******")
	// printCircle(c1)
	// printRectangle(r1)
	// fmt.Println("**************")

	fmt.Println("******* Revicer ******")
	print(c1)
	print(r1)
	fmt.Println("**************")

	// var s shap= circle{redius: 2.5}
	var s shap = rectangle{width: 3., height: 2.1}

	// s.volume
	// ทำงานอย่างไร
	ball, ok := s.(circle)

	if ok == true {
		fmt.Println("Ball Volume:%v\n", ball.volume())
	}

	switch value := s.(type) {
	case circle:
		fmt.Println("Circle: %#v", value)

	case rectangle:
		fmt.Println("Rectangle: %#v", value)
	}

}
