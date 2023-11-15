package main

import (
	"fmt"
	"strings"

	"encoding/binary"
	"regexp"
)

func main() {
	// pattern1 := "[ru]ssy" // r หรือ u และตามด้วย ssy
	// pattern2 := "[0-9]" // เช็คว่า แค่มีตัวเลขรึไม่
	// pattern3 := `\d`
	// pattern4 := "[s]" // มีตัว s
	// pattern5 := `[ูู^admin]` // มีตัว s
	// pattern6 := "admin|ADMIN" // มีค่า admin ADMIN

	// text1 := "ussy"
	// text2 := "essyerer"
	// text3 := "RUssy"
	// text5 := "admin"
	// text6 := "tony"
	// number1 := "12aa"

	// fmt.Println(Match_String(pattern1, text1))
	// fmt.Println(Match_String(pattern1, text2))
	// fmt.Println(Match_String(pattern1, text3))
	// fmt.Println("**********")
	// fmt.Println(Match_String(pattern4, text1))
	// fmt.Println(Match_String(pattern5, text5))
	// fmt.Println(Match_String(pattern6, text6))

	// fmt.Println("**********")
	// fmt.Println(Match_String(pattern2, number1))
	// fmt.Println(Match_String(pattern3, number1))

	//********************************************************

	fmt.Println("----------------")
	pt1 := "u+"   // มีตัว u
	pt2 := "e+"   // มีตัว e
	pt3 := "use+" // มีตัว use และตามด้วยตัวอะไรก็ได้

	t1 := "russy"
	r1 := regexp.MustCompile(pt1)
	r2 := regexp.MustCompile(pt2)
	r3 := regexp.MustCompile(pt3)
	
	fmt.Println(r1.Match([]byte(t1)))
	fmt.Println(r2.Match([]byte(t1)))
	fmt.Println(r3.Match([]byte(t1)))

	// testRegex()
	// testHelloworld()
	// testEmail()

	// Match()
	// MatchString()
	// fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))
	// Find() // หาคนที่ตรง ชุดเดียว
	// FindAll() // หาคนที่ตรงทั้งหมด
	// FindAllIndex() // หา index ที่ match หาทั้งหมด กับ หาเจอ จุดแรก
	// FindAllString() // หา string ที่ match
	// FindIndex() // หา index ของอักษรที่ match ตัวแรก
	// FindString() // หา string ที่ match
	// FindStringIndex() // หา index ของ string ที่ match
	// Longest() // คืออะไรไม่รู้
	// NumSubexp() // หาจำนวนชุด format รึ ป่าว
	// ReplaceAll() // แทนที้ด้วย string แต่ใช้ []byte ส่ง
	// ReplaceAllLiteralString() // แทนที้ด้วย string แต่ใช้ string ส่ง
	// ReplaceAllString() // แทนที้ด้วย string แต่ใช้ string ส่ง
	// ReplaceAllStringFunc() // ใช้ func กับ format ที่ตั้งไว้
	// Split() // ตัดคำหรืออักษรตาม format

	// ************  มีอีกหลาย Func ที่ยังไม่เข้าใจ *************
}

func Match_String(pattern, text string) bool {
	re := regexp.MustCompile(pattern)
	result := re.Match([]byte(text))
	// result, _ := regexp.Match(pattern, []byte(text))
	return result
}

func testRegex() {

	// text := "asdasrusss"
	// format := "jac" // เอาคำที่ตรงตามนี้เลย

	// text := "asd"
	//  format := "[fgj]" // ตัวใดตัวนึ่ง

	// text := "rys" // false
	// format := "[^russy]" // เอาตัวที่ไม่อยู่ในนี้ งง (rtys/ true) (rys/ false)***

	// text := "asdasjack" // true
	// format := "jack|russy" // เอาค่าใดค่านึ่ง ต้องติดกันไม่อย่างนั้นจะมองเป็นช่องว่าง อยู่ตรงส่วนไหนก็ได้

	// ******** + * ? {m} {m,n} **********
	// + ต้องการตัวอักษาที่อยู่ข้างหน้ากี่ตัว ต้องมีอักษรนั้นอยู่ด้วย ถึงจะ true
	// text := "russy jack"
	// text := "aba"
	// format := "e+" // false
	// format := "z+" // ต้องการตัวอักษาที่อยู่ข้างหน้ากี่ตัว ต้องมีอักษรนั้นอยู่ด้วย ถึงจะ true

	// *   ต้องการอักษรที่อยู่ข้างหน้า 0 ตัวขึ้นไป ทำยังไงก็ true ยกเว้น ช่องว่าง
	// text := "russy" //false
	// text := "r ussy" // true
	// ต้องการอักษรที่อยู่ข้างหน้า 0 ตัวขึ้นไป ทำยังไงก็ true ยกเว้น ช่องว่าง ***
	// format := " a*"

	// ?   ต้องการอักษรที่อยู่ข้างหน้า 0 หรือ 1 ตัว งง มัน true เกือบหมดเลย
	// text := "asdasjack"
	// format := "j?"

	// {m} m ตัว เท่านั้น
	// {m,n} m ตัว ถึง n ตัว
	// text := "aaa" //
	// text := "russy jack" //
	// format := "ja{1}" // ตัวอักษรอยู่ข้างหน้า n ตัว
	// format := "s{1,3}" // ตัวอักษรอยู่ข้างหน้า m-n ตัว 1-3 ตัว

	// ( ) จัดกลุ่ม
	// text := "rusay sys" //
	// format := "(sy)+" // ต้องการ pattren แบบนี้ 1 ครั้งขึ้นไป

	// กลุ่มอักษรพิเศษ
	text := "?" //

	n := 8
	// format := `\d`

	// format := `\d` // ตัวอักษร 0-9
	// format := `\D` // ตัวอักษรไม่ใช้ 0-9
	format := `\w` // ตัวอักษร a-z  A-Z 0-9
	// format := `\W` // ตัวอักษรไม่ใช้ a-z  A-Z 0-9 เช่น เครื่องหมายต่างๆ () ., +-

	bb := make([]byte, 4)
	binary.LittleEndian.PutUint32(bb, uint32(n))
	fmt.Println(bb)

	// ********* ตัวประมวณผล ************
	re := regexp.MustCompile(format)
	r1 := re.Match([]byte(text))
	// r1 := re.Match(bb)
	fmt.Println(r1)
}

func testHelloworld() {
	text1 := "Hello world"
	text2 := "jack world"
	_ = text2
	_ = text1

	format := `^(Hello|hello).*(world|World)$` // $ match จุดสิ้นสุด

	re := regexp.MustCompile(format)
	r1 := re.Match([]byte(text2))

	fmt.Println(r1)
}

func testEmail() {
	text := "55@gmail.co.th"
	format := `\w+@\w+.(com|net|co.th)` // + คือ

	re := regexp.MustCompile(format)
	r1 := re.Match([]byte(text))

	fmt.Println(r1)
}

// ถ้า match จะได้ true
func Match() {
	text := "russy jack"
	format1 := "[hte]"
	re := regexp.MustCompile(format1)

	// ถ้า match จะได้ true
	r2 := re.Match([]byte(text))
	fmt.Println(r2)
}

func MatchString() {
	text := "seafood"
	format1 := "ea.*"
	re := regexp.MustCompile(format1)

	r2 := re.MatchString(text)
	fmt.Println(r2)
}

func Find() {
	text := "seafoodeadd"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.Find([]byte(text))
	fmt.Println(string(r2))
}

func FindAll() {
	text := "seafoodeadd"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.FindAll([]byte(text), -1) // หาทั้งหมด
	for _, v := range r2 {
		fmt.Println(string(v))
	}
	fmt.Println("----------------")
	r3 := re.FindAll([]byte(text), 1) // เจอจุดแรก
	for _, v := range r3 {
		fmt.Println(string(v))
	}
}

// FindAllIndex
func FindAllIndex() {
	text := "sfoodeaddea"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.FindAllIndex([]byte(text), -1) // หาทั้งหมด
	r3 := re.FindAllIndex([]byte(text), 1)  // เจอครั้งแรก
	// for _, v := range r2 {
	// 	fmt.Println(string(v))
	// }
	fmt.Println(r2)
	fmt.Println(r3)
}

// FindAllString
func FindAllString() {
	text := "paranormalass"

	// ********* format ที่จะส่งเข้าไปปเช็ค *******
	format1 := "a."
	// ********* ตัวประมวณผล ************
	re := regexp.MustCompile(format1)

	// เจอค่าตัวแรก ก็ผ่าน
	r1 := re.FindAllString(text, -1) //หาทั้งหมด
	r2 := re.FindAllString(text, 1)  // เจอจุดแรก
	r3 := re.FindAllString(text, 2)  // เอา 2 จุด

	// if len(r1) <= 0 {
	// 	fmt.Println("Text dont match")
	// 	return
	// }

	for _, v := range r1 {
		fmt.Println("Value:", v)
	}
	fmt.Println("---------------")
	for _, v := range r2 {
		fmt.Println("Value:", v)
	}
	fmt.Println("---------------")
	for _, v := range r3 {
		fmt.Println("Value:", v)
	}
}

// คืออะไรไม่รู้
func FindAllStringSubmatch() {

}

func FindIndex() {
	text := "seafoodeadd"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.FindIndex([]byte(text))
	fmt.Println(r2)
}

func FindString() {
	text := "sfoodeadd"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.FindString(text)
	fmt.Println(r2)
}

func FindStringIndex() {
	text := "seafoodeadd"
	format1 := "ea.?"
	re := regexp.MustCompile(format1)

	r2 := re.FindStringIndex(text)
	fmt.Println(r2)
}

func Longest() {
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))
}

func NumSubexp() {
	re0 := regexp.MustCompile(`a.`)
	fmt.Printf("%d\n", re0.NumSubexp())

	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())
}

func ReplaceAll() {
	text := "seafoodeadd"
	format1 := "ea"
	re := regexp.MustCompile(format1)

	r2 := re.ReplaceAll([]byte(text), []byte("*"))
	fmt.Println(string(r2))
}

func ReplaceAllLiteralString() {
	text := "russy jack"
	format1 := "sy"
	re := regexp.MustCompile(format1)

	r2 := re.ReplaceAllLiteralString(text, "*")
	fmt.Println(string(r2))
}

func ReplaceAllString() {
	text := "russy jack"
	format1 := "sy"
	re := regexp.MustCompile(format1)

	r2 := re.ReplaceAllString(text, "*")
	fmt.Println(string(r2))
}

func ReplaceAllStringFunc() {
	text := "russy jack"
	format1 := "s.*"
	re := regexp.MustCompile(format1)

	r2 := re.ReplaceAllStringFunc(text, strings.ToUpper)
	fmt.Println(string(r2))
}

func Split() {
	a := regexp.MustCompile(`a`)
	fmt.Println(a.Split("banana", -1)) // ตัดทุกตัวอักษร
	fmt.Println(a.Split("banana", 0))  // ตัดหมด
	fmt.Println(a.Split("banana", 1))
	fmt.Println(a.Split("banana", 2)) // ตัดตักแรกที่เจอ
	fmt.Println("***********")
	zp := regexp.MustCompile(`z+`)
	fmt.Println(zp.Split("pizza", -1)) //  ตัดทุกตัว
	fmt.Println(zp.Split("pizza", 0))
	fmt.Println(zp.Split("pizza", 1))     // ไม่ตัด
	fmt.Println(zp.Split("pizzazzzy", 2)) // ตัดตัวแรก
}
