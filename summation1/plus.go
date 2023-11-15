package Plus

import "fmt"

// import (
// 	"fmt"
// )

// ตั้งค่าตัวอักษรตัวแรกเป็นตัวพิมพ์ใหญ่ เพื่อให้คนอื่นใช้ได้
func Sum(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func Add()string {
 return "Plus Add"
}

type Book struct {
	name string
	price int
}

func (Bo Book) SetDataBook() {
	Bo.name="Potter"
	Bo.price= 1500
	fmt.Println("Book DAta:", Bo)
}

// func  SetDataBook() {
	
// 	fmt.Println("Book DAta:")
// }