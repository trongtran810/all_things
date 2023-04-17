package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type MyType3 struct {
	V3Exported uint32
}

type MyType2 struct {
	V2         uint16
	V3Exported MyType3
}

type MyType1 struct {
	V1 uint8
	V2 MyType2
	Ip [4]uint8
}

// Encode any struct to []byte
func Cast(typ string, val uint64) interface{} {
	switch typ {
	case "byte":
		return byte(val)
	case "uint":
		return uint(val)
	case "uint8":
		return uint8(val)
	case "uint16":
		return uint16(val)
	case "uint32":
		return uint32(val)
	case "uint64":
		return val
	default:
		return val
	}
}

var ENCODING_TYPE = binary.BigEndian

func encodeValue(v reflect.Value, buf *bytes.Buffer) {
	k := v.Kind()
	if k == reflect.Array || k == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			encodeValue(v.Index(i), buf)
		}
		return
	}
	if k != reflect.Struct {
		binary.Write(buf, ENCODING_TYPE, Cast(v.Type().String(), v.Uint()))
		return
	}
	for i := 0; i < v.NumField(); i++ {
		encodeValue(v.Field(i), buf)
	}
}

func Bytes(v interface{}) []byte {
	buf := new(bytes.Buffer)
	encodeValue(reflect.ValueOf(v), buf)
	return buf.Bytes()
}

func TestEncoding() {
	x := MyType1{V1: 1, V2: MyType2{V2: 2}, Ip: [4]uint8{192, 168, 1, 0}}
	xBef := MyType1{V1: 1, V2: MyType2{V2: 2}, Ip: [4]uint8{192, 168, 1, 0}}
	y := Bytes(x)
	fmt.Printf("1.Encoding [%v] of [%+v] after using Bytes(): [% x]\n", ENCODING_TYPE, x, y)
	// Struct copy
	xDef := MyType1{Ip: [4]uint8{255, 255, 255, 255}, V1: 11, V2: MyType2{V2: 22, V3Exported: MyType3{V3Exported: 999}}}
	UpdateHeader(&x, &xDef)
	fmt.Printf("2.Update all fields of x[%+v] that is not set yet using the fields of xDef[%+v] => x[%+v]\n", xBef, xDef, x)
}
