package utils

import "reflect"

// The MustUpdate function recursively checks if each field of
// a struct is zero(is not set in config file), and if so, sets it to the corresponding value in another struct.
func MustUpdate(v reflect.Value, vDef reflect.Value) bool {
	if v.Kind() != reflect.Struct {
		if v.IsZero() {
			return true
		}
	} else {
		for i := 0; i < v.NumField(); i++ {
			u := MustUpdate(v.Field(i), vDef.Field(i))
			if u {
				t := vDef.Field(i)
				v.Field(i).Set(t)
			}
		}
	}
	return false
}

// The "UpdateHeader" method uses "MustUpdate" to
// update the fields of a struct from another struct
// that have the same type
func UpdateHeader(h interface{}, tmp interface{}) {
	v := reflect.ValueOf(h).Elem()
	vNew := reflect.ValueOf(tmp).Elem()
	if v.Type() != vNew.Type() {
		panic("UpdateHeader: incompatible types")
	}
	MustUpdate(v, vNew)
}
