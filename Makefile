build:
	@go build -o bin/jwt

run: build
	@./bin/jwt

test:
	@go test -v ./..@