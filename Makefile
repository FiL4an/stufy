include .env
export

service-run: 
	@go run main.go

service-deploy:
	docker compose up -d applications

servie-undeploy:
	docker compose down applications

migrate-up: 
	migrate -path migrations/ -database ${CONN_STRING} up

migrate-down: 
	migrate -path migrations/ -database ${CONN_STRING} down
