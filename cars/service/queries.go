package service

import (
	"context"
	"skeleton.service/cars/database"
)

func getCarByID(ctx context.Context, id string) (*database.Car, error) {
	return database.GetDbQueriesManagement().FindCarByID(ctx, id)
}
