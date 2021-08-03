package database

import (
	"context"
	"skeleton.service/database"
	"skeleton.service/env"
)

var dbCommands = DbCommands{
	PostCar: postCar,
}

func postCar(ctx context.Context, car *Car) (*Car, error) {
	col, err := database.DB.Collection(context.Background(), env.DBCarsCollectionName)
	if err != nil {
		return nil, err
	}

	meta, err := col.CreateDocument(context.Background(), car)
	if err != nil {
		return nil, err
	}

	car.ID = meta.ID.Key()

	return car, nil
}
