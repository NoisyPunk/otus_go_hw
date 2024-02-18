package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:35|regexp:\\d+"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$|len:100"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Response{
				Code: 201,
				Body: "test",
			}, // Place your code here.
		},
		{
			in: User{
				ID:     "ranmcurymipmrtomhyacepvnpdwaslhsrwws",
				Name:   "test",
				Age:    12,
				Email:  "ololool@ol.ru",
				Role:   "admen",
				Phones: []string{"92000000000", "920000000014", "920000000023"},
				meta:   nil,
			},
			expectedErr: ErrInterfaceNotStruct, // Place your code here.
		},

		// ...
		// Place your code here.
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			//require.Len(t, err, 0)
			var testErr ValidationErrors
			if errors.As(err, &testErr) {
				fmt.Println("wow")
			} else {
				fmt.Println("Haha")
			}

			// Place your code here.
			_ = tt
		})
	}
}
