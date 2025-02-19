package utils

import (
	"regexp"
	"strings"
)

func TransformCharsForDateofBirth(value *string) {
	*value = regexp.MustCompile(`[^0-9]`).ReplaceAllString(*value, " ")
	*value = strings.TrimSpace(*value)
	*value = regexp.MustCompile(`\s+`).ReplaceAllString(*value, "-")
}

func TransformCharsForPhoneNumber(value *string) {
	*value = regexp.MustCompile(`[^0-9]`).ReplaceAllString(*value, " ")
	*value = strings.TrimSpace(*value)
}
