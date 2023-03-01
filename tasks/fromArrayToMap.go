package tasks

import "fmt"

type MyStruct struct {
	Name string
	Age  int
}

func FromArrayToMap() {
	arr := []interface{}{"c", MyStruct{Name: "skdjhfj", Age: 23}, "sdfg", 123}
	m := make(map[int]interface{})
	for i, value := range arr {
		m[i] = value
	}
	fmt.Println("Array: ", arr, "\n", "Map: ", m)
}
