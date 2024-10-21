package main

import (
	"projetoEmail/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{Contacts: []campaign.Contact{{Email: ""}}}

	validate := validator.New()

	err := validate.Struct(campaign)

	if err == nil {
		println("Valid")
	} else {
		v := err.(validator.ValidationErrors)
		for _, v := range v {
			println(v.Error())
		}
	}
}
