package shared

import "strings"

type Validate interface {
	~string
	IsValid() bool
}

func ValidateEnum[T Validate](input string) *T {
	value := T(strings.ToUpper(strings.TrimSpace(input)))
	if !value.IsValid() {
		return nil
	}

	return &value
}
