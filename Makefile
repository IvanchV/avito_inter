include .env
export

migrate-up:
	migrate -path migrations -database postgres://Vasily:5525@localhost:5432/avito?sslmode=disable up

migrate-down:
	 migrate -path migrations -database postgres://Vasily:5525@localhost:5432/avito?sslmode=disable down

compose-up:
	sudo docker-compose up --build -d && docker-compose logs -f

compose-down:
	sudo docker-compose down --remove-orphans
