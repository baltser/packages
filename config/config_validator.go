package config

import (
	"fmt"
	"reflect"
	"strings"
)

func (c *Config) ensureRequiredFields() error {
	var empty []string

	v := reflect.ValueOf(c)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()

		if fieldValue == "" || fieldValue == 0 {
			empty = append(empty, v.Type().Field(i).Name)
		}
	}
	if len(empty) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.ToUpper(strings.Join(empty, ", ")))
	}
	return nil
}
