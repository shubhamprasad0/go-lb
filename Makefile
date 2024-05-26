APP_NAME := lb
BIN_DIR_NAME := bin
CMD_DIR_NAME := cmd

ROOT_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
CMD_DIR := $(ROOT_DIR)$(CMD_DIR_NAME)
BIN_DIR := $(ROOT_DIR)$(BIN_DIR_NAME)

all: build
build: $(BIN_DIR)/$(APP_NAME)

$(BIN_DIR)/$(APP_NAME):
	@echo "Building the application..."
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/$(APP_NAME)
	@echo "Built $(BIN_DIR)/$(APP_NAME)"

clean:
	@rm -rf $(BIN_DIR)