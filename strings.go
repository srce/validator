package validator

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type StringValid func(s string) error

func stringExecuter(s string, valids ...StringValid) error {
	for _, v := range valids {
		if err := v(s); err != nil {
			return err
		}
	}
	return nil
}

func StringRequire(s string, valids ...StringValid) Validation {
	return func() error {
		if utf8.RuneCountInString(s) == 0 {
			return ErrRequired
		}

		return stringExecuter(s, valids...)
	}
}

func StringOption(s string, valids ...StringValid) Validation {
	return func() error {
		if utf8.RuneCountInString(s) == 0 {
			return nil
		}

		return stringExecuter(s, valids...)
	}
}

var ErrWrongLength = errors.New("wrong length")

func Length(length int) StringValid {
	return func(s string) error {
		if utf8.RuneCountInString(s) != length {
			return ErrWrongLength
		}
		return nil
	}
}

var ErrTooShort = errors.New("too short")

func Longer(length int) StringValid {
	return func(s string) error {
		if utf8.RuneCountInString(s) < length {
			return ErrTooShort
		}
		return nil
	}
}

var ErrTooLong = errors.New("too long")

func Shorter(length int) StringValid {
	return func(s string) error {
		if utf8.RuneCountInString(s) > length {
			return ErrTooLong
		}
		return nil
	}
}

var ErrEmailInvalid = errors.New("invalid email format")

func Email() StringValid {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return func(s string) error {
		if !emailRegex.MatchString(s) {
			return ErrEmailInvalid
		}

		return nil
	}
}
