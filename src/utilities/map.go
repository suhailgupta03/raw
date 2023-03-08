package utilities

import "encoding/json"

func GetAllKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// Stringify Returns the JSON for the map input
func Stringify[K comparable, V any](m map[K]V) string {
	b, _ := json.Marshal(m)
	return string(b)
}
