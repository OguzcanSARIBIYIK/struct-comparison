package comparison

import (
	"fmt"
	"reflect"
)

type Structs struct {
	Structs []interface{}
}

type Response struct {
	FieldName string
	OldValue  interface{}
	NewValue  interface{}
}

//Comparison two struct by same key name
func Compare(oldStruct interface{}, newStruct interface{}, response *[]Response) {
	strct1 := reflect.ValueOf(oldStruct)
	strct2 := reflect.ValueOf(newStruct)

	for i := 0; i < strct1.NumField(); i++ {
		rVal1 := strct1.Field(i)
		rVal2 := strct2.FieldByName(strct1.Type().Field(i).Name)

		if reflect.ValueOf(rVal1.Interface()).Kind() == reflect.Ptr {
			rVal1 = rVal1.Elem()
		}

		if rVal2.Kind() == reflect.Ptr {
			rVal2 = rVal2.Elem()
		}

		if !checkValid(rVal1, rVal2) {
			continue
		}

		if reflect.ValueOf(rVal1.Interface()).Kind() == reflect.Struct {
			Compare(rVal1.Interface(), rVal2.Interface(), response)
		} else {
			if !reflect.DeepEqual(rVal1.Interface(), rVal2.Interface()) {
				*response = append(*response, Response{
					FieldName: fmt.Sprintf("%v.%v", strct1.Type().Name(), strct1.Type().Field(i).Name),
					OldValue:  rVal1.Interface(),
					NewValue:  rVal2.Interface(),
				})
			}
		}
	}
}

//If any value is not valid this method gonna return false
func checkValid(val1 reflect.Value, val2 reflect.Value) bool {
	if !val1.IsValid() || !val2.IsValid() {
		return false
	}

	return true
}
