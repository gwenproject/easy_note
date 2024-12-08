package reflect

import (
	"fmt"
	"reflect"
)

// GetValueByField 通过字段名获取字段值
func GetValueByField(v interface{}, fieldName string) (interface{}, error) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if !f.IsValid() {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}
	return f.Interface(), nil
}
