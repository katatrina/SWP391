package validator

import (
	"regexp"
)

type Validator struct {
	FieldErrors  map[string]string
	GenericError string
}

var (
	EmailRX   = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	PhoneRX   = regexp.MustCompile("^0\\d{9}$")
	TaxCodeRX = regexp.MustCompile(`\d{3}-\d{2}-\d{5}`)
)

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

func IsMatchRegex(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func IsAddressTooShort(value string) bool {
	return len(value) < 10
}

func IsPasswordTooShort(value string) bool {
	return len(value) < 8
}

func IsTitleValid(value string) bool {
	return len(value) > 10 && len(value) < 150
}

func IsDescriptionValid(value string) bool {
	return len(value) > 10 && len(value) < 500
}
