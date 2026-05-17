package gendiff

import (
	"fmt"
	"sort"
)

// GenDiff строит diff между двумя плоскими map[string]interface{}
func GenDiff(data1, data2 map[string]interface{}) string {
	// Собираем все ключи
	keysMap := make(map[string]struct{})
	for k := range data1 {
		keysMap[k] = struct{}{}
	}
	for k := range data2 {
		keysMap[k] = struct{}{}
	}

	// Сортируем ключи
	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := "{\n"

	for _, k := range keys {
		v1, ok1 := data1[k]
		v2, ok2 := data2[k]

		switch {
		case ok1 && ok2:
			if v1 == v2 {
				result += fmt.Sprintf("    %s: %v\n", k, v1)
			} else {
				result += fmt.Sprintf("  - %s: %v\n", k, v1)
				result += fmt.Sprintf("  + %s: %v\n", k, v2)
			}
		case ok1 && !ok2:
			result += fmt.Sprintf("  - %s: %v\n", k, v1)
		case !ok1 && ok2:
			result += fmt.Sprintf("  + %s: %v\n", k, v2)
		}
	}

	result += "}"
	return result
}
