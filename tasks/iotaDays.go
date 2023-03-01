package tasks

import (
	"fmt"
	"time"
)

const (
	Sunday time.Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	p2_0 = 1 << iota
	_
	p2_2
	_
	_
	p2_5
)

func IotaDays() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday)
	fmt.Println(p2_0, p2_2, p2_5)
}
