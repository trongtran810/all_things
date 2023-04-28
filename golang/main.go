package main

import (
	"tutorial/utils"
)

type MyType2 struct {
	b int
}
type MyType1 struct {
	a int
	MyType2
}

func main() {
	// utils.WatchFileChange(`D:\TrongTran\meta-projects\.POSPRO\source\meta-vi.com\pospro\golang\test\data`)
	utils.WatchFolderChange("")
	// utils.TestDecode()
	// // Test Encoding
	// utils.TestEncoding()
	// fmt.Println(MyType1{a: 1, MyType2: MyType2{b: 2}})
	// // Get local IP
	// fmt.Println(utils.GetOutboundIP())
}
