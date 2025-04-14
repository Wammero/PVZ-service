.PHONY: docker-start-service-debug docker-clean-data docker-start-test-debug

docker-start-service-debug:
	docker-compose build
	docker-compose up -d

docker-clean-data:
	docker-compose down -v

docker-start-test-debug:
	docker-compose -f docker-compose.test.yml build
	docker-compose -f docker-compose.test.yml up -d
	docker-compose -f docker-compose.test.yml exec -T PVZ-service-test go test ./...
	docker-compose -f docker-compose.test.yml down -v

lint:
	golangci-lint run