package gameroom

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./gameroom.sqlite3?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateNonZero(v interface{}, path string) error {
	vs := reflect.ValueOf(v)
	ts := reflect.TypeOf(v)
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		vs = reflect.ValueOf(v).Elem()
		ts = reflect.TypeOf(v).Elem()
	}

	for i := 0; i < vs.NumField(); i++ {
		if configTag := ts.Field(i).Tag.Get("config"); configTag == "optional" {
			continue
		}
		f := vs.Field(i)
		tag := ts.Field(i).Tag.Get("json")
		if tag == "" {
			continue
		}
		// log.Println(path+"."+tag, "value:", f.Interface())
		name := ts.Field(i).Name
		var err error
		switch f.Kind() {
		case reflect.Int:
			err = MustInt(name, f.Interface().(int))
		case reflect.String:
			err = MustString(name, f.Interface().(string))
		case reflect.Slice:
			switch ts.Field(i).Type.Elem().Kind() {
			case reflect.String:
				err = MustListString(name, f.Interface().([]string))
			default:
				return fmt.Errorf("UNHANDLED SLICE TYPE: %s at %s.%s",
					ts.Field(i).Type.Elem().Kind(),
					path,
					tag)
			}
		case reflect.Struct:
			err = ValidateNonZero(f.Interface(), path+"."+tag)
			// Return early since we already wrapped the error at the lowest level
			if err != nil {
				return err
			}
		case reflect.Ptr:
			if f.IsNil() {
				return fmt.Errorf("%s.%s missing from config", path, tag)
			}
			err = ValidateNonZero(f.Interface(), path+"."+tag)
			// Return early since we already wrapped the error at the lowest level
			if err != nil {
				return err
			}
		case reflect.Map:
			for _, k := range f.MapKeys() {
				err = ValidateNonZero(f.MapIndex(k).Interface(), path+"."+tag+"."+k.String())
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("UNHANDLED TYPE: %s at %s.%s", f.Kind(), path, tag)
		}
		if err != nil {
			return errors.Wrap(err, "At key: "+path+"."+tag)
		}
	}

	return nil
}

// MustString ensures that a string is not empty.
func MustString(key, val string) error {
	if val == "" {
		return fmt.Errorf("key %q is empty", key)
	}
	return nil
}

// MustListString ensures that a list of string values has a length of greater
// than zero, and ensures that each value is not an empty string.
func MustListString(key string, vals []string) error {
	if len(vals) == 0 {
		return fmt.Errorf("list of values %q has zero length", key)
	}
	for _, val := range vals {
		if err := MustString(key, val); err != nil {
			return err
		}
	}
	return nil
}

// MustInt ensures that a value is not zero.
func MustInt(key string, val int) error {
	if val == 0 {
		return fmt.Errorf("key %q has zero value", key)
	}
	return nil
}
