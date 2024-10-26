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

//func GetFieldName(iface interface{}) []string {
//	// fields := make([]reflect.Value, 0)
//	fields := make([]string, 0)
//	ifv := reflect.ValueOf(iface)
//	ift := reflect.TypeOf(iface)
//
//	for i := 0; i < ift.NumField(); i++ {
//		v := ifv.Field(i)
//
//		switch v.Kind() {
//		case reflect.Struct:
//			switch v.Interface().(type) {
//			case time.Time, sql.NullString, sql.NullTime, gorm.DeletedAt:
//				fields = append(fields, ift.Field(i).Name)
//			default:
//				fields = append(fields, GetFieldName(v.Interface())...)
//			}
//		default:
//			// fields = append(fields, v)
//			// fields = append(fields, v.Type().Field(i).Name)
//			fields = append(fields, ift.Field(i).Name)
//		}
//	}
//
//	return fields
//}
