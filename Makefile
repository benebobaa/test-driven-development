createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root tdd

dropdb:
	sudo docker exec -it postgres16 dropdb tdd

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/tdd?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/tdd?sslmode=disable" -verbose down 1