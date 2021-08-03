#!/bin/bash
PORT=8010 \
MOCK_ENABLED=false \
ARANGO_HOST=http://localhost:8529 \
ARANGO_DB_NAME=skeleton \
ARANGO_DB_USER_NAME=root \
ARANGO_DB_USER_PASSWORD=rootpassword \
DB_CARS_COLLECTION_NAME=cars \
go run app.go server-run
