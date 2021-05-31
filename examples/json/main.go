package main

import (
	"fmt"

	"github.com/srce/validator"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (u User) Valid() (bool, validator.Errors) {
	return validator.ByJSON(u, validator.Validators{
		"name":    validator.StringRequire(u.Name, validator.Shorter(2)),
		"email":   validator.StringRequire(u.Email, validator.Email()),
		"message": validator.StringOption(u.Name, validator.Shorter(10)),
	})
}

func main() {
	user := User{}

	ok, errs := user.Valid()
	fmt.Printf("validation: %v %v", ok, errs)
}
