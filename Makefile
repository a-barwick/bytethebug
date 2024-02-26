BUILD_DIR = build/
BINARY_NAME = bytethebug

TEST_DATA_FLAG = --testdata

dev:
	@templ generate
	@go run . $(TEST_DATA_FLAG)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

docker-run-release:
    docker run --rm -tdp 8080:8080 bytethebug:latest

.DEFAULT_GOAL := dev
.PHONY: all dev build
