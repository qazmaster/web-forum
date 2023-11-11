# Define variables
BINARY_NAME=bin/forum-auth
SRC_FILES=cmd/main.go cmd/config.go cmd/app.go

# Define targets
all: build

build:
	go build -o $(BINARY_NAME) $(SRC_FILES)

run:
	go run $(SRC_FILES)

test:
	go test -v

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run test clean
