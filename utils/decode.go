package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

// Decode a byte slice into a struct
func FromBytes(data []byte, v any) error {
	buf := bytes.NewReader(data)
	var ir io.Reader = buf
	return decodeValue(reflect.ValueOf(v).Elem(), ir)
}

func decodeValue(v reflect.Value, buf io.Reader) error {
	k := v.Kind()
	if k == reflect.Array || k == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if err := decodeValue(v.Index(i), buf); err != nil {
				return err
			}
		}
		return nil
	}
	if k != reflect.Struct {
		// Read bytes from buffer and cast to the correct type
		var val uint64
		switch k.String() {
		case "byte":
			var b [1]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = uint64(b[0])
		case "uint":
			var b [8]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = binary.BigEndian.Uint64(b[:])
		case "uint8":
			var b [1]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = uint64(b[0])
		case "uint16":
			var b [2]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = uint64(binary.BigEndian.Uint16(b[:]))
		case "uint32":
			var b [4]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = uint64(binary.BigEndian.Uint32(b[:]))
		case "uint64":
			var b [8]byte
			if _, err := buf.Read(b[:]); err != nil {
				return err
			}
			val = binary.BigEndian.Uint64(b[:])
		default:
			return fmt.Errorf("unsupported type: %s", k.String())
		}
		v.Set(reflect.ValueOf(Cast(k.String(), val)))
		return nil
	}
	// Iterate over struct fields
	for i := 0; i < v.NumField(); i++ {
		if err := decodeValue(v.Field(i), buf); err != nil {
			return err
		}
	}
	return nil
}

func TestDecode() {
	x := MyType1{V1: 1, V2: MyType2{V2: 2}, Ip: [4]uint8{192, 168, 1, 0}}
	z := MyType1{}
	y := Bytes(x)
	fmt.Printf("1.Encoding [%v] of [%+v] after using Bytes(): [% x]\n", ENCODING_TYPE, x, y)
	FromBytes(y, &z)
	fmt.Printf("2.Decoding [%v] of [% x] after using Bytes(): [%+v]\n", ENCODING_TYPE, y, z)
	// Struct copy
}
