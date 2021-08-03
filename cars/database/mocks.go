package database

import (
	"context"
	"errors"
	"fmt"
)

var dbQueriesMock = DbQueries{
	FindCarByID: findCarByIDMock,
}

var dbCommandsMock = DbCommands{
	PostCar: postCarMock,
}

var (
	carsByID    = make(map[string]*Car, 5)
	carsByBrand = make(map[string][]*Car, 5)
	carsByModel = make(map[string][]*Car, 5)
)

func AppendCarToMock(car *Car) {
	carsByID[car.ID] = car
	carsByBrand[car.Brand] = append(carsByBrand[car.Brand], car)
	carsByModel[car.Model] = append(carsByModel[car.Model], car)
}

func findCarByIDMock(ctx context.Context, id string) (*Car, error) {
	for i, _ := range carsByID {
		if i == id {
			return carsByID[id], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Car Not Found to provide ID. %s", id))
}

func postCarMock(ctx context.Context, car *Car) (*Car, error) {
	AppendCarToMock(car)
	return car, nil
}
