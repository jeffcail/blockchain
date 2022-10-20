package utils

import "encoding/json"

// JsonToArray
func JsonToArray(str string) ([]string, error) {
	var ss []string
	if err := json.Unmarshal([]byte(str), &ss); err != nil {
		return nil, err
	}
	return ss, nil
}
