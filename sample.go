package sample

import (
	"reflect"
	"strings"
)

const (
	Int    = -123456
	Uint   = 123456
	Float  = 12345.6
	String = "Foo"
	Bool   = true
	Bytes  = "Zm9vIGJhciBiYXNlCg=="
)

func Sample(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	k := t.Kind()
	for k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	var out interface{}
	switch k {
	case reflect.Interface:
		out = nil
	case reflect.String:
		out = String
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		out = Int
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uint:
		out = Uint
	case reflect.Bool:
		out = Bool
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		out = Float
	case reflect.Struct:
		out = sampleStructPtr(reflect.New(t).Interface())
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		elem := t.Elem()
		if elem.Kind() == reflect.Uint8 {
			return Bytes
		}
		out = []interface{}{Sample(reflect.New(elem).Interface())}
	}
	return out
}

func sampleStructPtr(v interface{}) map[string]interface{} {
	t := reflect.TypeOf(v).Elem()
	numFields := t.NumField()
	out := make(map[string]interface{}, numFields-1)
	for i := 0; i < numFields; i++ {
		sf := t.Field(i)
		if sf.Name == "_" {
			continue
		}
		var name string
		tag := sf.Tag.Get("json")
		if len(tag) > 0 {
			if tag[0] == '-' {
				continue
			}
			if idx := strings.Index(tag, ","); idx > -1 {
				name = tag[:idx]
			} else {
				name = tag
			}
		}
		if name == "" {
			name = sf.Name
		}
		out[name] = Sample(reflect.New(sf.Type).Interface())
	}
	return out
}
