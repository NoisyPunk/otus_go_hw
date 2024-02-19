package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInterfaceNotStruct = errors.New("given interface is not a struct")

	ErrLenIsNotEqual      = errors.New("len of value is not equal with required len")
	ErrRegExpIsNotEqual   = errors.New("value is not require with reqExp")
	ErrNotInRange         = errors.New("value is not in range of expected values")
	ErrValueIsMoreThenMax = errors.New("value is more than maximum")
	ErrValueIsLessThenMin = errors.New("value is less than minimum")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errs := make([]string, len(v))
	for _, err := range v {
		errs = append(errs, err.Err.Error())
	}
	stringErr := strings.Join(errs, ",")
	return stringErr
}

//nolint:gocognit
func Validate(v interface{}) error {
	validatedStruct := reflect.ValueOf(v)

	if validatedStruct.Type().Kind() != reflect.Struct {
		return ErrInterfaceNotStruct
	}

	validationErrors := ValidationErrors{}

	for i := 0; i < validatedStruct.Type().NumField(); i++ {
		tagValue := validatedStruct.Type().Field(i).Tag.Get("validate")
		if tagValue == "" {
			continue
		}
		value := validatedStruct.Field(i).Interface()
		name := validatedStruct.Type().Field(i).Name

		switch v := value.(type) {
		case int:
			validationErr, err := validateInt(v, tagValue, name)
			if err != nil {
				return err
			}
			if validationErr != nil {
				validationErrors = append(validationErrors, *validationErr)
			}

		case []int:
			for _, val := range v {
				validationErr, err := validateInt(val, tagValue, name)
				if err != nil {
					return err
				}
				if validationErr != nil {
					validationErrors = append(validationErrors, *validationErr)
				}
			}
		case string:
			validationErr, err := validateString(v, tagValue, name)
			if err != nil {
				return err
			}
			if validationErr != nil {
				validationErrors = append(validationErrors, *validationErr)
			}

		case []string:
			for _, val := range v {
				validationErr, err := validateString(val, tagValue, name)
				if err != nil {
					return err
				}
				if validationErr != nil {
					validationErrors = append(validationErrors, *validationErr)
				}
			}
		}
	}
	if len(validationErrors) != 0 {
		return validationErrors
	}
	return nil
}

func validateString(stringVal string, tagValues string, name string) (*ValidationError, error) {
	if len(stringVal) == 0 {
		return nil, nil
	}
	rules := getRules(tagValues)

	for _, rule := range rules {
		switch {
		case strings.HasPrefix(rule, "len:"):
			strLen, err := strconv.Atoi(strings.TrimPrefix(rule, "len:"))
			if err != nil {
				return nil, err
			}
			if len(stringVal) != strLen {
				return &ValidationError{
					Field: name,
					Err:   ErrLenIsNotEqual,
				}, nil
			}
		case strings.HasPrefix(rule, "regexp:"):
			regExp := strings.TrimPrefix(rule, "regexp:")
			valid, err := regexp.MatchString(regExp, stringVal)
			if err != nil {
				return nil, err
			}
			if !valid {
				return &ValidationError{
					Field: name,
					Err:   ErrRegExpIsNotEqual,
				}, nil
			}

		case strings.HasPrefix(rule, "in:"):
			var match bool
			in := strings.TrimPrefix(rule, "in:")
			values := strings.Split(in, ",")
			for _, value := range values {
				if stringVal == value {
					match = true
					break
				}
			}
			if !match {
				return &ValidationError{
					Field: name,
					Err:   ErrNotInRange,
				}, nil
			}
		}
	}
	return nil, nil
}

func validateInt(digit int, tagValues string, name string) (*ValidationError, error) {
	if digit == 0 {
		return nil, nil
	}
	rules := getRules(tagValues)

	for _, rule := range rules {
		switch {
		case strings.HasPrefix(rule, "min:"):
			minimum, err := strconv.Atoi(strings.TrimPrefix(rule, "min:"))
			if err != nil {
				return nil, err
			}
			if digit < minimum {
				return &ValidationError{
					Field: name,
					Err:   ErrValueIsLessThenMin,
				}, nil
			}
		case strings.HasPrefix(rule, "max:"):
			maximum, err := strconv.Atoi(strings.TrimPrefix(rule, "max:"))
			if err != nil {
				return nil, err
			}
			if digit > maximum {
				return &ValidationError{
					Field: name,
					Err:   ErrValueIsMoreThenMax,
				}, nil
			}
		case strings.HasPrefix(rule, "in:"):
			var match bool
			in := strings.TrimPrefix(rule, "in:")
			values := strings.Split(in, ",")
			for _, value := range values {
				intvalue, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				if digit == intvalue {
					match = true
					break
				}
			}
			if !match {
				return &ValidationError{
					Field: name,
					Err:   ErrNotInRange,
				}, nil
			}
		}
	}
	return nil, nil
}

func getRules(tagValues string) []string {
	rules := strings.Split(tagValues, "|")
	return rules
}
