package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
)

func StringBuild(args []string) string {
	var s strings.Builder
	for _, arg := range args {
		s.WriteString(arg)
	}
	return s.String()
}

func ConvertToJSON(data interface{}) (map[string]any, error) {
	var result map[string]any
	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, fmt.Errorf("error converting to JSON: %v", err)
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, fmt.Errorf("error converting to JSON: %v", err)
	}
	return result, nil
}
