package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// structPointer()

	// MarshalJson1()
	MarshalJson2()

}

func MarshalJson1() {

	data1 := []Data1{
		{
			Name:     "john",
			Lastname: "555",
			Age:      35,
		},
		{
			Name:     "",
			Lastname: "",
			Age:      23,
		},
	}
	_ = data1

	// json.Marshal รับ "struct"  เป็น json แต่จะเป็น []byte ของ json****
	d1, _ := json.Marshal(&data1) // แปลงเป็น []byte ต้องส่งเข้าไปเป็น pointer
	fmt.Println(d1)
	fmt.Println(string(d1))
	fmt.Println("****************")

	// แปลงกลับ จาก []byte ที่เป็น json แล้ว กลับมาเป็น struct
	data2 := []Data1{}                // ส่งมาแบบไหน struct ต้องเป็นแบบนั้น
	err := json.Unmarshal(d1, &data2) // แปลงกลับ รับ []byte ที่แปลงจาก json
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data2)

}

func MarshalJson2() {

	// เอาข้อมูล  json มาใช้ในโปรแกรม
	//****** data json เอามา json.Marshal ได้ค่าไม่ตรงตามต้องการ ****
	// json.Marshal  รับ struct เท่านั้น

	dataJson3 := `{"name":"Kong","last_name":"King","age":37}`
	// dataJson3:=`{"name":"Kong","age":37}`

	// d3,_:=json.Marshal(&dataJson3)
	// fmt.Println(d3)
	// fmt.Println("=============")
	// fmt.Println(string(d3)) //"{\"name\":\"Kong\",\"last_name\":\"King\",\"age\":37}"
	// fmt.Println("=============")
	// fmt.Println("*****************")
	// fmt.Println([]byte(dataJson3))
	// fmt.Println("=============")
	// fmt.Println(string([]byte(dataJson3))) // {"name":"Kong","last_name":"King","age":37}

	data3 := Data1{}
	// ข้อมมูลที่เอามาแปลงเป็น json อยู่ลแล้ว เลยใช้แค่ []byte
	// ** ถ้าเอาข้อมูลที่เป็น json อยู่แล้วไป marshal ค่าที่แปลงกลับจะเพี้ยน
	json.Unmarshal([]byte(dataJson3), &data3)
	fmt.Println(data3)
}

func structPointer() {
	userData1 := Data1{
		Name:     "Jack",
		Lastname: "Five",
	}
	userData2 := Data2{ // age  รับค่า nil struct ต้องเป็น pointer
		Name:     "Jack",
		Lastname: "Five",
		Age:      nil,
	}

	fmt.Println(userData1)
	fmt.Println(userData2)
}

type Data1 struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name,omitempty"`
	Age      int    `json:"age,omitempty"`
}

type Data2 struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	Age      *int   `json:"age",omitempty`
}
