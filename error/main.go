package main

import (
	"errors"
	validate "errortest/passwordvalidate"
	"fmt"
)

func main() {
	err := validate.VAlidate_PAssword("1ss1")
	// if err != nil {
	// 	// if ok:=err.(validate.PasswordError); ok == nil {
	// 	// 	fmt.Println("THis PAsswordError")
	// 	// }
	// 	fmt.Println("Error :", err.Error())
	// 	fmt.Println("**********************")
	// 	// ดึงค่าจาก err อีกแบบ ทำเพื่ออะไร*
	// 	fmt.Println("Error :", err.(*validate.PasswordError))
	// }

	if err != nil {
		// เอา error ที่ได้มาเช็ค Type ว่าตรง  กับ Error  ที่เราต้อวการรรึป่าว
		if errors.Is(err, err.(*validate.PasswordError)) {
			fmt.Println(err)
		}


		// ถ้าค่าไม่ตรงมัน Error เลยเหรอ
		// errTest := errors.Is(err, err.(*validate.ErrprTest))
		// fmt.Println(errTest)
		// if errTest {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("Not type PAsswordError")
		// }

	}

	fmt.Println("Done")
}
