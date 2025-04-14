package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func EncodedStructFields(s any, tag string, setValue func(key string, value []byte) error) error {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Pointer && v.Kind() != reflect.Interface {
		return errors.New("must be pointer or interface of struct 1")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("must be pointer or interface of struct")
	}

	typeOfS := v.Type()
	for i := range typeOfS.NumField() {
		val_bytes, err := json.Marshal(v.Field(i).Interface())
		if err != nil {
			return err
		}
		err = setValue(typeOfS.Field(i).Tag.Get(tag), val_bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func DecodeStructFields(s any, tag string, getValue func(key string) ([]byte, error)) error {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Pointer && v.Kind() != reflect.Interface {
		return errors.New("must be pointer or interface of struct 1")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("must be pointer or interface of struct")
	}

	typeOfS := v.Type()
	for i := range typeOfS.NumField() {
		byte_value, err := getValue(typeOfS.Field(i).Tag.Get(tag))
		if err != nil {
			return err
		}
		if len(byte_value) == 0 {
			continue
		}

		err = json.Unmarshal(byte_value, v.Field(i).Addr().Interface())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}
