build:
	docker-compose bild todo-app

run:
	docker-compose up todo-app

test:
	go test -v ./...

migrate:
	migrate -path ./migrations -database 'postgres://postgres:Qwer1111@0.0.0.0:5342/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go