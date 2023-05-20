package utils

import "fmt"

var dataList []interface{}

func DemoPointer() {
	dataList = append(dataList, 10)
	dataList = append(dataList, "Hello, World!")

	// Accessing the elements
	for i := range dataList {
		value := dataList[i].(type)
		fmt.Println(value)
	}
}
