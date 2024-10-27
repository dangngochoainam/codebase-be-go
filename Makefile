.PHONY:  mod run

mod:
	go mod tidy
	go mod vendor

run:
	go run cmd/server/main.go

