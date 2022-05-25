package main

import (
		"fmt"
		"reflect"
)





//查找字符是否在数组中
func InArray(obj interface{}, target interface{}) (bool) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func main() {
		aaa := "123" 
		bbb := []string{"234","34","56","1234"}
		if ! InArray(aaa,bbb) {
			fmt.Println("hahaha")
		} else {
			fmt.Println("no hah")
		}
}
