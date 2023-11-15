package unidirectional

import "fmt"

func Main_Unidirect() {
	ch := make(chan int)
	done := make(chan string)
	multiple := make(chan int)

	// go sender(ch)
	// go receiver(ch, done)
	// <-done

	// เรียก  done ตรfunc(งนี้ต้องมีการใช้ close ด้วย ไม่งั้น error
	// หลังจาก close done ไปแล้ว ตรงนี้จะได้ค่า default ของ done ที่เป็น bool
	// fmt.Println("", <-done)

	//******************
	go startNumber(ch, done)
	go procress_Number(ch, multiple, done)
	go showNumber(multiple, done)

	// fmt.Println(<-done)

	
	for v:= range done{
		fmt.Println("Done is :",v)
	}

}


// chan สำหรับส่งข้อมูลอย่างเดียว
// คือ รับข้อมูลเข้าตัวเองอย่างเดียว
func sender(ch chan<- int) {
	ch <- 88
	// ถ้าส่งข้อมูลออกจะ error เพราะมีหน้าที่รับอย่างเดียว
	// s:= <- ch
	// fmt.Println(s)
}

//  chan สำหรับ รับข้อมูลอย่างเดียว
// คือ ส่งข้อมูลออกจากตัวเองอย่างเดียว
func receiver(ch <-chan int, done chan string) {
	// ส่งข้อมูลเข้า chan ตรงนี้จะ error
	// ch <- 777
	fmt.Println("Receiver:", <-ch)
	done <- "End"
	close(done)
}

func startNumber(ch chan<- int, done chan<- string) {
	
	for i := 0; i <= 5; i++ {
		// fmt.Println("Start Value:", i)
		ch <- i
	}
	close(ch)
	done <- "startNumber"
}

func procress_Number(ch <-chan int, multi chan<- int, done chan<- string) {

	i := 0
	for v := range ch {
		multi <- v * i
		i++
	}
	close(multi)
	done <- "procress"
}

func showNumber(multi <-chan int, done chan<- string) {
	
	// chan เอามา loop มันจะได้ค่าเดียว ไม่มี index
	for v := range multi {
		fmt.Println("End Value :", v)
		// fmt.Println("Channel Value:", v)
	}
	done <- "end procress"

	// done เป็น sent only close ได้
	close(done)

	//multi เป็น recive only close ไม่ได้
	// close(multi)
}
