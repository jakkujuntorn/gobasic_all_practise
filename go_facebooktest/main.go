package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("")

	//*****************************************
	// numbers := [6]int{0, 1, 2, 3, 4, 5} // ส่ง array จะเปลี่ยนค่าไม่ได้
	// // numbers := []int{0, 1, 2, 3, 4, 5}  // ส่ง slice จะเปลี่ยนค่าได้
	// fmt.Println("==================")
	// CopyValuePase_Int(numbers)
	// fmt.Println(numbers)

	// // text:=[3]string{"AAA","BBB","CCC"}// ส่ง array จะเปลี่ยนค่าไม่ได้
	// text := []string{"AAA", "BBB", "CCC"} // ส่ง slice จะเปลี่ยนค่าได้
	// fmt.Println("Before:", text)
	// CopyValuePase_String(text)
	// fmt.Println("After:", text)

	//*******************************
	// Bracket_Scpoe()

	//*******************************
	// ParseFloat()

	//*************************
	// Loop_Float64()

	//*************  Pointer *****
	// p1:=&Personal{ // ถ้าส่ง pointer เปลี่ยนค่า และ recive Func ต้องเป็น pointer ด้วย
	// 	Name: "John5",
	// 	Age:50,
	// }
	// p2:= p1 // ถ้าส่ง pointer เปลี่ยนค่า และ recive Func ต้องเป็น pointer ด้วย
	// _=p2
	// p2.Increment_Age()
	// p2.Name="Jakkujuntorn"
	// fmt.Println(p1)
	// p2.Chang_Name()
	// fmt.Println(p1)

	mapString()

}

func CopyValuePase_Int(numbers [6]int) {
	fmt.Println(&numbers)
	for i, n := range numbers {
		numbers[i] = n * n
	}
}

func CopyValuePase_String(text []string) {
	for i, value := range text {
		text[i] = value + "1150"
	}
}

func Bracket_Scpoe() { // วงเล็บคือขอบเขตในการทำงาน
	x := 1
	fmt.Println("X before:", x)
	{
		// x=2 // แบบนี้ x จะเปลี่ยนค่า after เปลี่ยนค่าด้วย
		x := 2 // x ไม่เลี่ยนค่า แสดงแค่ใน inside
		fmt.Println("X inside1 :", x)
	}
	{ // ในวงเล็บถ้าไม่มี  x จะใช้ x ตัวบน ถ้าสร้างในวงเล็บจะใช้ x ในวงเล็บ
		// x:=15
		x = 15
		// _=x
		fmt.Println("x inside2:", x)
	}

	fmt.Println("X after:", x)
}

func ParseFloat() {
	// decimalString := "1.25" // จะได้ค่าเท่ากัน
	decimalString := "1.24"

	decimal32, err := strconv.ParseFloat(decimalString, 32)
	if err != nil {
		fmt.Println("Error ParseFloat 32")
	}
	decimal64, err := strconv.ParseFloat(decimalString, 64)
	if err != nil {
		fmt.Println("Error ParseFloat 64")
	}
	fmt.Println(decimal32)              // output 1.2000000476837158
	fmt.Println(decimal64)              // output  1.2
	fmt.Println(decimal32 == decimal64) // false
}

func Loop_Float64() { // เรื่องเดียวกับ ParseFloat
	existingVersions := []float64{0.2, 0.4, 0.7, 0.8, 1.2} // ถ้าเปลี่ยนค้เป็น 1.25 จะเท่ากัน

	newVersionStr := "1.2" // ถ้าเปลี่ยนค้เป็น 1.25 จะเท่ากัน

	newVersion, _ := strconv.ParseFloat(newVersionStr, 32)
	_ = newVersion

	isNewversion := true
	_ = isNewversion

	for _, version := range existingVersions {

		fmt.Println("existingVersions:", version)

		if newVersion <= version { // if แบบนี้ คืออะไร
			isNewversion = false
			break
		}

	}
	fmt.Println("NewVersion:", newVersion)
	fmt.Println(isNewversion)
}

type Personal struct {
	Name string
	Age  int
}

func (p *Personal) Increment_Age() {
	p.Age++
}

func (p *Personal) Chang_Name() {
	p.Name = "Russy"
}

func mapString() {
	type Item struct {
		Name  string
		Price float64
	}

	type Item2 struct {
		Name  string
		Price float64
	}

	items := map[string]Item{
		"apple":  {"Apple", 1.0},
		"banana": {"BAnana", 2.0},
	}

	fmt.Println(items)

	items["apple"] = Item{"a", 4.5}
	fmt.Println("------------")
	// items["apple"].Price = 4.5 // error เพราะมันระลุว่าต้องเป็น struct Item

	fmt.Println(items)
	// fmt.Println(items["apple"])

}
