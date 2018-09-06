package util

import validator "gopkg.in/go-playground/validator.v9"

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		res.Errors[e.Field()] = e.ActualTag()
	}
	return res
}
