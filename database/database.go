package database

import (
	"context"
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"skeleton.service/env"
)

var (
	DB driver.Database
)

func Init() {
	ctx := context.Background()

	if env.MockEnabled {
		log.Println("Mock database is enabled")
		return
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{env.ArangoHost},
		TLSConfig: &tls.Config{},
	})
	if err != nil {
		panic(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(env.ArangoDBUserName, env.ArangoDBUserPassword),
	})
	if err != nil {
		panic(err)
	}

	database, err := client.Database(context.Background(), env.ArangoDBName)
	if err != nil {
		database, err = client.CreateDatabase(context.Background(), env.ArangoDBName, &driver.CreateDatabaseOptions{})
		if err != nil {
			panic(err)
		}
	}

	DB = database

	if err = prepareCollection(ctx, env.DBCarsCollectionName); err != nil {
		panic(err)
	}
}

func prepareCollection(ctx context.Context, collectionName string) error {
	found, err := DB.CollectionExists(ctx, collectionName)
	if err != nil {
		return err
	}

	if found {
		return nil
	}

	options := &driver.CreateCollectionOptions{}
	if _, err = DB.CreateCollection(context.Background(), collectionName, options); err != nil {
		return err
	}
	return nil
}
