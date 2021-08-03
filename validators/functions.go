package validators

import (
	"log"
)

var Validator StructValidator

func Init() {
	log.Println("Init validators")

	Validator = StructValidator{}
	Validator.Initialize()
}
