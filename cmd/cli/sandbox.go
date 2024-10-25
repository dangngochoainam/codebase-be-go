package main

import (
	"log"
	"reflect"
)

type SubStruct struct {
	sub int
}

type TestStruct struct {
	a   string
	b   []string
	sub SubStruct
}

func GetFieldName(data interface{}) []string {
	fields := make([]string, 0)
	typeOfData := reflect.TypeOf(data)

	for i := 0; i < typeOfData.NumField(); i++ {
		fields = append(fields, typeOfData.Field(i).Name)
	}

	return fields
}

func main() {

	dataTest := TestStruct{
		a:   "aaa",
		b:   []string{"bbb", "ccc"},
		sub: SubStruct{sub: 111},
	}

	log.Println(GetFieldName(dataTest))
}
