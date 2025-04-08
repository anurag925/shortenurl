.PHONY: run
run:
	go run ./cmd/api/main.go

.PHONY: swag
swag:
	swag fmt -g ./cmd/api/main.go
	swag init -g ./cmd/api/main.go -o ./docs	