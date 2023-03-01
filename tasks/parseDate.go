package tasks

import (
	"fmt"
	"time"
)

func ParseDate() {
	myDate1 := time.Now()
	formattedDate := myDate1.Format("January 02, 2006 15:04:05")
	fmt.Println(formattedDate)
	parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")
	if err != nil {
		panic(err)
	}
	fmt.Println(parseTime)
}
