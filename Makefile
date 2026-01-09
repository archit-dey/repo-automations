.PHONY: *

# Build variables
PROJECT_NAME=demo-app
BUILD_DIR=build

build:
	mkdir -p $(BUILD_DIR)
	cd $(PROJECT_NAME) && go build -o ../$(BUILD_DIR)/$(PROJECT_NAME) .

test:
	cd $(PROJECT_NAME) && go test -v ./...

run: build
	./$(BUILD_DIR)/$(PROJECT_NAME)

fmt:
	cd $(PROJECT_NAME) && go fmt ./...

vet:
	cd $(PROJECT_NAME) && go vet ./...

clean:
	rm -rf $(BUILD_DIR)
	cd $(PROJECT_NAME) && go clean


pr: build test clean