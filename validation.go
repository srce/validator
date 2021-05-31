package validation

import (
	"errors"
	"reflect"
)

var ErrRequired = errors.New("field required")

type Errors map[string]error

func (e Errors) Empty() bool {
	return len(e) == 0
}

type (
	Validator interface {
		Valid() (bool, Errors)
	}

	Validation func() error

	Validators map[string]Validation
)

func ByJSON(s interface{}, validators Validators) (bool, Errors) {
	const tagName = "json"

	var (
		ok   = true
		errs = map[string]error{}
	)

	if s == nil || len(validators) == 0 {
		return ok, errs
	}

	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		panic("value should be struct")
	}

	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Tag.Get(tagName)
		if valid, exists := validators[name]; exists {
			if err := valid(); err != nil {
				errs[name] = err
				ok = false
			}
		}
	}

	return ok, errs
}
