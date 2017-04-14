package models

import (
	"errors"

	"github.com/beego"
	"github.com/beego/validation"
)

/*
	Get form error
*/
func IsValid(model interface{}) (err error) {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		for _, err := range valid.Errors {
			beego.Warning(err.Key, ":", err.Message)
			return errors.New(err.Message)
			// return errors.New(fmt.Sprintf("%s: %s", err.Key, err.Message))
		}
	}
	return nil
}
