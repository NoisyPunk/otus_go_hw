package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInterfaceNotStruct TechnicalErrors = errors.New("given interface is not a struct")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

type TechnicalErrors error

func (v ValidationErrors) Error() string {
	test := ValidationError{
		Field: "test",
		Err:   nil,
	}
	return test.Field
}

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

		switch value.(type) {
		case int:
			validationErr, err := validateInt(value.(int), tagValue, name)
			if err != nil {
				return err
			}
			if validationErr != nil {
				validationErrors = append(validationErrors, *validationErr)
			}

		case []int:
			for _, val := range value.([]string) {
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				validationErr, err := validateInt(intVal, tagValue, name)
				if err != nil {
					return err
				}
				if validationErr != nil {
					validationErrors = append(validationErrors, *validationErr)
				}
			}
		case string:
			validationErr, err := validateString(value.(string), tagValue, name)
			if err != nil {
				return err
			}
			if validationErr != nil {
				validationErrors = append(validationErrors, *validationErr)
			}

		case []string:
			for _, val := range value.([]string) {
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

func validateString(string string, tagValues string, name string) (*ValidationError, error) {
	rules := strings.Split(tagValues, "|")

	for _, rule := range rules {
		switch {
		case strings.HasPrefix(rule, "len:"):
			strLen, err := strconv.Atoi(strings.TrimPrefix(rule, "len:"))
			if err != nil {
				return nil, err
			}
			if len(string) != strLen {
				return &ValidationError{
					Field: name,
					Err:   fmt.Errorf("can't validate field '%s' len of value is not equal with required == %d", name, strLen),
				}, nil
			}
		case strings.HasPrefix(rule, "regexp:"):
			regExp := strings.TrimPrefix(rule, "regexp:")
			valid, err := regexp.MatchString(regExp, string)
			if err != nil {
				return nil, err
			}
			if !valid {
				return &ValidationError{
					Field: name,
					Err:   fmt.Errorf("can't validate field '%s' is not require with reqExp", name),
				}, nil
			}

		case strings.HasPrefix(rule, "in:"):
			var counter int
			in := strings.TrimPrefix(rule, "in:")
			values := strings.Split(in, ",")
			for _, value := range values {
				if string == value {
					counter++
				}
			}
			if counter == 0 {
				return &ValidationError{
					Field: name,
					Err:   fmt.Errorf("can't validate field '%s' is not in range of expected values", name),
				}, nil

			}

		}

	}
	return nil, nil
}

func validateInt(digit int, tagValues string, name string) (*ValidationError, error) {
	rules := strings.Split(tagValues, "|")

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
					Err:   fmt.Errorf("can't validate field '%s' value is less than minimum == %d", name, minimum),
				}, nil
			}
		case strings.HasPrefix(rule, "max:"):
			maximum, err := strconv.Atoi(strings.TrimPrefix(rule, "mix:"))
			if err != nil {
				return nil, err
			}
			if digit > maximum {
				return &ValidationError{
					Field: name,
					Err:   fmt.Errorf("can't validate field '%s' value is more than maximum == %d", name, maximum),
				}, nil
			}
		case strings.HasPrefix(rule, "in:"):
			var counter int
			in := strings.TrimPrefix(rule, "in:")
			values := strings.Split(in, ",")
			for _, value := range values {
				intvalue, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				if digit == intvalue {
					counter++
				}
			}
			if counter == 0 {
				return &ValidationError{
					Field: name,
					Err:   fmt.Errorf("can't validate field '%s' is not in range of expected values", name),
				}, nil

			}

		}
	}
	return nil, nil
}
