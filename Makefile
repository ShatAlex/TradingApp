.SILENT:

run: 
	docker-compose up --build
migrate:
	docker run -v ./schema:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:ShatAlex@localhost:5434/postgres?sslmode=disable up
