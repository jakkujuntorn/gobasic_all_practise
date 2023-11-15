package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Print("Time Druation:", time.Duration(int64(4)))

	fmt.Println(" ************* Start ************")
	for_Continue()
	fmt.Println(" *************End *************")
}

func for_Continue() {
	condition := true
	count := 0

	for {

		if condition {
			count++
			time.Sleep(1 * time.Second)
		}

		if count >= 10 {
			fmt.Println("Count in IF",count)
			fmt.Println("Break")
			return
		}

		fmt.Println("Count out  IF",count)

		// time.Sleep(1 * time.Second)
		// ใสหรือไม่ใสมีค่าเท่ากัน
		continue

	}

	// fmt.Println("Count in IF",count)

}
