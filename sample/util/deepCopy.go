package util

import "reflect"

func DeepCopy[T interface{}](src T) T {
	srcValue := reflect.ValueOf(src)
	srcType := srcValue.Type()

	dst := reflect.New(srcType).Elem()

	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		srcField := srcValue.Field(i)
		dstField := dst.Field(i)

		if field.Type.Kind() == reflect.Ptr && !srcField.IsNil() {
			srcFieldValue := srcField.Elem().Interface()
			dstFieldValue := DeepCopy(srcFieldValue)
			dstField.Set(reflect.New(field.Type.Elem()))
			dstField.Elem().Set(reflect.ValueOf(dstFieldValue))
		} else if field.Type.Kind() == reflect.Struct {
			srcFieldValue := srcField.Interface()
			dstFieldValue := DeepCopy(srcFieldValue)
			dstField.Set(reflect.ValueOf(dstFieldValue))
		} else {
			dstField.Set(srcField)
		}
	}

	return dst.Interface().(T)
}
