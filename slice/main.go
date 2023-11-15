package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

func main() {
	loc1, _ := time.LoadLocation("Asia/Bangkok")
	fmt.Println(loc1)
	loc2, _ := time.LoadLocation("Asia/Jakarta")
	fmt.Println(loc2)
	// sliceTest1()
	// n1 := []int{1, 5,2, 3, 4, 6,7,9,8}
	// n2 := []int{5,1,2,3,4,6}
	// n3 := []int{1, 2, 4,3, 5, 6, 7}
	// n4 := []int{1, 2, 4,3, 5, 7, 6}
	// _=n1
	// _=n2
	// _=n3
	// _=n4
	// checkSliceChaos(n1)

	//******* HackgeRAng **************
	// array_2dimension()
	// divideArray()
	// Staircase(6)

	// n1 := []int32{1, 2, 3, 4, 5}       // 10 14
	// n2 := []int32{7, 69, 2, 221, 8974} //299 9271
	// _ = n1
	// _ = n2
	// min_max_sum(n2) // เอาข้อมูลมา sort ก่อน เรียงจากน้อยไปมาก

}

func min_max_sum(arr []int32) {
	// sort array  ก่อน
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	total := 0
	for i, v := range arr {
		_ = i
		total = int(v) + total
	}

	min := total - int(arr[len(arr)-1])
	max := total - int(arr[0])

	// fmt.Println("Total:", total)
	fmt.Println(min, max)

}

func Staircase(n int32) {

	//******** แบบนี้ แสดงผลใกล้เคียง แต่ มันอยู่ผิดดด้าน ***********
	// for i := 0; i < n; i++ {
	// 	for ii := 0; ii <= i; ii++ {
	// 		// total = append(total, "#")
	// 		// fmt.Print("#")
	// 		fmt.Printf("%6s\n", strings.Repeat("#", ii))
	// 	}
	// 	fmt.Println("")
	// }

	//******** แบบนี้แก้โจทย์ได้ แต่มันทำงานยังไง **********
	// แสดงผลเหมือนแต่ไม่ผ่าน
	for i := 1; i < int(n)+1; i++ {
		fmt.Println(strings.Repeat(" ", int(n)-i), strings.Repeat("#", int(i)))
	}

	// ของตัวเอง แสดผลเหมือนแต่ไม่ ผ่าน
	// for i := 0; i <= int(n); i++ {
	// 	for ii := int(n); ii > i; ii-- {
	// // 		// fmt.Print(i)
	// // 		// fmt.Print("-")
	// // 		// fmt.Print(ii)
	// // 		// fmt.Print("-")
	// // 		// fmt.Println(int(n))
	// // 		// fmt.Println("-------------")

	// 		fmt.Print(" ")

	// // 		// วนตาม ค่า ii
	// 		if ii-i == 1 {
	// 			for iii := 0; iii < ii; iii++ {
	// 				fmt.Printf("#")
	// 			}
	// 			fmt.Println()
	// 		}

	// 		// fmt.Println() // เว้นมั่วมาก อัน เว้น อันเลย
	// 	}
	// 	// fmt.Println() // มี 2 ช่อง
	// }

	// internet
	// for i := 0; i < int(n); i++ {
	// 	for j := 0; j < int(n)-i-1; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Println()
	// }

}

func array_2dimension() {
	// var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	// fmt.Println(a[0][0])
	// fmt.Println(a[0][1])

	// fmt.Println(a[1][0])
	// fmt.Println(a[1][1])

	// fmt.Println(a[2][0])
	// fmt.Println(a[2][1])

	// fmt.Println(a[3][0])
	// fmt.Println(a[3][1])

	// fmt.Println(a[4][0])
	// fmt.Println(a[4][1])

	// 	var b = [][]int{}
	// b[0][0]= 0
	// b[0][1]= 1
	// b[1][0]= 2
	// b[1][1]= 3
	// b[2][0]= 4
	// b[2][1]= 5
	// b[3][0]= 6
	// b[3][1]= 7

	fmt.Println("**********")
	// fmt.Println(b[0][0])
	// fmt.Println(b[0][1])
	// fmt.Println(b[1][0])
	// fmt.Println(b[1][1])
	// fmt.Println(b[2][0])
	// fmt.Println(b[2][1])
	// fmt.Println(b[3][0])
	// fmt.Println(b[3][1])

	// aa := [][]int{
	// 	{0, 1, 5},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	aa := [][]int{ // ouput 15
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	var a, b, n1, n2 int
	// n2 = 19
	// n1 4
	for i, v := range aa {
		var index = len(aa)

		for ii, _ := range v {
			if i == ii {
				n1 = aa[i][ii]
				n2 = aa[i][index-1]
			}
			index--
		}
		index = len(aa)
		a = a + n1
		b = b + n2
	}
	fmt.Println(a)
	fmt.Println(b)

	// ****** ใช้ math.Abs แปลงค่าติดลบเป็นค่าบวก และได้ค่าเดิม *****
	fmt.Println(math.Abs(float64(a - b)))

	// for i, v := range aa { // v มี ค่า array อีกที
	// 	for _, vv := range v { // เอา v มา loop  อีกรอบ
	// 		fmt.Println(i, ":", vv)
	// 		// output
	// 		// 0 : 0
	// 		// 0 : 1
	// 		// 0 : 2
	// 		// 1 : 4
	// 		// 1 : 5
	// 		// 1 : 6
	// 		// 2 : 7
	// 		// 2 : 8
	// 		// 2 : 9
	// 	}
	// }

}

func divideArray() {
	var arr []int
	// arr = append(arr, -4, 3, -9, 0, 4, 1)
	arr = append(arr, 1, 2, 3, -1, -2, -3, 0, 0)

	side := len(arr)

	var positive int
	var negative int
	var zero int

	// เอา map มาใส ค่า ค่าที่ออกจะสุ่มออกมา ไม่ตรงโจทย์
	// russy := make(map[string]int)

	// แยกค่า บวก มีกี่ตัว
	// ค่าลบมีกี่ตัว
	// ค่า 0 มีกี่ตัว

	// แล้วเอามาหารด้วย len
	for _, v := range arr {

		switch { // switch ไม่ต้องใสค่า
		case v > 0:
			positive++
		case v < 0:
			negative++
		default:
			zero++
		}

		// if arr[i] > 0 && (arr[i]%2) != 0 {
		// 	total = float64(arr[i]) / float64(side)
		// 	s := fmt.Sprintf("%.6f", total)
		// 	fmt.Println(s)
		// }

	}

	p := fmt.Sprintf("%.6f", float64(positive)/float64(side))
	fmt.Println(p)

	n := fmt.Sprintf("%.6f", float64(negative)/float64(side))
	fmt.Println(n)

	z := fmt.Sprintf("%.6f", float64(zero)/float64(side))
	fmt.Println(z)

	// fmt.Println(positive)
	// fmt.Println(negative)
	// fmt.Println(zero)

	// russy["positive"]= positive
	// russy["negative"]= negative
	// russy["zerozero"]= zero
	// fmt.Println(russy)

	// for i := 0; i <= count+1; i++ {
	// 	total = float64(count) / float64(side)
	// 	s := fmt.Sprintf("%.6f", total)
	// 	fmt.Println(s)
	// 	count--
	// }

}

func checkSliceChaos(n []int) {
	// fmt.Println(len(n))

	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			// true
			fmt.Println("wrong:", n[i])
		} else {
			//false
			fmt.Println("true:", n[i])
		}
	}

	// เอาตัวมันเองมา เช็คกับตำแหน่งถัดไปว่าน้อยกว่ารึป่าว
	// ถ้าน้อยกว่าแสดงว่ามันอยู่ผิดตำแหน่ง

}

func sliceTest1() {
	var n [10]int
	var n2 [10]int

	n[0] = 1
	n[1] = 2
	n[2] = 3
	n[3] = 4
	n[4] = 5
	n[5] = 66
	n[6] = 77
	n[7] = 88
	n[8] = 99

	n2 = n
	fmt.Println(n)
	fmt.Println(n2)

	//  slice ตรงนี้มาดึงค่าใน array ไป มันเป็น Pointer
	// slice ไป copy by reference****
	s1 := n[0:5]
	s2 := n[5:]

	// แต่ถ้า copy แบบนี้ เป็น copy by value ********
	s3 := n2

	// copy by Referenec
	// s4:= n2[5:]
	_ = s1
	_ = s2

	// slice เปลี่ยนค่า array เดิม ก็เปลี่ยนค่าตามด้วย
	s1[1] = 777
	s2[0] = 111
	s3[0] = 888
	// s4[1]=9999
	fmt.Println("*** Copy By Reference ******")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("---- After Copy By Reference -----")
	fmt.Println(n)
	fmt.Println("---- After Copy By Value -----")
	fmt.Println(n2)
	fmt.Println(s3)
}
