package diff

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseFile(filePath, filePath2 string) (map[string]interface{}, map[string]interface{}, error) {
	data1, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось прочитать файл %s: %v", filePath, err)
	}

	var result1 map[string]interface{}
	err = json.Unmarshal(data1, &result1)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка парсинга JSON в файле %s: %v", filePath, err)
	}

	data2, err := os.ReadFile(filePath2)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось прочитать файл %s: %v", filePath2, err)
	}

	var result2 map[string]interface{}
	err = json.Unmarshal(data2, &result2)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка парсинга JSON в файле %s: %v", filePath2, err)
	}

	Difference(result1, result2)

	return result1, result2, nil
}

func Difference(data1, data2 map[string]interface{}) {
	for key, val1 := range data1 {
		val2, exists := data2[key]
		if !exists {
			fmt.Println("-", key, ":", val1)
		} else if val1 != val2 {
			fmt.Println("-", key, ":", val1)
			fmt.Println("+", key, ":", val2)
		}
	}

	for key, val2 := range data2 {
		_, exists := data1[key]
		if !exists {
			fmt.Println("+", key, ":", val2)
		}
	}
}
