package main

import (
	"log"
	"os"
	"skeleton.service/cli"
)

// @title Cars service Swagger API
// @version 1.0
// @description Swagger API for cars service.
// @contact.email dponomarov25@gmail.com
// @BasePath /api/
func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
