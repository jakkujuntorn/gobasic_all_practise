package main

import "fmt"

// type data interface {

// }

type data struct {
	name     string
	lastname string
	age      int
}

type email struct {
	username string
	password string
}

type salary struct {
	day   int
	money int
}

type dataInt struct {
	a int
	b int
}

type procressData interface {
	clearSalary()
	clearEmail()
}

type procressInt interface {
	plusInt() int
	minusInt() int
}

type Item interface {
	cost()
}

// interface ซ้อน interface
type Weapon interface {
	detail()
	Item
}
type Sword struct {
	nane string
}

func GetWeapon(w Weapon) {
	fmt.Println("Weapn Name:",w)
}

func (s Sword)detail()  {
	fmt.Println("Detail:",s.nane)
}

func (s Sword)Power()  {
	fmt.Println("Uppower :",s.nane)
}

func (s Sword)cost()  {
	fmt.Println("Cost :",s.nane)
}
// type ที่เอา struct คนอื่นมาใช้
type sword2 Sword

func main() {
	// s:= sword2{}
	// s.nane="Dogger"
	// fmt.Println(s.nane)


	// sw := Sword{"Dragon"}

	// GetWeapon(sw)
	// fmt.Println("*****....**")
	// sw.detail()
	// sw.Power()
	// sw.cost()
	// fmt.Println("OK")


}

func InterFaceTest() {

	fmt.Println("Iterface")

	a := data{
		name:     "Ruk",
		lastname: "Thong",
		age:      38,
	}
	fmt.Println("Show NAme:", a)

	a.cleaData()
	fmt.Println("Clear Name:", a)

	// moneyM := salary{
	// 	day:   20,
	// 	money: 100,
	// }

	// dataEmail := email{
	// 	username: "Jack@gmail.com",
	// 	password: "PassWoqd",
	// }

	d1 := dataInt{5, 1}
	// มันทำงานอย่างไร ****
	// เอา struct มารับค่าโดยทำแบบนี้ := ได้ไหม
	plus_minus(d1)

}

//  *************** function ที่รับค่าเป็น interface **************
func plus_minus(a procressInt) {
	pp := a.plusInt()
	mm := a.minusInt()
	fmt.Println("Interface", pp)
	fmt.Println("Interface", mm)

}

// *****************function ที่ return เป็น interface ศึกษาเพิ่ม  ****************
// func plus_minus(a procressInt) procressInt {
// 	pp := a.plusInt()
// 	mm := a.minusInt()
// 	fmt.Println("Interface", pp)
// 	fmt.Println("Interface", mm)

// 	return nil
// }

func (p dataInt) plusInt() int {
	result := p.a + p.b
	return result
}

func (m dataInt) minusInt() int {
	result := m.a - m.b
	return result
}

func (s salary) clearSalary() int {
	result := s.day * s.money
	return result
}

func (d *email) clearEmail() {
	d.username = ""
	d.password = ""
}

// เปลี่ยนค่า address ใช้ * และไม่ต้อง return
func (d *data) cleaData() {
	d.name = ""
	d.lastname = ""
	d.age = 0
	// return d
}
