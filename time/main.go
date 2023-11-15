package main

import (
	"fmt"
	"time"
)

func main() {
	// basicTime()

	// tc := time.Now().Add(29 * time.Minute)
	tc := time.Now()
	tf := time.Now().Add(30 * time.Minute)
	checkTime30(tc, tf)
}

func checkTime30(tc, tf time.Time) {
	fmt.Println(tc.Format("01-02-2006 15:04:05"))
	fmt.Println(tf.Format("01-02-2006 15:04:05"))
	fmt.Println("Time Current",tc.Unix()) // 1697274015
	fmt.Println("Time Future",int(tf.Unix())) // 1697274075

	total := int(tc.Unix()) - int(tf.Unix())

	if int(tc.Unix()) < int(tf.Unix()) {
		fmt.Println("Total :", total)
		
		fmt.Println("In Time 30 Minute")
	} else {
		fmt.Println("Time Out")
	}

}

func basicTime() {
	t1 := time.Now()
	t2 := time.Now()

	t3 := t2.Add(10 * time.Minute)         // เพิ่มเวลา
	t4 := time.Now().Add(10 * time.Minute) // เพิ่มเวลา

	fmt.Println(t1.Format("01-02-2006 15:04:05")) // format => 10-14-2023 14:41:52
	fmt.Println(t1.Format("01-02-2006 15:04:05")) // format => 10-14-2023 14:41:52
	fmt.Println("*********")
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

	fmt.Println("*********")
	fmt.Println(t1.Hour())
	fmt.Println(t2.Hour())
	fmt.Println(t3.Minute())
	fmt.Println(t4.Minute())

	fmt.Println("*********")
	fmt.Println("ti = t2 :", t1 == t2)
	fmt.Println("t1 = t3 :", t1 == t3)
	fmt.Println("t3 = t4 :", t3 == t4)

	fmt.Println("**** Hour *****")
	fmt.Println("ti.Hour = t2.Hour :", t1.Hour() == t2.Hour())
	fmt.Println("t1.Hour = t3.Hour :", t1.Hour() == t3.Hour())
	fmt.Println("t3.Hour = t4.Hour :", t3.Hour() == t4.Hour())

	fmt.Println("**** Minute *****")
	fmt.Println("ti.Minute = t2.Minute :", t1.Minute() == t2.Minute())
	fmt.Println("t1.Minute = t3.Minute :", t1.Minute() > t3.Minute())
	fmt.Println("t3.Minute = t4.Minute :", t3.Minute() > t4.Minute())
}


