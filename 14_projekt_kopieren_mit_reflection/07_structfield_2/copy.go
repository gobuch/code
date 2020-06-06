// Package structfield is used to handle different field
// in structs
package structfield

import (
	"fmt"
	"reflect"
)

// Copy struct data from src to dst, when both structs
// has different fields. It is important that dst is a
// pointer, otherwise no data can be copied.
func Copy(dst, src interface{}) error {
	typeDst := reflect.TypeOf(dst)
	if typeDst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not a pointer")
	}
	valDst := reflect.ValueOf(dst).Elem()
	valSrc := reflect.ValueOf(src)
	typeSrc := reflect.TypeOf(src)
	for i := 0; i < valSrc.NumField(); i++ {
		typeSrcField := typeSrc.Field(i)
		dstField := valDst.FieldByName(typeSrcField.Name)
		if !dstField.IsValid() {
			continue
		}
		if dstField.Kind() != valSrc.Field(i).Kind() {
			// src and dst has differen types
			continue
		}
		srcTag := typeSrc.Field(i).Tag
		if srcTag.Get("structfield") == "nocopy" {
			continue
		}
		dstField.Set(valSrc.Field(i))
	}
	return nil
}
