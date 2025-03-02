package utils

import (
	"log"
	"reflect"

	"github.com/spf13/viper"
)

func InitParams(v any) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		log.Fatal("InitParams requires pointer to struct as argument")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		log.Fatal("InitParams requires pointer to struct as argument")
	}

	for i := 0; i < rv.NumField(); i++ {
		fieldType := rv.Type().Field(i)
		tag := fieldType.Tag.Get("rstax")
		if tag == "" {
			continue
		}

		field := rv.Field(i)
		switch fieldType.Type.Kind() {
		case reflect.String:
			if field.String() == "" {
				field.SetString(viper.GetString(tag))
			}
		}
	}
}
