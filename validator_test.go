package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByJSON(t *testing.T) {
	var (
		success = func() error {
			return nil
		}

		errTest = errors.New("test")

		fail = func() error {
			return errTest
		}

		testStr = struct {
			One int `json:"one"`
			Two int `json:"two"`
		}{}
	)

	type exp struct {
		valid bool
		errs  Errors
	}

	cases := []struct {
		name string
		arg  Validators
		exp  exp
	}{
		{
			name: "empty",
			arg:  Validators{},
			exp:  exp{true, Errors{}},
		},
		{
			name: "non error",
			arg: Validators{
				"one": success,
				"two": success,
			},
			exp: exp{true, Errors{}},
		},
		{
			name: "one missed",
			arg: Validators{
				"one": success,
			},
			exp: exp{true, Errors{}},
		},
		{
			name: "one error",
			arg: Validators{
				"one": success,
				"two": fail,
			},
			exp: exp{false, Errors{
				"two": errTest,
			}},
		},
		{
			name: "two errors",
			arg: Validators{
				"one": fail,
				"two": fail,
			},
			exp: exp{false, Errors{
				"one": errTest,
				"two": errTest,
			}},
		},
		{
			name: "additional",
			arg: Validators{
				"one":   success,
				"two":   success,
				"three": fail,
			},
			exp: exp{true, Errors{}},
		},
	}

	for _, tc := range cases {
		ok, errs := ByJSON(testStr, tc.arg)
		assert.Equal(t, tc.exp.valid, ok, tc.name)
		assert.Equal(t, tc.exp.errs, errs, tc.name)
	}
}

func TestErrorsEmpty(t *testing.T) {
	errs := Errors{}
	assert.True(t, errs.Empty())

	var errTest = errors.New("test")
	errs["test"] = errTest
	assert.False(t, errs.Empty())
}
