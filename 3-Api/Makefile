createdb:
	docker exec -it teste-produtos createdb --username=root --owner=root produtos

dropdb:
	docker exec -it teste-produtos dropdb produtos

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/produtos?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/produtos?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc