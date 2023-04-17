package main

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"reflect"
// )

// type MyType3 struct {
// 	V3Exported uint32
// }

// type MyType2 struct {
// 	V2         uint16
// 	V3Exported MyType3
// }

// type MyType1 struct {
// 	V1 uint8
// 	V2 MyType2
// 	Ip [4]uint8
// }

// // Encode any struct to []byte
// func Cast(typ string, val uint64) interface{} {
// 	switch typ {
// 	case "byte":
// 		return byte(val)
// 	case "uint":
// 		return uint(val)
// 	case "uint8":
// 		return uint8(val)
// 	case "uint16":
// 		return uint16(val)
// 	case "uint32":
// 		return uint32(val)
// 	case "uint64":
// 		return val
// 	default:
// 		return val
// 	}
// }

// func encodeValue(v reflect.Value, buf *bytes.Buffer) {
// 	k := v.Kind()
// 	if k == reflect.Array || k == reflect.Slice {
// 		for i := 0; i < v.Len(); i++ {
// 			encodeValue(v.Index(i), buf)
// 		}
// 		return
// 	}
// 	if k != reflect.Struct {
// 		binary.Write(buf, binary.BigEndian, Cast(v.Type().String(), v.Uint()))
// 		return
// 	}
// 	for i := 0; i < v.NumField(); i++ {
// 		encodeValue(v.Field(i), buf)
// 	}
// }

// func Bytes(v interface{}) []byte {
// 	buf := new(bytes.Buffer)
// 	encodeValue(reflect.ValueOf(v), buf)
// 	return buf.Bytes()
// }

// // The MustUpdate function recursively checks if each field of
// // a struct is zero(is not set in config file), and if so, sets it to the corresponding value in another struct.
// func MustUpdate(v reflect.Value, vDef reflect.Value) bool {
// 	if v.Kind() != reflect.Struct {
// 		if v.IsZero() {
// 			return true
// 		}
// 	} else {
// 		for i := 0; i < v.NumField(); i++ {
// 			u := MustUpdate(v.Field(i), vDef.Field(i))
// 			if u {
// 				t := vDef.Field(i)
// 				v.Field(i).Set(t)
// 			}
// 		}
// 	}
// 	return false
// }

// // The "UpdateHeader" method uses "MustUpdate" to
// // update the fields of a struct from another struct
// // that have the same type
// func UpdateHeader(h interface{}, tmp interface{}) {
// 	v := reflect.ValueOf(h).Elem()
// 	vNew := reflect.ValueOf(tmp).Elem()
// 	if v.Type() != vNew.Type() {
// 		panic("UpdateHeader: incompatible types")
// 	}
// 	MustUpdate(v, vNew)
// }

// func main() {
// 	x := MyType1{V1: 1, V2: MyType2{V2: 2}, Ip: [4]uint8{192, 168, 1, 1}}
// 	y := Bytes(x)
// 	fmt.Println("the end")
// 	fmt.Printf("% x\n", y)

// 	// Struct copy
// 	xDef := MyType1{Ip: [4]uint8{1, 2, 3, 4}, V1: 11, V2: MyType2{V2: 22, V3Exported: MyType3{V3Exported: 3}}}
// 	UpdateHeader(&x, &xDef)
// 	fmt.Printf("%+v \n", x)
// 	fmt.Printf("%+v \n", xDef)
// }
