BUILD_DIR = build/
BINARY_NAME = bytethebug

TEST_DATA_FLAG = --testdata

.PHONY: all dev build

dev:
	templ generate
	go run . $(TEST_DATA_FLAG)

build-air:
	templ generate
	go build -o ./tmp/main .

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

setup-postgres:
	docker pull postgres
	docker volume create postgres-data
