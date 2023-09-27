createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost:5432/todoapi?sslmode=disable" -verbose up

migratedown:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost:5432/todoapi?sslmode=disable" -verbose down

.PHONY: migrate migratedown createmigration