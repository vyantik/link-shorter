package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid(body any) error {
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		return err
	}

	return nil
}
