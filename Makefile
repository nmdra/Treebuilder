APP_NAME = treebuilder

BUILD_DIR = bin

.PHONY: all build test clean run

all: clean build test

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Build complete! Executable is in $(BUILD_DIR)/$(APP_NAME)"

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@go clean

# A helper to quickly run the CLI (e.g., make run ARGS="-d my_structure.txt")
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME) $(ARGS)
