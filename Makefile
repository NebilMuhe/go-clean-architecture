sqlc:
	sqlc generate -f ./config/sqlc.yaml

migrate-create:
	migrate create -ext sql -dir internal/constant/query/schemas $(args)

swagger:
	- swag init -g initiator/initiator.go	