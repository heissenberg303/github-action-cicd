run:
	go run cmd/main.go

test:
	go test ./...

build_docker:
	docker build -t covid .

run_docker:
	docker run -it --rm covid