package wiz

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func extractBaseQuerySQL(i interface{}, table string) string {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic(errors.New("struct required"))
	}
	columns := []string{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dbTag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}
		columns = append(columns, dbTag)
	}
	return fmt.Sprintf("SELECT %s FROM %s ", strings.Join(columns, ","), table)
}
