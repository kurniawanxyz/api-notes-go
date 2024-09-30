package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidation(c *gin.Context, model any) map[string]string {
	validate := validator.New()
	if err := validate.Struct(model); err != nil {
		
		errors := make(map[string]string)
        for _, err := range err.(validator.ValidationErrors) {
            errors[err.Field()] = fmt.Sprintf("Error: %s %s",err.Field(),  err.Tag())
        }
		return errors
	}
	return nil
}