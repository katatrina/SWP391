package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors  map[string]string
	GenericError string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v *Validator) IsNoErrors() bool {
	return len(v.FieldErrors) == 0 && v.GenericError == ""
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) AddGenericError(message string) {
	v.GenericError = message
}

func IsNotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func IsStringNotExceedLimit(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func IsStringNotLessThanLimit(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func IsIntInList(value int, list ...int) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}

	return false
}

func IsMatchRegex(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
