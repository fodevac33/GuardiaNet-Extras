# Variables
APP_NAME := keylogger_virus
BIN_DIR := bin
SRC_DIR := .

# Default target
all: build

# Build the binary with a .pdf extension
build:
	@echo "Building $(APP_NAME).pdf..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME).pdf $(SRC_DIR)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BIN_DIR)/$(APP_NAME).pdf

# Run the program
run: build
	@echo "Running $(APP_NAME).pdf..."
	@./$(BIN_DIR)/$(APP_NAME).pdf

.PHONY: all build clean run
