package main

import (
	"fmt"
	
	// "log"
	"time"

	"sync"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(2) // สั่งให้ทำ กี่ครั้ง
	// wg.Done() // ลดการทำงานเมื่อเรียกใช้
	// wg.Wait() // รอให้ทำงานทั้งหมดเสร็จ ตาม Add สั่ง และ Done ลบออก

	// ************* 1 ***********
	// wg.Add(2)
	// wg.Add(2)
	// wg.Done()
	// wg.Done()
	// wg.Add(2)

	// go Goroutine1(&wg)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hi Jack5")
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hi 5")
	// }()

	// wg.Wait()

	// ***************** 2  ***********
	// Go_for()// output ขึ้นอยู่กับ  wg.Add() ว่าจะสั่งให้ทำงานกี่ครั้ง

	// *********  3  *********
	// Go_channal_SumArry() // output ไดเผลรวมอง array 2 ตัว
	// Go_Channal2() // output แสดงค่าออกเรื้อยๆ ถ้าอันไหนทำเสร็จก่อน ด้วย select

	// ***********  4  *********
	// Go_Job_Work() // ยากตรง factorial เพราะมันมีการเรียก  func ตัวเอง มันเลย loop

	// testFactorial(5)
	// fmt.Println(testFactorial(5))

	// ***********  5  *************
	// คำนวณชุดตัวเลข แต่ใช้ go routine เข้ามาช้วย
	// ง่ายกว่าข้างบนเพราะมันไม่มีการ loop ย้อนหลัง
	// แค่ส่งชถดเลขไปพร้อนกันตาม จำนวน worker เช่น 2 ชุดพร้อมกัน หรือ 4 ชุดพร้อมกัน

	// Go_Sum_Number()

	// chatGPT()

	// Chan_use_if()

	// defer fmt.Println("defer ERror") // อยู่ตรงนี้ ออก 2 ชุด
	// if err := testErr(); err != nil {
	// 	fmt.Println("have error")
	// 	return
	// }
	// defer fmt.Println("defer ERror") // อยู๋ตรงนี้  มีแค่ have error

	
}

func testErr() error {
	return fmt.Errorf("Err Func")
}

func Chan_use_if() {
	// cc := make(chan int,10)
	// result := make(chan int,10)

	// print ok
	// go func (c chan int)  {
	// 	for  i:= 0;  i< 10; i++ {
	// 		fmt.Println("CC:",<-c)
	// 	}
	// }(cc)

	// go func(c, result chan int) {
	// 	r:= <-c // chan ไม่สาามรถส่งค่าให้กับ chan ด้วยกันได้
	// 	result <- r
	// }(cc, result)

	// insert ok
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second * 1)
	// 	if i > 5 {
	// 		cc <- 0
	// 	}else {
	// 		cc <- i
	// 	}
	// }

	// nn := make(chan int,10)

	// this if error
	// _, ok := <-nn
	// if !ok {
	// 	fmt.Println("!ok")
	// }else {
	// 	fmt.Println("Have")
	// }

	time.Sleep(time.Second * 10)

}

var ii int

func testFactorial(n int) int {
	// ii++
	// fmt.Println("start i:", ii)
	// fmt.Println(" n:", n)
	// fmt.Println("-----------")
	if n <= 1 {
		time.Sleep(time.Second * 1)
		// n ที่ส่งเข้ามาจะหยุดเมื่อ n <=1 และจะ return ตรงนี้และจบ
		return 1
	} else {
		fmt.Println("n1", n) // 5, 4,3,2
		fmt.Println("***************")
		// time.Sleep(time.Second * 1)

		// ตรงนี้มันเป็นการ call Func อีกครั้งโดยส่ง n-1 เข้าไป
		// การ call จะหยุดเมื่อ n<=1 ตามด้านบน
		// เช่น n =5 * 5+4+3+2+1 = 15

		// แบบนี้ก็ได้
		//ส่ง 5 เข้ามา
		// คำนวณ testFactorial(2)  + testFactorial(2-1)

		//testFactorial(2) | 2+1 =3 // 3 จะไปเป็นตัว บวกของ testFactorial(3)
		//testFactorial(3) | 3+3 = 6 // 6 จะไปเป็นตัว บวกของ testFactorial(4)
		//testFactorial(4) | 4+6 = 10 // 10 จะไปเป็นตัว บวกของ testFactorial(5)
		//testFactorial(5) | 5+ 10 = 15 // ผลลัพธ์

		// ทำไมมันทำ testFactorial(2) ก่อน *****
		// testFactorial(n-1) คือเอาค่าที่คำนวณครั้งที่แล้วมา + กับ n ปัจจุบัน
		// testFactorial(n-1) มันจะต้องย้อนไปถึงค่าเริ่มต้นก่อน แล้วคำนวณกลับมา

		//*****************  ***************
		// เช่น testFactorial(4) n คือ 4 แต่ testFactorial(n-1 มันคือ 3 เลยต้องย้อนกลับไปว่า testFactorial(3) คือค่าอะไร
		// และมันต้องย้อนไปอีกว่า testFactorial(3-1 ซึ่งคือ 2 มันคือค่าอะไร))
		// และจะย้อนไปถึงค่าต่ำสุด คือ 1 และให้ testFactorial(1) return  ค่าออกมาก่อน
		// หรือต้อองรอค่าให้  n-1 มันคำนวค่าออกมาก่อน ถึง  n ปัจจุบันจะทำงานต่อได้
		//************************
		n2 := n + testFactorial(n-1)
		// fmt.Println("n1------1", n) // 2,3,4,5 ทำไมผ่าน กาาคำนวณ  n ถึงตรงข้ามกับ n ด้านบน
		// fmt.Println("-------------------")
		// fmt.Println("n2:", n2) // 3, 6, 10, 15
		return n2
	}
}

func chatGPT() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Message from Channel 2"
	}()

	// go routine ทำงานแค่ครั้งเดียวไม่มี loop for เลยต้องวน 2 รอบ
	for i := 0; i < 2; i++ {
		// time.Sleep(time.Second * 7)
		select {
		case msg1 := <-ch1:
			fmt.Println("MSg1:", msg1)
			// close(ch1)
		case msg2 := <-ch2:
			fmt.Println("Msg2:", msg2)
			// close(ch2)
			// default:
			// 	fmt.Println("default")

		}

	}

}

// *********  1  *************
func Goroutine1(wg *sync.WaitGroup) {
	// var wg sync.WaitGroup

	// wg.Add(2) // ต้องใสให้พอดี กับ go routine ไม่งั้น จะเอาไปทำงานไม่หมด จะทำงานแค่จำนวนที่ระบุ
	go output(wg, "1.Jack")
	go output(wg, "2.Russy")
	go output(wg, "3.John")
	go output(wg, "4.King")

	// wg.Wait()
}

func output(wg *sync.WaitGroup, msg string) {
	defer wg.Done() // ถ้าไม่มี Done() error deadlock เพราะ Done ทำหน้าที่ เอาค่า Add ลบออก
	//  wg.Done()
	fmt.Println(msg)
}

// ***********  2  *********
func Go_for() {
	var wg sync.WaitGroup

	wg.Add(1) // => เลขออกแค่ 1 ตัว

	for i := 0; i < 10; i++ {
		fmt.Println("Start Loop")
		//  wg.Add(1) // 1- 0 แบบไม่เรียง
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}

// **********  3  *************
func Go_channal_SumArry() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{10, 20, 30, 55}

	sumArr1 := make(chan int)
	sumArr2 := make(chan int)

	go func(a1 []int, sum chan int) {
		var result int
		for i, _ := range a1 {
			result = result + a1[i]
		}
		// fmt.Println(result)

		// **** sum คือ sumArr1  คล้าย pointer ส่งค่าในนี้แต่ด้านนอกก็เปลี่ยนค่า****
		// time.Sleep(10 * time.Second)
		sum <- result // ต้องมี go routine ด้วยไม่งั้น error
		close(sum)
	}(arr1, sumArr1) // paramiter ที่ส่งเข้าไป

	// go func(arr2 []int, sum chan int) {
	// 	var result int
	// 	for i, _ := range arr2 {
	// 		result = result + arr2[i]
	// 	}
	// 	// fmt.Println(result)

	// 	// **** sum คือ sumArr2 คล้าย pointer ส่งค่าในนี้แต่ด้านนอกก็เปลี่ยนค่า ****
	// 	sum <- result // ต้องมี go routine ด้วยไม่งั้น error
	// 	close(sum)
	// }(arr2, sumArr2)

	go go_sum(arr2, sumArr2)

	sum1 := <-sumArr1
	sum2 := <-sumArr2
	// _ = sum1
	// _ = sum2

	// nn <- sumArr2
	// n3 := 3
	// n1 := make(chan int)
	// n1 <- n3
	// fmt.Println(<-n1) // error chan เอาไว้ส่งค่าข้าม func กัน ไม่สามารถส่งค่าในท func เดียวกันได้

	// มี for ได้ 0 ตลอด จริงๆมันมีผลรวม แต่มันปล่อยค่าออกไๆปแล้ว และมัน loop ด้วยความเร็วเราเลยไม่เห็น
	//*** น่าจะมัน loop ด้วยความเร็ว จับค่าได้ได้ก็ แสดงผล ******
	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case sum1 := <-sumArr1:
	// 		fmt.Println("sum1", sum1)

	// 	case sum2 := <-sumArr2:
	// 		fmt.Println("sum2", sum2)

	// 	}
	// }

	fmt.Println("sum1:", sum1)
	fmt.Println("-------------")
	fmt.Println("sum2:", sum2)
	// fmt.Println("=======")
	// fmt.Println(<-sumArr2) // chan ที่ส่งค่าไปแล้ว ตัวมันจะไม่มีค่า

}

func go_sum(n []int, sum chan int) {
	var result int

	for i, _ := range n {
		result += n[i]
	}
	// chan ส่งค่าในนี้แต่ chan เดิมก็เปลี่ยนค่า
	// time.Sleep(time.Second * 10) // ใสดีเลย์ไม่มีผล อะไรเลย หรือเป็นที่คอม
	sum <- result
	close(sum)
}

func Go_Channal2() {
	result1 := make(chan string)
	result2 := make(chan string)

	go func() {
		for {
			result1 <- "100 s"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			result2 <- "300 s"
			time.Sleep(time.Second * 3)
		}
	}()

	// for {
	// 	select {
	// 	case <-result1:
	// 		fmt.Println(<-result1)
	// 	case <-result2:
	// 		fmt.Println(<-result2)
	// 	}
	// }

	//******* ถ้าไม่ใช้ for จะส่งค่าครั้งเดียวแล้วจบเลย ******
	fmt.Println(<-result1)
	fmt.Println(<-result2)

}

// *********  4  *************
func Go_Job_Work() {
	now := time.Now()
	// make (chan type, capacity)
	jobs := make(chan int, 20)    // jobs 10 // ตำสุดคือ 5
	results := make(chan int, 20) // result 10 // ตำสุดคือ 5
	// jobs chan -> งานที่ต้องทำ
	// results chan -> เอาไว้รอรับผลลัพธ์จาก jobs
	// fmt.Println(len(jobs)) // 0
	// do work first
	// ....... 3 .........
	go worker(jobs, results) // work
	//  worker(jobs, results) // error ตอนทำงานมันไม่มีค่าอะไรส่งมาดพราะมันอยู่ด้านบน

	// workerNumber := 4
	// for i := 0; i < workerNumber; i++ {
	// 	go worker(jobs, results)
	// }
	// .......... 1 ........
	// if job has receive an i then the result chan will assign a result
	// loop ตามจำนวนความจุ chan
	for i := 0; i < 20; i++ {
		jobs <- i
	}
	// go worker(jobs, results)

	func() {
		// ใสคา่ที่เดียว error เพราะ result loop  จะวน 10 รอบ
		//พอ chan ปล่อยค่าแล้ว result  จะ loop ต่อ เลยทำให้ error
		// ต้องปรับ loop result ให้วน รอบเดียว
		// jobs <- 9

		// ทยอยใสค่าไม่ error
		// time.Sleep(time.Second *1)
		// jobs <- 0
		// time.Sleep(time.Second *1)
		// jobs <- 1
		// time.Sleep(time.Second *1)
		// jobs <- 2
		// time.Sleep(time.Second *1)
		// jobs <- 3
		// time.Sleep(time.Second *1)
		// jobs <- 4
		// time.Sleep(time.Second *1)
		// jobs <- 5
		// time.Sleep(time.Second *1)
		// jobs <- 6
		// time.Sleep(time.Second *1)
		// jobs <- 7
		// time.Sleep(time.Second *1)
		// jobs <- 8
		// time.Sleep(time.Second *1)
		// jobs <- 9
	}()

	// close job chan when assign a value for all job have completed
	// close(jobs)

	// go worker(jobs, results) // work
	//  worker(jobs, results) //work ที่มันทำงานได้เพราะ มันมี jobs ส่งเข้ามาแล้ว

	// loop นี้ใช้รับ print ผลลัพธ์
	// ........ 2 ..........
	// loop 1 รอบ จะไม่มค่าอะไร
	// ........ 5 .......
	// result รอบนี้จะมีค่าออกมา
	// .... 7 ........
	// แสดงผล result อีกครั้ง
	for i := 0; i < 20; i++ { // ถ้า chan ไม่มีค่าปล่อย for มันจะรอ // ถ้า loop เเล้ว chan ไม่มีค่าปล่อยก็ error
		// fmt.Println("for result :", i+1)
		// results มีผลลัพธ์ที่ใสมาจาก func worker
		fmt.Println("REsult: ", <-results)
	}

	// go worker(jobs, results) //วางตรงนี้ error
	//  worker(jobs, results) //วางตรงนี้ error น่าจะ เพราะ results

	// close a result chan when all result have used
	close(results)
	fmt.Println("time elapse:", time.Since(now))
	// time.Sleep(5 *time.Second)
}

// ทำงานครั้งเดียว
func worker(jobs <-chan int, results chan<- int) {
	fmt.Println("worker")
	// jods มีคา่ 0-9 เอามาจาก for ด้านบน
	// results ไม่มีค่า  เพระา result จะรับค่าได้อย่างเดียวจากการประกาศด้านบน chan<- int
	for i := range jobs {
		// .......... 4 ..........
		// results รับ ผลลัพธ์ ****
		// ตรงนี้จะ call func ไปเรื้อยจะกว่า n <=1 factorial() ถึงจะหยุดทำงาน
		results <- factorial(i)
	}
}

var i int // เอาไว้นับรอบการทำงานเฉื้อยๆ
// ...... 6 ..........
// จะมาที่นั้ อีกรอบ
func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		// ดู ตัวดำเนินการด้วยว่าตัวไหนสำคัญ จะทำงานตรงนั้นก่อน
		// e1:= n*(n-1)
		// e2:= n*n-1
		// fmt.Println(e1)
		// fmt.Println(e2)

		// จริงๆมัยคำนวณแบบนี้ (n*n)-1 ตัวคูณทำงานก่อน
		// return n * n-1

		// ตรงนี้ทำให้ factorial  ถูกเรียกซ้ำ ***
		// ตรงนี้มัน loop  เช่น ส่ง 5
		// 5*4*3*2*1

		// ตรงนี้ไป call Func factorial(n-1)
		// n ต้อง <=1 ก่อนถึงจะจบการทำงาน เพราะถ้า  <=1 Func จะ return 1
		return n * factorial(n-1)
	}
}

// *******************   5  *******************
type number struct {
	a int
	b int
}

type result struct {
	sum int
}

//paramiter chan มีสัญลักษณ์ที่ต่างกัน
// <-chan number ส่งได้อย่างเดียวรับไม่ได้
// chan<- result รับได้อย่างเดียวส่งไม่ได้
// เอา jobsCh มาคำนวณ แล้วใสค่าลงใน result **********
func worker_Sum(workId int, jobsCh <-chan number, resultsCh chan<- result, errsCh chan<- error) {
	// var i int

	// jobsCh เป็น slice เราเลย loop มันเพื่อเอาค่าแต่ละตัวออกมาใช้งาน
	for job := range jobsCh {

		// index <- i+1
		// job.a, job.bjob
		fmt.Printf("work_id: %d\n", workId+1)

		// ส่ง job เข้าไปที่ sum เพื่อฟาผลรวม / job มีค่า array number
		sumResult, err := sum(job)
		if err != nil {
			// ถ้ามี error จะส่งไปที่ errsCh
			errsCh <- err

			// แก้ error เพราะถ้าเข้าเงื่อนไขนี้ resultsCh จะมีค่า -1 ออกไป
			//ไม่งั้น loop นี้ resultsCh  จะไม่มีค่าให้ออกไปปล่อยแล้ว error

			// time.Sleep(5 * time.Second) //
			resultsCh <- result{0}
			return
		}
		// ห้ามลืม!!! ยัด nil ลง errsCh ที่มันว่างทกครั้ง ไม่งั้นมันจะค้างงงง
		errsCh <- nil

		// time.Sleep(5 * time.Second) //
		// resultsCh เป็น chan ของ result เลยต้องส่งค่าแบบนี้
		resultsCh <- result{sumResult}

		time.Sleep(2 * time.Second) //
		fmt.Println("***********************************")
	}
}

func sum(num number) (int, error) {
	// return num.a +num.b

	if num.a+num.b == 7 {
		return -1, fmt.Errorf("error!")
	}
	return num.a + num.b, nil
}

func Go_Sum_Number() {
	now := time.Now()
	nums := []number{
		{a: 1, b: 2},
		{a: 4, b: 3},
		{a: 7, b: 2},
		{a: 8, b: 2},
		{a: 7, b: 7},
		{a: 20, b: 2},
		{a: 10, b: 2},
		{a: 15, b: 2},
	}

	// make chan ความจุตาม len(nums)
	jobsCh := make(chan number, len(nums))
	resultsCh := make(chan result, len(nums))
	errsCh := make(chan error, len(nums))

	// ใสค่า nums  ลงใน jobdCh
	for _, v := range nums {
		jobsCh <- v // ใสค่าให้ jobsCh
	}
	close(jobsCh) // close เพื่อไม่ให้รับค่าอีก

	// ตัวนี้กำหนดการทำงานว่าจะทำกี่ครั้งพร้อมกัน อันนี้ ทำ 2 ครั้ง พร้อมกัน
	// คือจะส่งไปกี่ค่า ภายในการทำงาน 1 ครั้งถ้า 2 ก็จะส่งเลข 2 ชุดเข้าไป
	// ถ้า 4 ก็จะส่งเลข 4 ชุดเข้าไป
	numWorkers := 3
	// ประกาศ for ครอบ go routine เพราะ Go จะเข้าใจว่าให้แยก go routine
	// ตามจำนวนของ func ที่จะทำงาน
	for w := 0; w < numWorkers; w++ {
		go worker_Sum(w, jobsCh, resultsCh, errsCh)
	}

	//****** แสดงผลลัพธ์ ***********
	// ตอนเริ่มโปรรแกมมันทำงานตรงนี้แล้ว แต่ยังรอ chan result อยู่
	// result จะรอ chan ของ  result รับค่าเสร็จ
	for r := 0; r < len(nums); r++ {
		// fmt.Println("Reslut Na ja")
		// ถ้า มี errsCh จะส่งไปที่ err เพื่อแสดง log
		err := <-errsCh
		if err != nil {
			// log.Fatal(err)
			// resultsCh <- result{0} // แก้ error ถ้ามี err จากการ sum ต้องใสค่าให้ resultsCh ด้วยจะใสจากตรงไหนก็ได้
			fmt.Println("Error Sum is 7 :", err)
		}
		// result := <-resultsCh

		fmt.Printf("Result: %v \n", <-resultsCh)
		fmt.Println("------------------c")
		// time.Sleep(time.Second*3)
	}

	close(resultsCh)
	close(errsCh)
	fmt.Println("Time: ", time.Since(now))

}
