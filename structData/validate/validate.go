package validate

import (
	_ "errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator"
	_"regexp"
)

type Login struct {
	Username string `json:"username" validate:"required,notblank,haveSpace"`
	Password string `json:"password" validate:"required,notblank,haveSpace"`
}

type Address struct { 
	Address string
	Country string
}

type Personal struct {
	Name     string
	LastName string
	Email    string
	Age      int
	Sex      string
}

type dataType1 interface {
	Address | Personal
}

type dataType2 interface {
	Login
}

var validate *validator.Validate

func Validate1[T dataType1](dataUser *T) error {
	return validate.Struct(dataUser)
}

func Validate2[T dataType2](dataUser *T) error {
	err := validate.Struct(dataUser)
	return err
}

func init() {
	fmt.Println("Start Validater")
	validate = validator.New()

	//register custom validation
	validate.RegisterValidation("notblank", NotBlank)
	validate.RegisterValidation("haveSpace", CheckSpace)
	validate.RegisterValidation("checkScript", CheckScript)
}

func StartValidate() {
	fmt.Println("Start Validater")
	validate = validator.New()

	//register custom validation
	validate.RegisterValidation("notblank", NotBlank)
	validate.RegisterValidation("haveSpace", CheckSpace)
}

// custom Validate Not Blank
func NotBlank(fl validator.FieldLevel) bool {
	
	return strings.TrimSpace(fl.Field().String()) != ""
}
func CheckSpace(value validator.FieldLevel) bool {
	for _, v := range value.Field().String() {
		if unicode.IsSpace(v) {
			// fmt.Println("Is Space")
			return false
		}
	}
	return true
}

func CheckScript(fl validator.FieldLevel) bool{


	return true 
}

// น่าจะต้อง reegister func validate เพิ่ม
func ValidateDataUser(dataUser interface{}) error {
	//validate.Struct จะเอา struct ที่ส่งเข้ามาไปเช็คกับ การตั้งค่าของเราว่า validate
	// อะไรบ้าง เช่น validate:"required,notblank"
	err := validate.Struct(dataUser)
	// fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}




func Validate_Var(value string) error {
	return validate.Var(value, "required,min=3,max=10,notblank")
}
