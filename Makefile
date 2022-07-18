setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/database.yml.example internal/config/db/database.yml
	@cp internal/config/example/rest.yml.example internal/config/server/rest.yml
	@echo " --- Done Setup and generate configuration --- "

rest:
	@go run main.go
	@#go run cmd/server/restful/main.go

swagger:
	@swag init -g cmd/server/restful/main.go --output pkg/shared/document/swagger

build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/server/restful/main.go