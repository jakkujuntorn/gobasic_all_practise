package main

// tool input output
import (
	_ "encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	_ "unicode/utf8"

	// "gobasic/channels"
	// "gobasic/interfacetest"
	"gobasic/pointer"
	_ "gobasic/map"
	// reciver "gobasic/reciver-func"
	// books "gobasic/structData" // สร้างชื่อตัวแปร
	// "gobasic/summation1"
	errCustom "gobasic/error" // ใสเป็น event นอก
	// jsonTest "gobasic/jsonunmarshal"
// _ "gobasic/validation"
// _"gobasic/randomnumber"
)

// ** Global Variable ตัวแปรที่ใครจะใช้ก็ได้ ประกาศนอก func ***
var name = "Russy Jack"

// func init()  {
// 	fmt.Println("###################################Init Func #################################################################")
// 	for i:=0; i<5; i++ {
// 		fmt.Println("Value I in Init: %v",i)
// 	}

// }

func main() {
	pointer.PlusPointerAddress()
	// pointer variable
	// msg := "Pointer"
	// var msgPointer *string = &msg
	// fmt.Println(".....Pointer.......")
	// fmt.Println("Pointer Variable:", *msgPointer)

	// var name = "Russy Jack"
	// name = "Kwan ****"
	// // int ถ้าไม่ใสต้องใส :

	// // Implicit Declration ไม่ต้องระบุประเภท
	// age := 25
	// // age = 44

	// /* **************************** */
	// n1, n2 := 10, 5
	// fmt.Println("hello Go Pro gram")

	// fmt.Println("My name is", name)
	// // fmt.Pintf("Age : %v \n",age)
	// showData("Men Now", age)
	// fmt.Println("Number : ", n1+n2)
	// fmt.Println("......func return value......")

	// // func return value
	// var price = priceTotal()
	// fmt.Println("Price", price)

	// //  func return 2 value
	// fmt.Println(".....func return 2 value.......")
	// name2, lastname2 := name2()
	// fmt.Println("name2", name2)
	// fmt.Println("lastname2", lastname2)

	// fmt.Println(".....func more paramiter.......")
	// var total = summation(10, 20, 3, 4, 5)
	// fmt.Println("Total:", total)

	// fmt.Println(".....Type.......")
	// var product = TProduct{
	// 	name:     "Pen",
	// 	price:    150,
	// 	discount: 10,
	// }
	// product.price = 200

	// fmt.Println("Product :", product.name)
	// fmt.Println("Product :", product.price)
	// fmt.Println("Product :", product.discount)

	// fmt.Println("Before Product setPrice :", product)
	// product = product.setPrice(100)
	// // product = product.Clear()
	// fmt.Println("After Product setPrice :", product)

	// product.Clear()
	// fmt.Println("Clear  Struct:")
	// fmt.Println("Product :", product.name)
	// fmt.Println("Product :", product.price)
	// fmt.Println("Discount :", product.discount)

	// fmt.Println(".....import outside.......")

	// // ชื่อ package.ชื่อ func
	// fmt.Println("Sum from outside:", Plus.Sum(10, 50))
	// fmt.Println("Add from outside:", Plus.Add())

	// showData2()

	// fuForLangth()

	// switchCaseInt()

	// switchCaseText()

	// appendData()

	// objMap()

	// // var book =Plus.Book{name:"dd",price:1150}
	// // book.SetDataBook()

	// var book = books.SetDataBook("Potter", 750)

	// fmt.Println("Book from struct:", book)

	// var aa, bb = swap(1, 5)
	// fmt.Println("Swap", aa, bb)

	// var xx, yy = split(10)
	// fmt.Println("Split X:", xx)
	// fmt.Println("Split Y:", yy)

	// forLoop()
	// ifElse()
	// // pointerTest()
	// arrayTest()

	// fmt.Println("************************************")

	// // pointer
	// pointer.Pointer()
	// pointer.Pointer_setValue()

	// // interface
	// interfacetest.InterFaceTest()

	// // channels
	// channels.Channels()

	// reciver.Reciver()

	// fmt.Println(CalNumber(0))

	// channels.TestSelect()

	// GetDay();

	// fmt.Println(Dd())
	// maptest.CheckArray()
	// maptest.CheckArray2()
	// maptest.CheckArray3()

	// x:= xx{Value: "Jack"}
	// // RecvX(x) //x มันไม่ได้ถูกเปลี่ยนค่า เพราะ ไม่ใช้ pointer
	// RecvX(&x)
	// fmt.Println(x)

	// Dd_Array()
	// Dd_Map()
	// LoopStar()

	// บทความ Go
	// x:=1
	// fmt.Println(x)
	// { // ตั้ง scrop ให้ตัวแปล แบบนี้ก็ได้ เรียก variable shadowing
	// 	x:=2
	// 	fmt.Println(x)
	// }
	// fmt.Print(x)

	// การกำหนดค่า nil
	// var xx string= nil // string มีค่า nil ไม่ได้
	// var xx1 interface{}= nil
	// var xx2 = nil // ตัวแปรไม่ระบุ type เป็นค่า  nil ไม่ได้
	// fmt.Print(xx)
	// fmt.Print(xx1)
	// fmt.Print(xx2)

	// ดูค่า cap ของ map ไม่ได้  ค่า len ด้วยรึป่าว
	// mm:=make(map[int]int,20)
	// cap(mm)
	// len(mm)

	// func แบบไ่มีชื่อ แอนโนนิมันส func มั้ง
	// func(a string)string{ return a }("ss")

	// paramiter ไม่ใช้ pointer ค่าที่ส่งเข้าไป ค่าเดิมจะไม่เปลี่ยนค่า ต่อให้ใสค่าในใน Func

	// เช็ค map ว่ามีค่าของ key ที่มีอยู่รึ ป่าวให้ดึงค่า ok ออกมาด้วย
	// if _,ok := x["one"] ; !ok แสดงว่าไม่มี key นั้น

	// แก้ไขข้อมูล string บางตำแหน่งต้องใช้ []byte()
	// s:= "Russy"
	//s1:= []byte(s)
	// s1[0]= "J"
	// string(s1)
	// output "Jussy"

	//ถ้า ภาษาไทยหรือบางภาษา ต้องใช้ []rune()
	//s1:= []rune(s)
	// s1[0]= "ท"

	// เข้าถึง string ด้วย index
	// d:= "Deek"
	// d[0] จะ แสงค่าอะไรออกมา

	// เช็คค่า string  utf8
	// data:= "Data"
	// data2 := "sd#2"
	// utf8.ValidString(data) / retuen อะไรออกมา
	// utf8.ValidString(data2) / retuen อะไรออกมา

	// เช็คว่ามยาว string
	// dd:= "Dook"
	// len(dd)  จะ ได้ค่าอะไรออกมา

	// log.Fatal and log.Panic จะแสดผลต่างกัน เลือกใช้ให้เหมะสม

	// loop for range type string ดูว่าค่าที่ได้จะเป็นอะไร
	// อาจใช้ for  _,value:= range []byte(...) เทสดูว่าไแบบนี้ได้ค่าอะไร

	// for range ด้วย map ค่าที่แสดงจะ random ไม่เรียง

	// 28
	// struct ที่ขึ้นต้นด้วยตัวเล็ก
	// type Data struct {
	// 	Name string
	// 	age int
	// }
	// dd:= Data{
	// 	Name: "Russy",
	// 	age: 15,
	// }
	// en,_:=json.Marshal(dd)
	// fmt.Println(string(en))

	// d2:=Data{}
	// json.Unmarshal(en,&d2)
	// fmt.Println(d2) // output {Russy 0}

	// 29 go routines
	// ระบบจะไม่รอ go routines ให้เสร็จ
	// ต้องใช้ sync.WaitGroup
	// ตัวอย่างอยู่ใน go rutines ในเว็บ

	// 30 channel
	// 31 ส่งข้อมูลไปหา channel ที่ปิดไปแล้วมันจะ panic

	//32 channel ที่เป็น nil จะเกิด dead rock คืออะไร

	//33
	// type DAta struct {
	// 	Age int
	// 	Name string
	// }

	// // ใช้ * ค่าจะเปลี่ยนตอนเรียก Func
	// func (d DAta) DataTEst1()  {
	// 	d.Age = 8
	// 	d.Name = "Russy"
	// }
	// // ไม่ใช้ * ค่าจะไม่เปลี่ยนตอนเรียก Func
	// func (d *DAta) DataTEst2()  {
	// 	d.Age = 8
	// 	d.Name = "Russy"
	// }

	// d:= DAta{45,"Jack"}
	// fmt.Println(d)
	// d.DataTEst1()
	// fmt.Println(d)
	// d.DataTEst2()
	// fmt.Println(d)
	// output
	// {45 Jack}
	// {45 Jack}
	// {8 Russy}

	// กรณีใช้ net/http
	// 34 ปิด http conect คืออะไรไม่รู้
	// 35 ปิด http response Body น่าจะเป็นการอ่านค่าจาก font end
	// defer.Body.Close()
	// หรือใช้ if เช็คก่อนว่าเป็น nil หรือไม่   if resp != nil

	// 36 json Encoder Adds a Newline Character / แปลงค่า และเอาค่ามาเทียบและอาจจะไม่ตรงกัน
	// json.NewEncodes().Encode()  // คืออะไร

	// 37

	//********* ERROR ***********
	type eerr struct {
		Message string
		Code    int
	}
	ec := errCustom.CustomError{}
	ec.Code = 1150
	ec.Message = "Error Na ja"
	// fmt.Println(ec)// error
	// fmt.Println(ec.Error())// error
	// fmt.Println(ec.Message)// Error Na ja

	// fmt.Println(ec.MessageError2()) // error
	// fmt.Println(ec.MessageError2().Error()) //error

	// เอา struct มาครอบ
	// errr := eerr{}
	// errr.Message = ec.Message
	// errr.Code = ec.Code
	// fmt.Println(errr)

	// fmt.Println(ec) // error
	// fmt.Println(ec.Message)
	// fmt.Println(ec.Code)

	// fmt.Println(ec.MessageError2().Message)
	// fmt.Println(ec.MessageError2().Code)

	// ee:= ec.MessageError1("Na ja",1155)
	// fmt.Println(ee) // error
	// fmt.Println(ee.Message)
	// fmt.Println(ee.Code)

	// err:= er_test1()
	// if err != nil {
	// 	fmt.Println(err) // error
	// 	// fmt.Println(err.Error()) /error
	// }

	// err:= er_test2()
	// if err!= nil {
	// 	fmt.Println(err) // error
	// }

	// err3:= errCustom.MessageError3()
	// if err3 != nil {
	// 	fmt.Println(err3)
	// }

	// err4:= errCustom.MessageError4()
	// if err4 != nil {
	// 	fmt.Println(err4)
	// }

	// echann:= make(chan error)

	// errChan := errCustom.ErrChan{
	// 	ErrChan: echann,
	// }
	// go errCustom.ErrorChan(echann)

	// ปริ้น log ได้เพราะ มี Func Error () string
	// n1,err:=er_test2("15S")
	// if err!= nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(n1)

	//********** JSON**********
	// jsonTest.JsonMarshil1()

	

	// ******  Validation *****
	// validation.ValidaterTest()

	// random number 
	randomnumber.GuessNumber()

}

func er_test1() error {
	ec := errCustom.CustomError{
		Message: "Test Error Now",
		Code:    112,
	}
	return ec
}

func er_test2(s string) (int,error) {
	n,err:=strconv.Atoi(s)
if err != nil {
	// return 0,err
	return 0,errCustom.Err_Russy{
		Message: err.Error(),
		Code: 1112,
	}
}
	return n,nil
}

func LoopStar2() {
	// n1:= []int{3,4,6}

	// for _, num:= range n1 {

	// }
}

func LoopStar() {
	// n1 := 5
	n2 := []int{3, 4, 6}
	result := []string{}
	clear := []string{}
	// result2 := ""

	for index, num := range n2 {
		_ = num
		fmt.Printf("Number = %v \n", num)
		for i := 0; i < n2[index]; i++ {
			result = append(result, "*")
			// fmt.Println(result) // [* * *] ยังมีเครื่องหมาย []

			// ใช้แปลง array ให้ไม่มี []
			var ss string
			ss = strings.Join(result, " ")
			fmt.Println(ss)
		}

		fmt.Println("================")
		result = clear
	}

}

type xx struct {
	Value string
}

// ไม่ใช้ pointer ค่าจะไม่เปลี่ยน
func RecvX(c *xx) {
	c.Value = "Russy"
}

// เช็คเลขซ้ำ Array
// ใช้ array จะมีช่องว่างของ array และมันเริ่มจาก index 0
func Dd_Array() {
	arr := []int{1, 2, 3, 5, 4, 5, 1, 7, 7}
	groupNum := make([][]int, len(arr))

	for _, num := range arr {
		groupNum[num] = append(groupNum[num], num)
	}
	// ลบ index ที่ว่างได้ไหม ****
	// fmt.Println(groupNum) //[[] [1 1] [2] [3] [4] [5 5] [] [7 7] []]

	//  result2  :=make([][]int,10)
	// 	for i, num :=range groupNum{
	// 		if groupNum[i][0] >= 1 {
	// 			result2 = append(result2, num)
	// 		}
	// 	}
	// fmt.Println(result2)

}

// เช็คเลขซ้ำ Map
// ใช้ map จะไม่มี ช่องว่าง น่าจะเหมาะกว่า
func Dd_Map() {
	arr := []int{1, 2, 3, 5, 4, 5, 1, 7, 7}
	groupNum := make(map[int][]int, len(arr))

	// แยกกลุ่มด้วย map
	for _, num := range arr {
		groupNum[num] = append(groupNum[num], num)
	}
	fmt.Println(groupNum) // map[1:[1 1] 2:[2] 3:[3] 4:[4] 5:[5 5] 7:[7 7]]

	// ถอด key ออกแล้วใสค่าใน array ใหม่
	var result [][]int
	for _, num := range groupNum {
		result = append(result, num)
	}
	fmt.Println(result)      // [[5 5] [4] [3] [7 7] [1 1] [2]]
	fmt.Println(len(result)) // [[5 5] [4] [3] [7 7] [1 1] [2]]
}

func D2(value []int) {
	for index := range value {
		value[index] = value[index] + 5
	}
}

/* **************************** */

func CalNumber(n int) (int, error) {
	// เอาค่าตัวเองมา คูณ ก็ได้ด้วย

	if n != 0 {
		return n*n*n - n, nil
	} else {
		return 0, errors.New("Number Value = 0")
	}

}

// Array
func arrayTest() {
	// array แบบกำหนดขนาด
	var ar = [3]int{1, 5, 7}
	var as = [3]string{"", "CMD", "Code"}
	fmt.Println(ar)
	fmt.Println(as)

	//แบบ ไม่กำหนดขนาด
	var n1 = []int{1, 2, 4, 5, 8}
	// append เพิ่มข้อมูลต่อท้าย
	n1 = append(n1, 77)
	fmt.Println("Before slice:", n1)

	// *** slice ตัด array ***
	// ช่องที่ 4 ขึ้นไป ***
	// n1 = n1[4:]
	// 0-4 ***
	// n1 = n1[:4]
	// ระหว่าง 2-4 ***
	n1 = n1[2:4]
	fmt.Println("After slice:", n1)

	var number []int
	number = append(number, 17)
	number = append(number, 7)
	// เพิ่มค่า  array ด้วย array แต่ต้องแตกค่าออกมาด้วย ... ก่อน ***
	number = append(number, []int{1150, 1180}...)

	fmt.Println("Number Array:", number)

}

//if
func ifElse() {
	status := "a dmin"

	if status == "Admin" || status == "admin" {
		fmt.Println("Admin login")
	} else {
		fmt.Println("Not Admin")
	}
}

//for loop แบบต่างๆ
func forLoop() {
	// For ปกติ
	// for i :=0; i<=10; i++ {
	// 	fmt.Println("For", i)
	// }

	// For แบบไม่กำหนดค่าเริ่มต้น
	// ss := 1
	// for ss< 500 {
	// ss+=ss
	// fmt.Println("For", 	ss)

	//  // // ใช้ออกจาก loop
	// // return
	// // break
	// }
	// fmt.Println("For", 	ss)

	// For range
	course := []string{"IOS", "Windows", "Android"}
	for index, item := range course {
		fmt.Println(index+1, "Course :", item)
	}

	//  For range _ ไม่รับค่าตัวแปรนั้น
	// course := []string{"IOS", "Windows", "Android"}
	// for _, item := range course {
	// 	fmt.Println("Course :", item)
	// }
}

// return ค่าโดยๆม่ใสตัวแปร
func split(sum int) (x, y int) {
	x = sum + 10
	y = sum * 10
	return
}

//swap
func swap(a, b int) (int, int) {
	return b, a
}

// ตัวแปร แบบ struct  คล้าย class แต่ใน go ไม่มี  class
type TProduct struct {
	name     string
	price    int
	discount int
}

// Function Struct , เพิ่มฟังชั่นลงไปใน struct TProduct
//func (ตัวแปร Struct) ฟังชั่น() รีเทริน {
func (TP *TProduct) Clear() {
	TP.name = ""
	TP.price = 0
	TP.discount = 0
	// return TP
	// return
}

// clear ไม่ได้
func (TP TProduct) Clear2() TProduct {
	TP.name = ""
	TP.price = 0
	TP.discount = 0
	return TP
}

func (TP TProduct) setPrice(disC int) TProduct {
	// TP.name = ""
	// TP.price = 0
	TP.price = TP.price - disC
	// TP.discount = 0
	return TP
}

//ตัวแปร map Obj
func objMap() {
	var number = map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println("Number Map:", number)
	fmt.Println("Number Map:", number["three"])

	// make สร้าง map ป่าวขึ้นมาก่อน
	var position = make(map[string]string)
	position["IT"] = "It support"
	position["EN"] = "Enginer"
	position["MD"] = "manager"
	fmt.Println("All Position :", position)
	fmt.Println("Position :", position["MD"])

	// map in map, obj ซ้อนกัน
	var course = make(map[string]map[string]int)

	// *** ต้อง make หรือ map ค่าก่อนใสค่าจริงลงไป ****
	course["Android"] = make(map[string]int)
	course["Android"] = map[string]int{"price": 200}

	course["Android"]["code"] = 1111
	course["Android"]["Teacher"] = 333

	course["IOS"] = make(map[string]int)
	course["IOS"]["code"] = 2222
	course["IOS"]["Teacher"] = 444

	// มี Obj แล้วเปลี่ยนค่าได้
	course["Android"]["Teacher"] = 9999
	fmt.Println("All Course", course)
	fmt.Println("Android", course["Android"])
	fmt.Println("IOS", course["IOS"])

}

func appendData() {
	// ไม่จองพื้นที่ mem
	var number []int
	// var n1 = []int{1, 2, 4, 5, 8}
	number = append(number, 17)
	number = append(number, 7)
	// number = append(n1)

	// จอง พื้นที่ mem
	var n2 = make([]int, 0, 5) // len = 0, cap = 5
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)
	n2 = append(n2, 7)

	fmt.Println("Append Data1:", number)
	// fmt.Println("Append Data2:", n2)

	fmt.Println("len=%d , cap=%d , slice=%v", len(number), cap(number), number)
	fmt.Println("len=%d , cap=%d , slice=%v", len(n2), cap(n2), n2)
}

func switchCaseInt() {
	index := 1
	switch index {
	case 0:
		fmt.Println("Case 1")
		break
	case 1:
		fmt.Println("Case 2 ***")
		break
	case 2:
		fmt.Println("Case 3 ***")
		break

	default:
		fmt.Println("*** Default Case ***")
		break
	}

}

func switchCaseText() {
	index := "User"
	switch index {
	case "Manager":
		fmt.Println("Case Manager")
		break
	case "Admin":
		fmt.Println("Case Admin ***")
		break
	case "User":
		fmt.Println("Case User")
		break

	default:
		fmt.Println("*** Default Case No ***")
		break
	}
}

func showData2() {
	fmt.Println("Name Global Variable:", name)
}

func showData(sex string, age int) {
	fmt.Println("Sex:", sex)
	fmt.Println("Age:", age)
}

//   func return
func priceTotal() int {
	return 1150
}

//   func return 2 value
func name2() (string, string) {
	var name2 = "kwan"
	var lastname2 = "tar"
	return name2, lastname2
}

//  func more paramiter
func summation(number ...int) int {
	total := 0

	// for langth
	for _, value := range number {
		total += value
	}
	return total
}

func fuForLangth() {

	course := []string{"IOS", "Windows", "Android"}
	// course2 []string=[]string{"IOS", "Windows", "Android"}
	var n1 = []int{1, 2, 4, 5, 8}

	var n2 [2]int

	n2[0] = 1
	n2[1] = 2

	fmt.Println("Array :", n1)
	fmt.Println("Array :", n2)

	// For ปกติ
	// for i :=0; i<=10; i++ {
	// 	fmt.Println("For", i)
	// }

	// For แบบไม่กำหนดค่าเริ่มต้น
	// ss := 1
	// for ss< 50 {
	// 	ss +=ss
	// }

	// For range
	// for index, item :=range course {
	// 	fmt.Println("Course :",index+1, item)
	// }

	//  For range _ ไม่รับค่าตัวแปรนั้น
	for _, item := range course {
		fmt.Println("Course :", item)
	}
}
