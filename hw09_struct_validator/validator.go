package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
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
	// Place your code here.
	validatedStruct := reflect.ValueOf(v)

	if validatedStruct.Type().Kind() != reflect.Struct {
		return ErrInterfaceNotStruct
	}

	//var validationErrors ValidationErrors

	validationErrors := ValidationErrors{}

	for i := 0; i < validatedStruct.Type().NumField(); i++ {
		field := validatedStruct.Type().Field(i).Tag.Get("validate")
		if field == "" {
			continue
		}
		value := validatedStruct.Field(i).Interface()
		name := validatedStruct.Type().Field(i).Name

		switch value.(type) {
		case int:
			err := validateInt(value.(int), field)
			if err != nil {
				return err
			}
		case string:
			validationErr, err := validateString(value.(string), field, name)
			if err != nil {
				return err
			}
			if validationErr != nil {
				validationErrors = append(validationErrors, *validationErr)
			}

		case []string:
			err := validateStringSlice(value.([]string), field)
			if err != nil {
				return err
			}

		case []int:
			err := validateIntSlice(value.([]int), field)
			if err != nil {
				return err
			}

		}
	}
	if len(validationErrors) != 0 {
		return validationErrors
	}
	return nil
}

func validateIntSlice(ints []int, rules string) error {
	fmt.Println(ints)
	return nil
}

func validateStringSlice(strings []string, rules string) error {
	fmt.Println(strings)
	return nil
}

func validateString(string string, rules string, name string) (*ValidationError, error) {
	switch {
	case strings.HasPrefix(rules, "len:"):
		strLen, err := strconv.Atoi(strings.TrimPrefix(rules, "len:"))
		if err != nil {
			return nil, err
		}
		if len(string) != strLen {
			return &ValidationError{
				Field: name,
				Err:   fmt.Errorf("can't validate field '%s' len of value is not equal with required == %d", name, strLen),
			}, nil
		}
	case strings.HasPrefix(rules, "regexp:"):

	case strings.HasPrefix(rules, "in:"):
	}
	return nil, nil
}

func validateInt(int int, rules string) error {
	fmt.Println(int)
	return nil
}
