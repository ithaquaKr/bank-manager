.PHONY: run build test

# ==============================================================================
# Main

run:
	go run main.go

build:
	go build .

test:
	go test -cover ./...

