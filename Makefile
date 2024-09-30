.PHONY: migrate

DATABASE_URL := "mysql://root:2006@tcp(localhost:3306)/notes_go_db?parseTime=true"

migrate-up:
	@echo "Checking if database exists..."
	@mariadb -u root -p2006 -e "CREATE DATABASE IF NOT EXISTS notes_go_db"
	migrate -path ./db/migrations -database ${DATABASE_URL} up

migrate-down:
	migrate -path ./db/migrations -database ${DATABASE_URL} down

migrate-create:
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "Error: name is not set"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir ./db/migrations -seq $$name

migrate-force:
	@read -p "Enter version to force: " version; \
	if [ -z "$$version" ]; then \
		echo "Error: version is not set"; \
		exit 1; \
	fi; \
	migrate -path ./db/migrations -database ${DATABASE_URL} force $$version
