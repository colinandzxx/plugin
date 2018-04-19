package appbase

import "reflect"

func GetNameByType(v interface{}) string {
	refV := reflect.ValueOf(v)
	return reflect.Indirect(refV).Type().String()
}
