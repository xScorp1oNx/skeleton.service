package database

import (
	"context"
	"skeleton.service/env"
)

type DbQueries struct {
	FindCarByID func(ctx context.Context, id string) (*Car, error)
}

type DbCommands struct {
	PostCar func(ctx context.Context, car *Car) (*Car, error)
}

func GetDbQueriesManagement() DbQueries {
	if env.MockEnabled {
		return dbQueriesMock
	}
	return dbQueries
}

func GetDbCommandsManagement() DbCommands {
	if env.MockEnabled {
		return dbCommandsMock
	}
	return dbCommands
}
