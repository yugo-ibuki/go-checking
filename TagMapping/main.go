package main

import (
	"log"
	"reflect"
	"strconv"
)

func main() {
	type MapStruct struct {
		Str     string  `map:"str"`
		StrPtr  *string `map:"str"`
		Bool    bool    `map:"bool"`
		BoolPtr *bool   `map:"bool"`
		Int     int     `map:"int"`
		IntPtr  *int    `map:"int"`
	}
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
		"int":  "12345",
	}
	var ms MapStruct
	Decode(&ms, src)
	log.Println(ms)
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
}

func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		sv, ok := src[key]
		if !ok {
			continue
		}

		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			if k == reflect.Ptr {
				continue
			}
			isP = true
		}
		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err == nil {
				if isP {
					e.Field(i).Set(reflect.ValueOf(&b))
				} else {
					e.Field(i).SetBool(b)
				}
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err == nil {
				if isP {
					n := int(n64)
					e.Field(i).Set(reflect.ValueOf(&n))
				} else {
					e.Field(i).SetInt(n64)
				}
			}
		}
	}
	return nil
}
