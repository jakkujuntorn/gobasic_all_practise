package main

import (
	"fmt"
	"strings"
	_"structtest/validate"

	"github.com/gofiber/fiber/v2"
	_ "unicode"
	"regexp"
)

type book struct {
	name  string
	price int
}

type Data struct {
	Name     string
	Lastname string
	Age      int
}

type Idata interface {
	AddData(newData Data)
	EditData(editData Data)
}

func (d Data) AddData(ad Data) {
	fmt.Println(ad)
}

func (d Data) EditData(ed Data) {
	fmt.Println(ed)
}

func main() {
	// validate.StartValidate()
	pt:=`<script[^>]*>([\s\S]*?)<\/script>`
	r1 := regexp.MustCompile(pt)
	
	text:= `<script> let x;x = 6;
	document.getElementById("demo").innerHTML = x;
	</script>`

	text2:=`<div> 1150</div>`
	fmt.Println("Script TEst",r1.Match([]byte(text)))
	fmt.Println("Script TEst",r1.Match([]byte(text2)))


	app := fiber.New()
	_ = app

	// app.Post("/login", func(c *fiber.Ctx) error {
	// 	login := validate.Login{}
	// 	err := c.BodyParser(&login)
	// 	if err != nil {

	// 		return c.JSON(fiber.Map{
	// 			"error": err.Error(),
	// 		})

	// 	}

	// 	// validator *********
	// 	errValidate := validate.ValidateDataUser(login)
	// 	if errValidate != nil {
	// 		// ตัดคำด้วย "\n"
	// 		errText := strings.Split(errValidate.Error(), "\n")

	// 		errLoop := func(text []string) []string {
	// 			var errLoop []string
	// 			for _, value := range text {
	// 				errLoop = append(errLoop, value)
	// 			}
	// 			return errLoop
	// 		}(errText)

	// 		return c.JSON(fiber.Map{
	// 			"error": errLoop,
	// 		})
	// 	}

	// 	return c.SendString(login.Username)
	// })

	// app.Listen(":8000")

	// name := "Jassssssss1"
	// name := "     "
	// err:=validate.Validate_Var(name)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Println("Trim: ", strings.Trim("!#!Hello world##&#", " ")) //เอาตัวอักษรที่เรากำหนดออก => Hello world##&
	// fmt.Println("TrimPrefix: ",strings.TrimPrefix("##Jack$$ world"," "))// เอาตัวอักษรที่กำหนดออก ต้องเหมือนที่เรากำหนด  =>  Jack$$
	// fmt.Println("TrimSpace: ",strings.TrimSpace("          John 5   Sulatar     "))
	// fmt.Println("Trim: ",strings.Trim("          John 5   Sulatar     "," "))
	// fmt.Println("TrimSuffix: ",strings.TrimSuffix("Hello world","world")) // เอาคำออกตามรูปแบบที่เรากำหนด => Hello
	// fmt.Println("TrimLeft: ",strings.TrimLeft("  Hello world"," "))

	// strings.
	fmt.Println("****************")

	// ss:="jack hight"
	// ss:="jackhight"

	// for _,value:=range ss{

	// 	if unicode.IsSpace(value){
	// 		fmt.Println("Have Space")
	// 		return
	// 	}
	// }
	// fmt.Println("Not Space")

	// dd := Data{}

	// bb := Data{
	// 	Name:     "JAckBB",
	// 	Lastname: "Five",
	// 	Age:      39,
	// }

	// dd.AddData(bb)

	// login := validate.Login{
	// 	Username: "John",
	// 	Password: "1150",
	// }
	// address := validate.Address{
	// 	Address: "1150",
	// 	Country: "Bangkok",
	// }
	// _ = address

	// err1 := validate.ValidateDataUser(&login)
	// // fmt.Println(err1)
	// if err1 != nil {
	// 	fmt.Println("****Type 1 Not Format******")
	// 	fmt.Println(err1.Error())
	// }
	// fmt.Println("Type 1 Complete")

	// err2 := validate.ValidateDataUser(&address)
	// // fmt.Println(err2)
	// if err2 != nil {
	// 	fmt.Println("****Type 2 Not Format******")
	// 	fmt.Println(err2.Error())
	// }
	// fmt.Println("Type 2 Complete")

	// ด้านล่าง ถ้าใสผิดtype มัน error เลย
	// err11 := validate.Validate1(&address)
	// if err11 != nil {
	// 	fmt.Println("****Type 1 Not Format******")
	// 	// fmt.Println(err1.Error())
	// }
	// fmt.Println("Type 1 Complete")

	// err22 := validate.Validate2(&login)
	// if err22 != nil {
	// 	fmt.Println("Type 1 Not Format")
	// 	fmt.Println(err22.Error())
	// }
	// fmt.Println("Type 2 Complete")

}

// แยก text error สำหรับ validater
func LoopSplitError(text []string) []string {
	var errLoop []string
	// fmt.Println(text)

	for _, value := range text {
		errLoop = append(errLoop, value)
	}

	// chat GPT
	// มันเริ่มที่ตำแหน่ง 1 เพราะ ตำแหน่ง 0 มันเป็นค่าว่างมาอยู่ แล้ว ****
	// for _, value := range text[1:] {
	// 	// value = strings.TrimSpace(value)
	// 	errLoop = append(errLoop, value)
	// }

	return errLoop
}

func SetDataBook(nameB string, priceB int) book {

	// if (nameB == "" ) {
	// 	var message = "Data is Emptry"
	// 	return message
	// } else {
	// 	var b = book{name: nameB, price: priceB}
	// 	return b
	// }

	var b = book{name: nameB, price: priceB}
	return b

}

// ใช้ pointer ช้วยเช็คการสนร้าง Obj ด้วย  nil เช็คว่ามีการสร้างแล้วรึยัง จะได้ค่าตัวแรกที่สร้าง  วีดีโอ 34 point จะเป็นตัวบอกตำแหน่งถ้าสร้างแล้ว pointer  จะมีตำแหน่ง
