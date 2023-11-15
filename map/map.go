package maptest

import (
	"fmt"
	"sort"
)

// 1.
// แยกกลุ่ม array ซ้ำ
// แนวคิด แยก กลุ่มของ ตัวเลขก่อนว่ามีกี่กลุ่มด้วย make(map[int]bool)
// ค่า 1 ค่าจะสร้าง map แค่ 1 อัน
//=========================
// 2.
// loop map ที่แบ่งกลุ่มหาค่าซ้ำแล้ว  for ใน for loop original หาค่าซ้ำ อีกที
// เอาค่า map มี่แยกกลุ่มครั้งแรก  ไปเช็คกับ original ถ้า เท่ากัน ก็ใสค่านี้ใน  array ใหม่
//  array ชุดนั้นจะได้ค่า เท่ากับ กลุ่มที่แยกตั้งแต่แรก เช่น กลมที่ 1 จะมีแค่ค่า [1,1] กลุ่มที่ 2 จะมีแค่ค่า [2]
// จบ loop 1 ครั้งก็ append ลงใน result
func CheckArray() {
	original := []int{3, 1, 2, 3, 4, 1}

	countMap := make(map[int]bool)

	var result [][]int // output [[3 3] [1 1] [2] [4]] array ที่เก็บค่า array

	// 1.
	// เอา countMap ใสค่า true ให้หมด ใส เพื่ออะไร *******
	// แยกกลุ่ม array ก่อน เพราะได้เลข ไม่ซ้ำ
	// สร้างกลุ่ม ตาม ตัวเลข  กลุ่ม 1,2,3,4
	for _, num := range original {
	
		// ทำไมใสค่าเลขไม่ซ้ำกัน
		// เพราะมันเป็น key คล้ายๆของ array 1 index มีได้ 1 ค่า
		countMap[num] = true
	}
	// fmt.Println("Count Map: %v", countMap) // output map[[1:true] [2:true] [3:true] [4:true]]

	// 2.
	// for _,num ไม่ได้เพราะ countMap[int]bool num จะได้ true
	//วน รอบ ตาม countMap (จำนวนเลขที่ไม่ซ้ำ ตามกลุ่มที่แบ่งจาก for ด้านบน)
	for numCountMap := range countMap { //  1,2,3,4 หลังจากแบ่งกลุ่ม
		// fmt.Println("Num in for %v",num) // output 1,2,3,4
		var arrayGroup []int

		for _, value_Origrnal := range original { // {3, 1, 2, 3, 4, 1}
			// fmt.Println("VAlue  in for 2 %v", value)
			// เอากลุ่มที่แบ่งมาเช็ค ตามรอบของจำนวน original
			// loop กลุ่ม ที่ 1,2,3,4 ตามลำดับ
			if numCountMap == value_Origrnal { // เอาเลขมาเช็คว่าตรงกับกลุ่มไหน
			// apped เอาค่าเดิมก่อนหน้าลงไปด้วย
				arrayGroup = append(arrayGroup, value_Origrnal) // เลขที่มาเช็คมีค่าเท่ากับกลุ่มที่แยกไว้
			}
			// fmt.Println("in sert result")
		}

		// พอวนจบ 1 รอบ ก็จะมา append ค่า 1 ครั้ง
		// ทุกครั้งที่ append จะได้ค่าที่เหมือนกับ กลุ่มที่แยกไว้มา append
		// arrayGroup [1,1]
		result = append(result, arrayGroup)
	}

	fmt.Println(result) // output [[3 3] [1 1] [2] [4]] 

}

// 1.
// เอาค่า original
// แยกกลุ่ม array โดยใช้แนวคิด index
// แนวคิด  เอา สร้าง map คล้าย อันแรก
// ค่าใน map  key จะใสค่าที่เท่ากันตาม key เช่น ค่า key 1 จะใสแต่ค่า 1
// output ออกมา คือแยกกลุ่มพร้อม จำนวนข้อมูลอยู่ข้างใน map[1:[1,1] 2:[2,2] 3:[3,3,3,3]]
// 2.
// เอาค่าที่แยก กลุ่มตาม key loop เพื่อถอด key เอาแต่ value มาใส array ใหม่ 
//
func CheckArray2() {
	original := []int{1, 1, 2, 3, 4, 4, 3, 3, 2}

	// อันนี้เก็บ int:[]int / 1:[1,1]
	countMap := make(map[int][]int)

	// ทำไต้องประกาศตัวแปรแบบนี้ เก็บค่า [[1 1] [2 2] [3 3 3] [4 4]]
	// จองขนาด เป็น 0
	// ขยายความจุตาม len(countMap)
	result := make([][]int, 0, len(countMap)) // array ที่บรรจุ array

	// ประกาศแบบนี้ก็ได้ แต่จะไม่จองขนาด
	//   result2 := make([][]int, 0) // array ที่บรรจุ array

	// ประกาศแบบนี้ คือ จองความจุ ไว้ 2 
	// [[] [] [1 1] [2 2] [3 3 3] [4 4]]
	// result := make([][]int, 2) 

	// 1.
	// loop ค่า ที่ส่งเข้ามา เพื่อใสใน map ตามค่า key นั้นๆ
	// เช่น ค่า key 1 ก็จะค่า 1 ลงไปเท่านั้น เพราะมันถูกอ้างด้วย key  ********
	for _, num := range original { // []int{1, 1, 2, 3, 4, 4, 3, 3, 2}
		// ใช้ key ในการอ้างตำแหน่ง เช่น key 1 ก็ใสแต่ค่า 1
		// key 3 ก็ใสแต่ค่า 3
		// key ตำแหน่งนั้นๆ จะใสค่า นั้นค่าเดียว
		countMap[num] = append(countMap[num], num)
	}
	// fmt.Println(countMap) // output map[1:[1,1] 2:[2,2] 3:[3,3,3,3]]

	// 2. 
	// loop เพื่อเอา key ออก เอาแค่ค่า [ ] ใส result เท่านั้น
	for _, nums := range countMap { // nums เป็น array ของ countMap
		// fmt.Print(i)
		// fmt.Println(nums)
		// เอา countMap ที่แยก กลุ่มไว้แล้วมาใสใน result
		// countMap อยู่ในค่า map
		result = append(result, nums)
	}
	fmt.Println(result)
}

// แยกกลุ่ม array ซ้ำ
// แนวคิด เอา array มาเรียงลำดับ
// แล้ว หาช่วงที่เลขเหมือนกัน แล้วตัดแยกเป็นชุดๆ
func CheckArray3() {
	original := []int{1, 1, 2, 3, 4, 4, 3, 3, 2}

	//เรียงลำดับ original จะได้ [1 1 2 2 3 3 3 4 4]
	sort.Ints(original)

	result := make([][]int, 0)

	start := 0

	for i := 1; i <= len(original); i++ { // original  [1 1 2 2 3 3 3 4 4]
		// fmt.Println(original[i])
		// fmt.Println(original[i-1])
		fmt.Println("================================")

		if i == len(original) || original[i] != original[i-1] {
			fmt.Println("IF")
			// ตรงนี้เป็นการ หาระยะ กุล่มเลขที่เหมือนกัน
			group := original[start:i]
			// fmt.Println(group)
			// คล้ายทำแบบนี้
			fmt.Println(original[0:2])
			fmt.Println(original[2:4])
			fmt.Println(original[4:7])
			fmt.Println(original[7:8])

			result = append(result, group)
			start = i
		}

	}
	fmt.Println(result)
}
