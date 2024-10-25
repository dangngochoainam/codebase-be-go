package structhelper

import "reflect"

func GetFieldName(data interface{}) []string {
	fields := make([]string, 0)
	typeOfData := reflect.TypeOf(data)

	for i := 0; i < typeOfData.NumField(); i++ {
		fields = append(fields, typeOfData.Field(i).Name)
	}

	return fields
}
