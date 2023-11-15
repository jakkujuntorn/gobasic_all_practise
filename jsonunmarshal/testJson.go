package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DataUser struct {
	Name string
	Age  int
}

type Data2 struct {
	Name string
	Age  int
}

func main() {
	jsonMarshalIndent()
	// jsonMarchil_Normal()
}

func jsonMarshalIndent() {
	d := DataUser{
		Name: "Jack",
		Age:  39,
	}
	// แค่จัดรูปแบบการแสดง
	js1, _ := json.MarshalIndent(d, "", "\n")
	js2, _ := json.MarshalIndent(d, "", "   ")

	fmt.Println(js1)
	fmt.Println("--------------")
	fmt.Println(string(js1))
	fmt.Println(bytes.NewBuffer(js1))

	fmt.Println("***************************")
	fmt.Println(js2)
	fmt.Println("--------------")
	fmt.Println(string(js2))

}

func jsonMarchil_Normal() {
	d := DataUser{
		Name: "Jack",
		Age:  39,
	}
	js1, _ := json.Marshal(d)

	fmt.Println(js1)
	fmt.Println("--------------")
	fmt.Println(string(js1))

}

func JsonMarshil1() {
	myData := `{"Name":"Russy","Age":38}`
	_ = myData

	data2 := Data2{ // แปลงเป็น []byte ไม่ได้
		Name: "JAck",
		Age:  17,
	}
	_ = data2

	data3 := "63" // แปลงเป็น []byte ได้
	_ = data3
	dd := DataUser{}
	_ = dd

	n1 := []byte("6")
	n2 := []byte("7")
	n3 := []byte("8")
	n4 := []byte("9")
	n5 := []byte("0")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)

	// n6:= []byte("880")
	// n66:= []byte("159")
	// fmt.Println("===========")
	// // nn:= []byte("92")
	// fmt.Println(n6)
	// fmt.Println(n66)
	// // fmt.Println(nn)

	err := json.Unmarshal([]byte(myData), &dd)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(dd)

	// bu, _ := json.Marshal(dd)
	// fmt.Println(bu)

	// json.Unmarshal(bu, &dd)
	// fmt.Println(dd)

}
