package main

import "fmt"

func main() {
	// r := r_Func()
	// // ตรงนี้จะเป็นการเรียกต่อเนื่องกัน *******
	// fmt.Println(r())
	// fmt.Println(r())
	// fmt.Println(r())
	// fmt.Println(r())
	// fmt.Println("**********")
	// // ตรงนี้ถือว่าเป็นการเรียกคนละครั้งกัน *********
	// fmt.Println(r_Func()())
	// fmt.Println(r_Func()())

	//************
	flist := []func(){}
	flist = createF([]int{108, 777, 258, 456, 999})

	for _, v := range flist {
		v()
	}


	// เคสนี้ ปกติ *****
	dd:= []int{1,2,4,5,6,7,8,}
	dd = Dod(dd)
	
	for _,v := range dd {
		fmt.Println(v)
	}
}

func Dod(list ...int) []int{
reslut := []int{}

for _,v:=range list{
	reslut = append(reslut,v)
}
	return reslut
}

func createF(list []int) []func() {
	result := []func(){}

	for _, v := range list {
		//ต้องหาตัวแปรมาเก็บค่าด้วย ไม่งั้นมันจะเป็น่าสุดท้าย ค่าเดียว ****
		insideValue:= v
		result = append(result, func() {
			fmt.Println(insideValue)
		})

	}
	// fmt.Println(result) // 999 999 999 ทำไมได้ค่าเดียว
	return result
}

// ทำงานคล้าย เอา finc มา  for ถ้ามีการเรียกหลายครั้ง *****
//anonymous Func ****
func r_Func() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func r_Func2() int {
	var i int
	i++
	return i
}
