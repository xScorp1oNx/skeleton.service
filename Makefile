BASE_PATH       = ./
MAIN_GO_PATH    = /app.go

swag-init:
	- swag init --parseDependency --parseInternal -d $(BASE_PATH) -g $(MAIN_GO_PATH)

