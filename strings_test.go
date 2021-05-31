package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	cases := []struct {
		name    string
		length  int
		message string
		exp     error
	}{
		{name: "zero", length: 0, message: "", exp: nil},
		{name: "empty", length: 2, message: "", exp: ErrWrongLength},
		{name: "less", length: 2, message: "1", exp: ErrWrongLength},
		{name: "equal", length: 2, message: "12", exp: nil},
		{name: "more", length: 2, message: "123", exp: ErrWrongLength},
		{name: "much more", length: 2, message: "1234", exp: ErrWrongLength},
	}

	for _, tc := range cases {
		err := Length(tc.length)(tc.message)
		assert.ErrorIs(t, err, tc.exp, tc.name)
	}
}

func TestShorter(t *testing.T) {
	cases := []struct {
		name    string
		length  int
		message string
		exp     error
	}{
		{name: "zero", length: 0, message: "", exp: nil},
		{name: "much shorter", length: 2, message: "", exp: nil},
		{name: "shorter", length: 2, message: "1", exp: nil},
		{name: "equal", length: 2, message: "12", exp: nil},
		{name: "longer", length: 2, message: "123", exp: ErrTooLong},
		{name: "much longer", length: 2, message: "1234", exp: ErrTooLong},
	}

	for _, tc := range cases {
		err := Shorter(tc.length)(tc.message)
		assert.ErrorIs(t, err, tc.exp, tc.name)
	}
}

func TestLonger(t *testing.T) {
	cases := []struct {
		name    string
		length  int
		message string
		exp     error
	}{
		{name: "zero", length: 0, message: "", exp: nil},
		{name: "much shorter", length: 2, message: "", exp: ErrTooShort},
		{name: "shorter", length: 2, message: "1", exp: ErrTooShort},
		{name: "equal", length: 2, message: "12", exp: nil},
		{name: "longer", length: 2, message: "123", exp: nil},
		{name: "much longer", length: 2, message: "1234", exp: nil},
	}

	for _, tc := range cases {
		err := Longer(tc.length)(tc.message)
		assert.ErrorIs(t, err, tc.exp, tc.name)
	}
}

func TestEmail(t *testing.T) {
	cases := []struct {
		name  string
		email string
		exp   error
	}{
		{name: "empty", email: "", exp: ErrEmailInvalid},
		{name: "one word", email: "local", exp: ErrEmailInvalid},
		{name: "without domain", email: "local@local", exp: nil},
		{name: "with domain", email: "local@local.com", exp: nil},
		{name: "numbers", email: "111@222.com", exp: nil},
	}

	for _, tc := range cases {
		err := Email()(tc.email)
		assert.ErrorIs(t, err, tc.exp, tc.name)
	}
}
