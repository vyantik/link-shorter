package req

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func IsValid(body any) error {
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		log.Printf("[Req] - [Validate] - [ERROR] : %s", err)
		return err
	}

	return nil
}
