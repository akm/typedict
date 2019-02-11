package reflectjson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/akm/reflectjson/typedict"
)

func Process(objectMap map[string][]interface{}, ptn *regexp.Regexp) {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := typedict.New(objects).Dig().Types(func(t reflect.Type) bool {
			return ptn.MatchString(t.PkgPath())
		})

		dataTypes := []*DataType{}
		for _, t := range types {
			dt := NewDataType(t)
			dataTypes = append(dataTypes, dt)
		}

		res[key] = dataTypes
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to json.MarshalIndent because of %v\n", err)
		return
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write because of %v\n", err)
		return
	}
}
