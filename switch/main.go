package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	p1 := Data{
		Name: "JAck5",
		Age:  22,
	}

	p2 := Personal{
		Name: "king",
		Age:  78,
	}
	_ = p1
	_ = p2

	// SwitchPointer(p1)  // copy value
	// SwitchPointer(&p2) // pointer

	// Switch_ReflactType(p1)
	// Switch_ReflactType(p2)

	// Switch_ReflactType2(p1)
	// Switch_ReflactType2(p2)

	// day()
	// hour()
	// Fallthrough()
	// exitWithBreak() // ดู func นี้ใหม่ มันทำงานยังไง งง เหมือนกัน
	// executionOrder() // ไปดู  output ว่ามันออกค่านี้ได้ยังไง  งง มาก

}

func day() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func hour() {
	hour2 := time.Now().Hour()
	fmt.Println(hour2)
	switch hour := time.Now().Hour(); { // missing expression means "true"

	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func Fallthrough() {
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}

func exitWithBreak() {
Loop: // ตรงนี้คืออะไร ****
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}

func executionOrder() {
	// Foo(1)
	// Foo(2)
	// Foo(3)

	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
	// output ทำไมถึงไดเค่านี้
	// 2
	// 1
	// 2
	// First case
	// Second case

}

func Foo(n int) int {
	fmt.Println(n)
	return n
}

type Data struct {
	Name string
	Age  int
}

type Personal struct {
	Name string
	Age  int
}

func SwitchPointer(data interface{}) {

	switch data.(type) {
	case *Data: // สามารถ แยก pointer ได้
		fmt.Println("Pointer")
	case Data:
		fmt.Println("Copy VAlue")
	}

}

func Switch_ReflactType(data interface{}) {
	t1 := reflect.TypeOf(data).Name()
	_ = t1

	// switch t {
	switch t := reflect.TypeOf(data).Name(); t {

	// case reflect.TypeOf(Personal{}).Name():  // ใช้แบบไหนก็ได้
	case "Personal": // ใช้แบบไหนก็ได้
		fmt.Println("Personal")

	case reflect.TypeOf(Data{}).Name():
		fmt.Println("Data")
	}

	// topic string
	// switch topic {
	// case reflect.TypeOf(events.OpenAccount_Event{}).Name():
	// 	fmt.Println("step 2 Handlet switch ")
	// 	event := &events.OpenAccount_Event{}
	// 	err := json.Unmarshal(eventBytes, event)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
}

func Switch_ReflactType2(data interface{}) {
	t1 := reflect.TypeOf(data).Name()
	_ = t1

	switch data.(type) {
	// switch t := reflect.TypeOf(data).Name(); t {

	// case reflect.TypeOf(Personal{}).Name():  // ใช้แบบไหนก็ได้
	case Personal: // ใช้แบบไหนก็ได้
		fmt.Println("Personal")

	case Data:
		fmt.Println("Data")
	}

}
