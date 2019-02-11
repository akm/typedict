package reflectjson

import (
	"reflect"
	"sort"

	"github.com/akm/reflectjson/typedict"
)

func SeriazlizableWithCategories(objectMap map[string][]interface{}, filters ...func(reflect.Type) bool) map[string][]*DataType {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := typedict.New(objects).Structs(filters...)

		dataTypes := []*DataType{}
		for _, t := range types {
			dt := NewDataType(t)
			dataTypes = append(dataTypes, dt)
		}

		sort.Slice(dataTypes, func(i, j int) bool {
			return (dataTypes[i].PkgPath + "." + dataTypes[i].Name) < (dataTypes[j].PkgPath + "." + dataTypes[j].Name)
		})

		res[key] = dataTypes
	}

	return res
}
