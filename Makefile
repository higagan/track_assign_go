.PHONY: build run test

build:
	docker build -t analytic-service .

run:
	docker run -p 8080:8080 analytic-service

test:
	go test ./...
