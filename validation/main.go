package main

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

type UserRegister struct {
	// check  ห้ามว่าง มีอักษรมากกว่า 5 ตัว
	Username string `json:"username" validate:"notblank,textLimit"`

	// custom check format email
	Email string `json:"email" validate:"notblank"`

	// เช็ค  มากกว่า 18 แต่ไม่เกิน 100
	Number int `json:"number" validate:"numberLimit"`
}

type UserData struct {
	// ตัวอักษร ห้ามมีตัวอักษรพิเศษ ไม่เกิน 30 ตัวอัษร
	Name string
	// ตัวอักษร ห้ามมีตัวอักษรพิเศษ ไม่เกิน 30 ตัวอัษร
	Lastname string
	// age >18 and <120 โดยใสแค่ วันเดือนปีเกิด
	Age int `validate:"required,gte=5,lte=130"`
	// เช็ค array address
	Address []*Address
}

type Address struct {
	// เช็คไม่เกิน 120 ตัวอักษร ห้ามมีอักษรพิเศษ
	Address string `json:"address" validate:"required"`
	// เช็ค ้ามีอักษรพิเศษ
	City string `json:"city" validate:"required"`
	// เช็ค ตัวเท่านั้น และตัวเลข 5 ตัวเท่าน้น
	Zipcode int `json:"zipcode" validate:"required"`
}

type Error_Response struct {
	Field string
	Tag   string
}



func main() {
	// app := fiber.New()

	// app.Post("/register", RegisterUser)
	// app.Post("/login", LoginUser)
	// app.Get("/user", User)

	// app.Listen(":8000")

	// ValidaterTest()


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

func RegisterUser(c *fiber.Ctx) error {
	registerUser := UserRegister{}
	if err := c.BodyParser(&registerUser); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error Body parse",
		})
	}

	errRegister := validate.Struct(registerUser)
	if errRegister != nil {

		// errText := strings.Split(errRegister.Error(), "Key:")
		// แยกด้วย \n เพราะ lib ใสมาให้
		errText := strings.Split(errRegister.Error(), "\n")
		errLoop := LoopSplitError(errText)

		return c.Status(500).JSON(fiber.Map{
			"message": "false",
			// "err":     errRegister.Error(),
			"err": errLoop,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    registerUser,
	})
}

func LoginUser(c *fiber.Ctx) error {
	return nil
}

func User(c *fiber.Ctx) error {
	return c.Status(200).JSON("Ok User")
}

//************** Validater lib **********
var validate *validator.Validate

//sample validate struct
func init() {
	validate = validator.New()

	//register custom validation
	validate.RegisterValidation("notblank", NotBlank)
	validate.RegisterValidation("ageVAlidater", AgeValidator)
	validate.RegisterValidation("userNamr_Russy", UsernameValidator)
	validate.RegisterValidation("numberLimit", NumberLimit)
	validate.RegisterValidation("textLimit", textCheckFormat)
}

// custom Validate Not Blank
func NotBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func NumberLimit(fl validator.FieldLevel) bool {

	// fmt.Println(fl.Field().Int())

	if fl.Field().Int() < 18 || fl.Field().Int() > 100 {
		return false
	}

	return true

}

func textCheckFormat(fl validator.FieldLevel) bool {

	// fmt.Println(len(fl.Field().String()))

	//limit len
	tx := strings.TrimSpace(fl.Field().String())
	if len(tx) > 5 {
		return true
	}

	// format type  ใช้ regex

	return false
}

func ValidaterTest() {
	validater := validator.New()

	data := UserRegister{
		Username: "",
		Email:    "sssss.com",
	}

	// validater.RegisterValidation("ageValidate",AgeValidator)
	// validater.RegisterValidation("UValidate", UsernameValidator)

	// dd := strings.Trim("russy@jack","@")
	// fmt.Printf("Data: %v",dd)

	// dd := strings.Trim("                 russy  jack    "," ")
	// fmt.Printf("Data: %v \n",dd)

	// d2:= strings.Trim("!!  J  Ack!!","")
	// fmt.Println(d2)

	// if data.Username == " " {
	// 	fmt.Println("Username is null")
	// }

	err := validater.Struct(&data)

	Error_s := []Error_Response{}

	if err != nil {
		fmt.Println(len(err.(validator.ValidationErrors))) // array ของ error

		for _, value := range err.(validator.ValidationErrors) {

			// fmt.Println(value.Field(),value.Tag())
			// fmt.Println(value.Field()) // struct name Username. Age
			// fmt.Println(value.Tag()) // validator  => required
			// fmt.Println(value) // Key: 'User.Age' Error:Field validation for 'Age' failed on the 'required' tag

			fmt.Println(value.Namespace()) // User.Username, User.Age  | Struct ที่ Error
			fmt.Println("=====================")
			fmt.Println(value.Field()) // Username, Age   | ฟิว ที่ Error
			fmt.Println("=====================")
			fmt.Println(value.StructNamespace()) //User.Username, User.Age   | Struct ที่ Error
			fmt.Println("=====================")
			fmt.Println(value.Tag()) // required, lte  | ประเภทที่ error
			fmt.Println("=====================")
			fmt.Println(value.ActualTag()) // required, let  | ประเภทที่ error
			fmt.Println("=====================")
			fmt.Println(value.Kind()) //string, int  | ประเภทข้อมูล
			fmt.Println("=====================")
			fmt.Println(value.Type()) // string, int | ประเภทข้อมูล
			fmt.Println("=====================")
			fmt.Println(value.Value()) // ว่าง, 150  | ค่าที่ได้มา
			fmt.Println("=====================")
			fmt.Println(value.Param()) // ว่าง,130  | ค่าที่กำหนด
			fmt.Println("=====================")
			fmt.Println(value.Error()) // Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag

			er := Error_Response{
				Field: value.Field(),
				Tag:   value.Tag(),
			}
			Error_s = append(Error_s, er)
		}
	}

	fmt.Println(Error_s)

	// MyEmail := "myegmail.com"
	// errEmail := validater.Var(MyEmail, "required,email")
	// if errEmail != nil {
	// 	fmt.Println(errEmail)
	// }

	// MyAge := 70
	// errAge := validater.Var(MyAge, "required,lte=50")
	// if errAge != nil {
	// 	fmt.Println(errAge)
	// }

}

func AgeValidator(data validator.FieldLevel) bool {
	// strings.Contains(data.Field().String(),"russy")// หาคำว่า russy
	value, ok := data.Field().Interface().(int)
	fmt.Println(value)
	if !ok {
		return false
	}
	fmt.Println(value)
	return value > 1
}

func UsernameValidator(data validator.FieldLevel) bool {
	fmt.Println(data.Field())                               // russy
	return strings.Contains(data.Field().String(), "russy") // หาคำว่า russy

}
