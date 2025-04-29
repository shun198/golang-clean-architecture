prepare:
	docker-compose up -d --build

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

seed:
	docker-compose exec app go run ./cmd/seed/seed.go

migrate:
	docker-compose exec app go run ./internal/infrastructures/migrations/migration.go
