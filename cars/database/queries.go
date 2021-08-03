package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver"
	"skeleton.service/database"
	"skeleton.service/env"
)

var dbQueries = DbQueries{
	FindCarByID: findCarByID,
}

func findCarByID(ctx context.Context, id string) (*Car, error) {
	bindVars := map[string]interface{}{
		"car_key": id,
	}

	query := fmt.Sprintf("FOR c IN %s FILTER c._key == @car_key RETURN c", env.DBCarsCollectionName)

	cursor, err := database.DB.Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var doc Car
	meta, err := cursor.ReadDocument(ctx, &doc)
	if driver.IsNoMoreDocuments(err) {
		return nil, errors.New(fmt.Sprintf("Car Not Found to provide ID. %s", id))
	}
	if err != nil {
		return nil, err
	}
	doc.ID = meta.Key

	return &doc, nil
}
