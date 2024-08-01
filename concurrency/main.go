package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type something int

type address struct {
	Name string `json:"name" validate:"required,min=3,max=30"`
}

func main() {
	addr := address{
		Name: "ma",
	}

	validate := validator.New()
	err := validate.Struct(addr)

	if err != nil {
		// Validation failed
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), "is required")
		}
	} else {
		// Validation passed
		fmt.Println("Validation passed!")
	}
}
