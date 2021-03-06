package typedict

import (
	"fmt"
	"reflect"
)

func KeyOf(t reflect.Type) string {
	name := t.Name()
	if name == "" {
		switch t.Kind() {
		case reflect.Array, reflect.Chan, reflect.Ptr, reflect.Slice:
			return KeyOf(t.Elem()) + ":" + t.Kind().String()
		case reflect.Map:
			return fmt.Sprintf("map[%s]%s", KeyOf(t.Key()), KeyOf(t.Elem()))
		default:
			return "no_name_given"
		}
	}

	pkgPath := t.PkgPath()
	if pkgPath == "" {
		return name
	} else {
		return pkgPath + "." + name
	}
}
