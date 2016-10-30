package common

import "reflect"

// StructDeepCopy 定义       未完，内部结构只实现了指针引用
func StructDeepCopy(srcStruct interface{}, desStruct interface{}) {
	if srcStruct == nil || desStruct == nil {
		return
	}

	input := reflect.ValueOf(srcStruct)
	if input.Type().Kind() == reflect.Ptr {
		input = input.Elem()
	}
	inputType := input.Type()
	output := reflect.ValueOf(desStruct)
	if output.Type().Kind() == reflect.Ptr {
		output = output.Elem()
	}
	for i := 0; i < inputType.NumField(); i++ {
		inputField := input.Field(i)
		outputField := output.FieldByName(inputType.Field(i).Name)

		if !outputField.IsValid() {
			continue
		}

		if !outputField.CanSet() {
			continue
		}

		inputFieldType := inputField.Type().Kind()
		outputFieldType := outputField.Type().Kind()
		if inputFieldType == outputFieldType {
			outputField.Set(inputField)
			continue
		}
		if inputFieldType == reflect.Ptr {
			inputField = inputField.Elem()
			inputFieldType = inputField.Type().Kind()
		}
		if outputFieldType == reflect.Ptr {
			outputField = outputField.Elem()
			outputFieldType = outputField.Type().Kind()
		}
		if (inputFieldType == reflect.Struct && outputFieldType == reflect.Struct) || (inputFieldType == reflect.Ptr && outputFieldType == reflect.Ptr) {
			StructDeepCopy(inputField, outputField)
		}
		if inputFieldType == outputFieldType {
			outputField.Set(inputField)
			continue
		}
	}
	return
}
