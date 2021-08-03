package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	PublicKey  []byte
	PrivateKey []byte

	MockEnabled bool

	Port string

	ArangoHost           string
	ArangoDBName         string
	ArangoDBUserName     string
	ArangoDBUserPassword string

	DBCarsCollectionName string
)

func Initialize() {
	PublicKey = getRSAKeyEnv("PUBLIC_KEY")
	PrivateKey = getRSAKeyEnv("PRIVATE_KEY")

	if val, ok := os.LookupEnv("PORT"); ok {
		Port = val
	}

	if val, ok := os.LookupEnv("MOCK_ENABLED"); ok {
		if v, err := strconv.ParseBool(val); err != nil {
			log.Println("Cannot parse value of MOCK_ENABLED. Using false as default.")
			MockEnabled = false
		} else {
			MockEnabled = v
		}
	} else {
		MockEnabled = false
	}

	if val, ok := os.LookupEnv("ARANGO_HOST"); ok {
		ArangoHost = val
	}
	if val, ok := os.LookupEnv("ARANGO_DB_NAME"); ok {
		ArangoDBName = val
	}
	if val, ok := os.LookupEnv("ARANGO_DB_USER_NAME"); ok {
		ArangoDBUserName = val
	}
	if val, ok := os.LookupEnv("ARANGO_DB_USER_PASSWORD"); ok {
		ArangoDBUserPassword = val
	}

	if val, ok := os.LookupEnv("DB_CARS_COLLECTION_NAME"); ok {
		DBCarsCollectionName = val
	}
}

func getRSAKeyEnv(key string) []byte {
	return []byte(strings.Replace(os.Getenv(key), `\n`, "\n", -1))
}
