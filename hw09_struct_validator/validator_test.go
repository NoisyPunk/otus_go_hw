package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Codes struct {
		Code []int `validate:"in:350,12,22"`
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

func TestValidateErrors(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Response{
				Code: 201,
				Body: "test",
			},
			expectedErr: ErrNotInRange,
		},
		{
			in: User{
				Phones: []string{"92000000000", "92000000001", "920000000023"},
			},
			expectedErr: ErrLenIsNotEqual,
		},
		{
			in: User{
				ID: "ranmcurymipmrtomhyacepvnpdwaslhsrwwsasdas",
			},
			expectedErr: ErrLenIsNotEqual,
		},
		{
			in: User{
				Age: 11,
			},
			expectedErr: ErrValueIsLessThenMin,
		},
		{
			in: User{
				Age: 102,
			},
			expectedErr: ErrValueIsMoreThenMax,
		},
		{
			in: User{
				Email: "test",
			},
			expectedErr: ErrRegExpIsNotEqual,
		},
		{
			in: Codes{
				Code: []int{1, 22, 12},
			},
			expectedErr: ErrNotInRange,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			var testErr ValidationErrors
			require.ErrorAs(t, err, &testErr)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestRegularErrors(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "test",
			expectedErr: ErrInterfaceNotStruct,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestSuccessFlow(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: App{
				Version: "adsad",
			},
		},
		{
			in: &App{
				Version: "adsad",
			},
		},
		{
			in: Token{
				Header: []byte{},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.NoError(t, err)
		})
	}
}
