package main

import (
	"errors"
	"fmt"
	"sync"

	// "log"
	// "runtime"
	// uni "channels/unidirectional"
	"time"
)

func main() {
	// ch := make(chan int)
	// done := make(chan string)
	// multiple := make(chan int)

	// go startNumber(ch, done)
	// go showNumber(multiple, done)
	// go procress_Number(ch, multiple, done)

	// // fmt.Println(<-done)
	// // fmt.Println(<-done)

	// for v:= range done {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Exit")

	// *********** คอร์ส  Udemy ************************
	// Citricla_Session() // โจทย์นี้ไม่เหมือนในคอร์ส เพราะ มี หลาย go routine แย่งกันไปเรียก Func เดียว น่าจะไปดู ctritcal section

	// buffer_Channel()// Buffer channel

	// Check_Len_Channel()

	// uni.Main_Unidirect()

	//************** รับและส่งค่า chan ถ้าปล่อยค่าไปแล้ว chan นั้นจะไม่มีค่า***********
	// c := make(chan int)
	// go writeToChannel(c, 10)
	// time.Sleep(1 * time.Second) // ไม่มี time sleep ก็ไม่มีอะไรออกมาเพราะ โปรแกรมมันทำงานเสร็จก่อน
	// fmt.Println(<-c)

	//******* print ตรงๆ ******
	// int to channel
	// p1:= 1150
	// p2 := make(chan int,1)
	// p2 <- p1
	// close(p2)
	// fmt.Println(int(<-p2))

	// channel to int
	// var p3 int
	// p4:= make(chan int,2)
	// p4 <-777
	// p3 = <-p4
	// close(p4)
	// fmt.Println(p3)

	//********   Int To Channel   ***********
	// n1 := 10
	// nChan1 := Go_Return_IntTo_Channels(n1)
	// fmt.Println(int(<-nChan1)) // ได้ int
	// fmt.Println(<-nChan1)      // ได้ 0 เพราะมันปล่อยค่าออกไปแล้ว

	// for Channel
	// f1:= make(chan int,4)
	// f2:= 45
	// f1 <- 1150
	// f1 <- 2500
	// f1 <- f2
	// close(f1)
	// for c:=range f1 {
	// 	fmt.Println(c)
	// }

	// ******   channel to Int  *******
	// n1 := 700
	// n11:= 7
	// n2 := make(chan int, 1) // buffer 1 แบบนี้ไม่ error เพราะ ได้ปล่อยค่าไปแล้ว และรับค่าใหม่
	// // n2 := make(chan int)
	// n2 <- n1
	// fmt.Println("Before 1",int(<-n2)) // ได้ 700 เพราะ มันยังไม่คืนค่า ตอนนี้มี 700 กับ 7 คืนค่า 700 ไปแล้ว ใน Func เลยได้ 7
	// n2 <- n11
	// fmt.Println("Before 2",int(<-n2)) // ทำไม่ได้ 700 แต่เข้าไปใน Func  แล้วได้ 7
	// nChan2 := Go_Channels2_toInt(n2)
	// _ = nChan2
	// fmt.Println("Main: ", nChan2)

	//****** Chennel Array ****
	// a1 := make(chan []int, 3)
	// a2 := []int{1, 5, 9}
	// _ = a2
	// // a1 <- a2
	// a1 <- []int{1, 5, 9}
	// Array_Channel(a1)

	// **********  check if *****
	// Check_if("ping")

	// ********  ส่งค่าของ 2 Func  ****************
	// ตัวแปร chan คล้าย goble variable  Func ไหนเอาไปใช้แล้วค่าเปลี่ยน
	// คนที่เอาไปใช้ต่อจะได้ค่านั้น
	// Channels() // some func error check agiant

	// ***********  รับส่งค่า chan ธรรมดา แต่ chan รับค่า struct *****
	// มีรายละเอีอด ถ้าว่าง close กับ ให้ chan ปล่อยค่า
	// มี buffer chan จะแสดงผลผิด ไม่แสดงใน go func เพราะอะไร ****
	// Channels_Struct()

	//********** ให้ go routine ดีเลย์ 1 วิ แล้วรอให้ chan ส่งค่าออกมาถึงทำงานบรรทัดต่อไปได้ ********
	// ถ้า go routine ไม่ส่งค่ามาทาง chan โปรแกรมจะไม่ทำงานต่อและ error
	// fmt.Println("send first")
	// // cc:= <- scheduledNotification(time.Second)
	// // fmt.Println(cc) // output {}
	//  <- scheduledNotification(time.Second)
	// fmt.Println("secondly send")
	// <- scheduledNotification(time.Second)
	// fmt.Println("lastly send")

	// *********** ให้ for loop ไปเรื้อยๆ รอจนกว่า quit จะมีค่าเ้าไปเพื่อให้ หยุด  ****
	// c := make(chan string,10)
	// quit := make(chan string)

	// c <- "Run"
	// // quit <- "QQQ" // ถ้า quit มีค่าจะจบทำงานทันที

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	quit <- "QQQ"
	// }()

	// // ส่ง  c <- "Run" task แรก จะออกค่า run และ c จะถูกใสค่าใหม่เป็น Running...
	// // c ต้องใส buffer ด้วยไม่งั้น error เพราะอะไรไม่รู้ ส่วน quit ไม่ต้องมี buffer
	//  task(c, quit)

	// ************ task2 *************
	// c := make(chan string, 5)
	// quit := make(chan string, 1)

	// c <- "switch"
	// c <- "Run"

	// task2(c,quit) // wait... อย่าเดียวเพราะ  chan อยู่ด้านล่างและ func ไม่ใช้ go routine
	// go task2(c,quit) // ทำงานปกติ

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	// ต้องใสค่า 2 ครั้ง
	// 	c <- "Switch" // ใช้ที่ switch case
	// 	c <- "Runing" // ใช้ แสดงผล
	// 	close(c)
	// }()

	// go func() {
	// 	time.Sleep(7 * time.Second)
	// 	quit <- "Switch"
	// 	quit <- "Now"
	// }()

	// go func() {
	// 	for i := 0; i < 15; i++ {
	// 		c <- "Switch"
	// 		c <- "Run"
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// go task2(c, quit) // work
	// task2(c, quit) // work

	// defer close(c)
	// defer close(quit)

	//****** อีกแบบ
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	quit <- "Switch"
	// 	quit <- "Now"
	// }()
	//  task2(c,quit)

	// time.Sleep(10 * time.Second)

	n := make(chan int)
	err := make(chan string)
	go func() {
		i, errCart := cart()

		if errCart != nil {
			err <- errCart.Error()
		}

		n <- i

	}()

	go func() {
		select {
		case msg := <-err:
			fmt.Println(msg)
		case msg := <-status:
			fmt.Println(msg)
		}
	}()

	cartItem()
	time.Sleep(30 * time.Second)

	fmt.Println("End")

}

var rollBack = make(chan string)
var commit = make(chan string)
var status = make(chan string)

func cart() (int, error) {
	fmt.Println("Begin Cart")

	 select {
	case <-rollBack:
		status <- "Roll Back"
		fmt.Println("rollback Cart")
	case <-commit:
		status <- "Success"
		fmt.Println("commit Cart")
	}


	textError := errors.New("")
	_ = textError
	// return 2, nil
	return 2, errors.New("Cart Error")

}

func cartItem() {
	time.Sleep(10 * time.Second)
	// commit <- "commit"
	rollBack <- "rollBack"
}

func buffer_Channel() {
	fmt.Println("Buffer")

	m := make(chan string, 2)
	done := make(chan string)
	_ = done

	go func(mm chan string) {
		for i := "A"; i <= "G"; {
			fmt.Println("_>", i)
			mm <- i
			//  คืออะไร - ทำให้ตัวอักษรเลื่อน ด้วยรหัสแอสกี้
			i = string([]byte(i)[0] + 1)
		}
		fmt.Println("Close Channel")
		close(mm)
	}(m)

	// go func(mm chan string, d chan string) {
	// 	for i := 0; i < 7; i++ {
	// 		fmt.Println("\t|",<-mm)
	// 	}
	// 	// ปริ้นค่าให้เสร็จก่อน แล้วค้อยใสค่า done
	// 	d <- "Done"
	// 	fmt.Println("Done")
	// }(m, done)

	go showData_Buffer(m, done)

	// โปแกรมจะรอจนกว่าจะมีค่า done ออกมา
	<-done

	// 	for v := range m {
	// 		fmt.Println("VAlue: ", v)
	// 	}
	// <-done
}
func showData_Buffer(mm chan string, do chan string) {
	for i := 0; i < 7; i++ {
		fmt.Println("\t|", <-mm)
	}
	do <- "ShowData"
}

func showData(ch chan int, done chan string) {
	// ด้านบนทำงาน
	// fmt.Println("-------------------------")
	for v := range ch {
		fmt.Println("VAlue:", v)
	}
	// ด้านล่าง ไม่ทำงานเลย ***********
	fmt.Println("-------------------------")
	// done <- "Show Data"
}

func Check_Len_Channel() {
	ch4 := make(chan int, 5)
	done := make(chan string)

	go showData(ch4, done)

	// go func(ch chan int, donE chan bool) {
	// 	for v := range ch4 {
	// 		fmt.Println("VAlue:", v)
	// 	}
	//0 	donE <- true
	// 	// close(done)
	// }(ch4, done)

	for v := range ch4 {
		fmt.Println("Value is :", v)
	}
	fmt.Println("Done Status:", <-done)
	// time.Sleep(10 * time.Second)

}
func insertComparison_Int(ch chan int, done chan string, counter sync.WaitGroup) {
	defer counter.Done()
	for i := 0; i < 5; i++ {
		ch <- 1
	}
	done <- "ch1 succress"

}

func Citricla_Session() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	_ = ch1
	_ = ch2
	_ = ch3
	done := make(chan string, 1) // ไม่มี buffer ก็ยังทำงานได้

	counter := sync.WaitGroup{}

	counter.Add(2)

	go func() {
		counter.Wait()
		close(done)
	}()

	// ch1
	go insertComparison_Int(ch1, done, counter) // ส่งเข้า go routine ไปใสค่า

	// //ch2
	go func(ch chan int, done chan string) {
		defer counter.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		done <- "ch2 succress"
	}(ch2, done) // ส่งเข้า go routine ไปใสค่า

	// ch3
	go func(ch chan int, done chan string) {
		defer counter.Done()
		for i := 0; i < 5; i++ {
			ch <- i * 10
		}
		done <- "ch3 succress"
	}(ch3, done)

	go showData(ch1, done)
	go showData(ch2, done)
	go showData(ch3, done)

	// ปัญหาคือ done เอามา loop ไม่ได้
	// ต้องใช้ sync.WaitGroup มาช้วยจัดการ
	// คือให้ channel มันทำงานให้เสร็จก่อน เพราะไม่งั้น done จะไม่มีค่าให้มา loop
	for v := range done {
		_ = v
		fmt.Println("Done status", v)
	}

	// v, ok := <-done
	// if !ok {
	// 	fmt.Println("No Done")
	// }
	// fmt.Println("OK", v)

	// go func() {
	// 	for v := range done {
	// 		_ = v
	// 		fmt.Println("Done status:", v)
	// 	}
	// }()

	// fmt.Println("Done :", <-done)
	// fmt.Println("Done :", <-done)
	// fmt.Println("Done :", <-done)

	// time.Sleep(1 * time.Second)
	fmt.Println("Succress")
}

func Array_Channel(n chan []int) {
	ar := []int{}
	_ = ar

	// chan ส่งไปหา ar
	// ar = <-n
	// // work
	// for i,v := range ar {
	// 	fmt.Printf("Arrayindex [%v] : %v \n",i, v)
	// }

	// work
	// for i, v := range <-n {
	// 	fmt.Printf("Arrayindex [%v]: %v \n", i, v)
	// }
}

func f1(out chan<- int64, in <-chan int64) {
	// fmt.Println(x)
	// c <- x
}

func f2(out chan int64, in chan int64) {
	// fmt.Println(x)
	// c <- x
}

// ยังไม่ได้เทส
func willClose() {
	willClose := make(chan int, 10)
	willClose <- -1
	willClose <- 0
	willClose <- 2
	<-willClose
	<-willClose
	<-willClose
	close(willClose)
	read := <-willClose
	fmt.Println(read)
}

// ทำ delay 1 วิ
func scheduledNotification(t time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(t)
		ch <- struct{}{}
	}()

	// time.Sleep(t)
	// ch <- struct{}{}

	return ch
}

type emptyStruct = struct{}

// รับส่งค่า chan
func Channels_Struct() {
	// ประกาศแบบนี้ เหมือนกัน chan เป็น struct {}
	// completed := make(chan emptyStruct)
	// completed := make(chan struct{})
	completed := make(chan string, 1)

	go func() {
		fmt.Println("ping")
		time.Sleep(time.Second * 3)
		// <-completed // receive a value from completed channel
		// fmt.Println(<-completed) // ไม่มีค่าอะไรออกมา
		// น่าจะทำหน้าที่คล้าย wg.Done
		<-completed
	}()
	defer close(completed) // ถ้าเปิด
	//  จบการทำงานของ chan แต่ chan นี้รับ struct เลยต้องส่ง struct{}{}
	// completed <- struct{}{} // blocked waiting for a notification
	completed <- "End" //

	fmt.Println("pong")
	fmt.Println(<-completed) // ถ้า coment pong ping แต่ต้องมี buffer ด้วย
	// time.Sleep(time.Second * 5) // coment output pong only
}

// เทส การปล่อยค่า ถ้าปล่อยค่าปแล้วจะ error ถ้ามีการทำงานต่อ
func writeToChannel(c chan int, x int) {
	// 1
	// c <- x
	// fmt.Println("Before", x) // print แต่บางครั้งก็ไม่ print
	// fmt.Println("After", x) //  print แต่บางครั้งก็ไม่ print

	// 2
	fmt.Println("Before", x) // print
	c <- x
	fmt.Println("After", x) //  no print เพราะ chan ปล่อยค่าออกไปแล้ว

	// 3
	// fmt.Println("Before", x) // print
	// fmt.Println("After", x)  //   print
	// c <- x

	close(c)
}

func Go_Return_IntTo_Channels(n1 int) <-chan int {

	n2 := make(chan int, 2)
	n3 := 30

	// แบบ 1
	// ค่าที่ได้จะเป็น ค่าแรกที่ใสค่าลงไป
	// defer close(n2)
	n2 <- n1
	n2 <- n3
	// defer close(n2)

	//  แบบ 2
	// close(n2) //close ตรงนี้ error chan เพราะปล่อยค่าไปแล้วยังเรียกใช้
	// n2 <- n1
	// close(n2) //close ตรงนี้ error chan เพราะปล่อยค่าไปแล้วยังเรียกใช้
	// n2 <- n3
	// close(n2) // close อยู่ตรงนี้ทำงานได้

	//ทำงานได้
	// go func() {
	// 	defer close(n2)
	// 	n2 <- n1
	// }()
	// fmt.Println("in Func",int(<-n2))

	return n2
}

func Go_Channels2_toInt(n2 chan int) int {

	// fmt.Println("In Func1",<-n2)

	var n1 int

	// n1 = <- n2
	// defer close(n2)
	// fmt.Println("n1 in Func",n1)// output 7

	go func() { // n1 ด้านล่างะเป็น 0
		defer close(n2)
		n1 = <-n2
	}()
	fmt.Println("n1 in Func", n1) // output 0

	return n1
}

func task(c, quit chan string) {
	// 	fmt.Println("--------------")
	// fmt.Println("c :",<-c)// -> Run
	// fmt.Println("quit :",<-quit) // -> QQQ
	// fmt.Println("-------------------")

	// ถ้า c มีค่าก็ทำ case แรก ถ้า c ไม่มีค่าก็ทำ case <-quit (c ต้องไม่ใสค่าอะๆรมาเลย)
	for {
		select {
		case c <- "Running...": // อันนี้คืออะไร ****
			time.Sleep(1 * time.Second)
			fmt.Println("Task :", <-c)

		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Println("Waiting ......")
			time.Sleep(1 * time.Second)

		}
	}
}

func task2(c, quit chan string) {
	for {
		select {
		// case c<- "Running...":
		case <-c:
			fmt.Println("Task :", <-c)
			time.Sleep(1 * time.Second)

		case <-quit:
			fmt.Println("Quit", <-quit)
			return
		default:
			fmt.Println("Waiting ......")
			time.Sleep(1 * time.Second)

		}
	}
}

var c = make(chan string, 1)

func Check_if(text string) {
	defer close(c)
	// close(c) // error

	c <- text
	// close(c) //อยู่ตรงนี้ไม่ error ปิดก่อน ส่งค่าได้ แต่ปิดก่อนรับค่าไม่ได้
	str := <-c
	// close(c)
	fmt.Println(str)

	if str == "ping" {
		fmt.Println("pong")
	} else {
		fmt.Println("error")
	}

}

func Channels() {
	//ใช้ 1 core มั้ง
	// runtime.GOMAXPROCS(1)

	// go pong()
	// go russy()
	// go pong()
	// fmt.Println("Step 1")

	// ใสค่าไปใน  channels ใช้ <-
	// เมื่อโยนค่ามาที่ c จะเกิดการ blocking  แล้วจะไปทำงาน go routine อื่นๆ

	// c <- "ping"

	// ถ้าใสไปอีกค่า จะ deadlock เพราะไม่มี go routine
	// c <- "ping"

	// fmt.Println("Step 2")
	// time.Sleep(time.Second)

	cc := make(chan string)
	dd := make(chan string)
	n1 := make(chan int)
	n2 := make(chan int)

	// ส่ง chan เปล่าเข้าไปใน func
	go printString("KKK", cc, n1)
	go printString("John", dd, n2)

	fmt.Printf("%v: %v \n", <-n1, <-cc)
	fmt.Printf("%v: %v \n", <-n2, <-dd)

	// for  Error
	// for msg := range <-cc {
	// 	fmt.Println(msg)
	// 	// fmt.Printf("%v: %v",<-n1,msg)
	// }

	// for Error
	// for msg := range <-dd {
	// 	fmt.Println(msg)
	// 	// fmt.Printf("%v: %v",<-n2,msg)
	// }

	// i := 1
	// go func() {
	// 	for msg := range cc {
	// 		fmt.Printf("%v : %v \n", i, msg)
	// 		i++
	// 	}
	// 	i = 1
	// }()

	// go func() {
	// 	for msg := range dd {
	// 		// fmt.Println(i,msg)
	// 		fmt.Printf("%v : %v \n", i, msg)
	// 		i++
	// 	}
	// 	i = 1
	// }()

	time.Sleep(time.Second * 3)
	fmt.Println("****** End *****")
}

func printString(str string, cc chan string, n chan int) {
	for i := 1; i < 100; i++ {
		// n <- i
		// เอา string ที่ส่งมาใสใน chan
		cc <- str

		// time.Sleep(time.Second * 1)
	}
	cc <- "Done" // จบ loop cc จะได้ค่า Done ออกไปด้วย

	// ปิด chan ไม่งั้นเกิด dead loop
	close(cc)

}
