package service

import (
	"context"
	"skeleton.service/cars/database"
	"skeleton.service/constants"
	"time"
)

func postCar(ctx context.Context, request PostRequest) (*database.Car, error) {
	car := &database.Car{
		Brand:   request.Brand,
		Model:   request.Model,
		Created: time.Now().Format(constants.TimeLayoutISO),
	}

	return database.GetDbCommandsManagement().PostCar(ctx, car)
}
