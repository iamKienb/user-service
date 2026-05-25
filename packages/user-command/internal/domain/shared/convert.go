package shared

import (
	"strings"

	"github.com/google/uuid"
)

func ParseToRawID[T ~[16]byte](ID string) (T, error) {
	var result T
	parts := strings.Split(ID, "_")
	idStr := parts[len(parts)-1]

	parsed, err := uuid.Parse(idStr)
	if err != nil {
		return result, err
	}

	return T(parsed), nil
}

func Strings[T ~string](items []T) []string {
	result := make([]string, len(items))
	for i, v := range items {
		result[i] = string(v)
	}
	return result
}

func FromStrings[T ~string](items []string) []T {
	result := make([]T, len(items))
	for i, v := range items {
		result[i] = T(v)
	}
	return result
}
