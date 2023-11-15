package main

import (
	"time"
	"fmt"
)

func GetDay()  {
	
	year := time.Now().Year()
fmt.Println("Year",year)

	date:= time.Date(year,time.January,1,0,0,0,0, time.Local)
	fmt.Println("date : ",date)

	dayOfWeek := date.Weekday()
	fmt.Println("dateOfWeek : ",dayOfWeek)
	fmt.Println("dateOfWeek INT : ",int(dayOfWeek))

	daysToAdd := 14 - int(dayOfWeek)
	fmt.Println("daysToAdd : ",daysToAdd)

	secondStaurday := date.AddDate(0,0, daysToAdd)

	fmt.Println(secondStaurday)
}