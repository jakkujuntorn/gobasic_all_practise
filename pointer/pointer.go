package main

import (
	"context"
	"fmt"
	_ "reflect"
	"time"
)

type Personal struct {
	name string
	age  int
}

type User struct {
	username string
	Personal
}

type BianryTree struct {
	value int
	left  *BianryTree
	right *BianryTree
}

type noteBook struct {
	Content []byte
}

// confrom ตาม interface Write interface
func (n *noteBook) Write(p []byte) (int, error) {
	// fmt.Println("Start Write ")
	n.Content = append(n.Content, p...)
	return len(p), nil
}

func (n noteBook) String() string {
	// fmt.Println("Start String")
	return string(n.Content)
}

type exit int

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// ใส time.Sleep แล้ว ctx มันก็ทำงานเสร็จ มันเลยเข้า  case
	time.Sleep(10 * time.Second)
	ch := time.After(200 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("CTX Done")
	case <-ch:
		fmt.Println("Ch ")
	}
	// การ new จะได้  pointer เหรอ ****
	// ss := 23
	// ex := new(exit)
	// fmt.Println(ex)
	// ex = &(*exit)(ss) // มันขึ้นมาให้เองคืออะไร

	// fmt.Println(pp.name)

	// nb := noteBook{}
	// fmt.Println(nb)

	// // ตรงนี้ไปเรียกใช้ func (n *noteBook)Write(p []byte) ด้วย
	// // เพราะ Fprintln มันไปเรียก io.Write ซค่งเรา confrom ตาม intrface WRite ให้กับ REcive func ไว้แล้ว
	// fmt.Fprintln(&nb, "hello world")

	// // ตรงนี้ไปเรียกใช้  func (n *noteBook)String() string {
	// fmt.Println(nb)  // อยู่ที่ recive func ว่าใช้ pointer รึป่าว ถ้าไม่ใช้ ก็ส่งค่าปกติเข้าไป
	// fmt.Println(&nb) // recive Func จะใช้หรือไม่ใช้ pointer ถ้าส่ง pointer มันจะใช้ได้

	// // fmt.Println(string(nb.Content))

	// fmt.Println(0.1)
	// fmt.Println(0.1 == 0.1)
	// fmt.Println(0.1+0.2 == 0.3)

	// var sum float32
	// for i := 0; i < 10; i++ {
	// 	sum += 0.1
	// }

	// fmt.Println(sum)
	// fmt.Println(sum == 1.0)
	// fmt.Println(sum == 1.0000001)

	// PlusPointerAddress()
	// binaratreeTEst()

	// d1:=Personal{
	// 	name: "N",
	// 	age: 45,
	// }
	// d2:=Personal{
	// 	name: "S",
	// 	age: 54,
	// }

	// dd:= []Personal{d1,d2}
	// fmt.Println(dd)
	// fmt.Println(reflect.TypeOf(d2))

}

func binaratreeTEst() {
	root := BianryTree{value: 1}
	left := BianryTree{value: 2}
	right := BianryTree{value: 3}
	root.left = &left
	root.right = &right

	// root:= new(BianryTree)

	// fmt.Println(root)
	// fmt.Println(&root.right) // ได้ที่อยู่
	// fmt.Println(*root.right) // ไม่ได้อะไรเลย

	// ทำไมถึงทำงาน 2 รอบ - ตอบ เพราะ มันถูกเรียกตัวเองใน func มันเอง
	binaryTEst2(&root)
}

func binaryTEst2(node *BianryTree) {
	// fmt.Println(node.value)
	// fmt.Println(node.left.value)
	i := 1
	if node != nil {
		binaryTEst2(node.left)
		fmt.Println("V:", node.value) // ทำไมบรรทัดนี้แสดงค่าทุก ค่า
		binaryTEst2(node.right)
		i++

		// pointer
		// fmt.Println(node.right.value)
		// fmt.Println(node.left.value)
	}
	// fmt.Println("i:",i)

}

func PlusPointerAddress() {
	var x *int
	y := 5
	x = &y
	*x++

	fmt.Println(*x, x)

	// ใช้ new สร้าจะได้ pointer
	a := new(Personal)
	// a = &Personal{name: "JAck", age: 15}
	*a = Personal{name: "JAck", age: 15}
	a.age = 55
	fmt.Println(a)

	uu := User{}

	uu.username = "zilch@gg.com"
	uu.name = "John"
	uu.age = 38
	fmt.Println(uu)

	ss := User{
		"Gmail.com",
		Personal{"John5", 44},
	}

	fmt.Println(ss.name)

	// ควรใส tag เพื่ออ่านง่าย และ เวลาใครมาย้าย ตำแหน่งในstruct จะไม่ error
	dd := User{
		username: "Hotmail.com",
		// เบิ้ล ค่าลงไป
		Personal: Personal{
			name: "JKKK",
			age:  45,
		},
	}

	fmt.Println(dd)
	fmt.Println(dd.name)
}

func Pointer() {

	// i := 777
	// //  & คือเอาค่า address มาใส
	// p := &i
	// fmt.Println("pointer Value1:", p)
	// fmt.Println("pointer address1:", *p)

	// *p = *p / 2
	// fmt.Println("pointer Value1:", *p)
	// fmt.Println("pointer address1:", p)

	// สร้างตัวแปร pointer
	// var s *int
	// *s = *p
	// fmt.Println("pointer address2:",*s)

	// fmt.Println("pointer:", s)

	fmt.Println("POINTER")
}

func Pointer_setValue() {
	a := "John"

	b := &a
	fmt.Println("Before A Pointer Address: ", &a)
	fmt.Println("Before A Pointer Value:", a)

	// b เป็น pointer อยู่แล้ว
	// fmt.Println("After B Pointer Address1:",&b)

	// แสดง address ของ pointer
	fmt.Println("After B Pointer Address2:", b)
	// * ดู value ของ pointer
	fmt.Println("After B Pointer Value:", *b)

	fmt.Println("***************")

	// ใสค่า b ต้องมี * เพราะคือการใสค่าเข้า address
	*b = "John 5"

	a = "Russy"

	// &  ดู address ของตัวแปร
	fmt.Println("After A Pointer Address: ", &a)
	fmt.Println("After A Pointer Value:", a)

	// b เป็น pointer อยู่แล้ว
	// fmt.Println("After B Pointer Address1:", &b)

	fmt.Println("After B Pointer Address2:", b)
	fmt.Println("After B Pointer Value:", *b)

	// a2:= 45
	changeValue(&a)
	fmt.Println("Change Func", a)
}

// ส่ง pointer เข้ามาเปลี่ยนค่า
// *string คือ address ที่เป็น string
func changeValue(x *string) {
	// ค่าจะใสใน address  เลยไม่ต้อง return ค่าออกไป
	*x = "KKK"
	// fmt.Println("Change:", x)
}

func t_pointer(a *string) {
	var ss string = "1150"
	*a = "Hello :" + ss // เอาค่ามาต่อตรงๆได้เลย ?
}
