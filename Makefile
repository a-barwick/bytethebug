APP_NAME := e-commerce
BUILD_DIR := ./build
BIN_DIR := ./bin

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./*.go

run:
	@echo "Running $(APP_NAME)..."
	@go run ./*.go

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(BIN_DIR)

.DEFAULT_GOAL := run
