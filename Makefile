# It is primarily used to simplify long, complex commands,
# ensuring that developers or CI/CD pipelines execute the exact same command without the risk of typos.

# Vars
BUILD_DIR = bin

.PHONY: tidy fmt build-exe

tidy:
	go mod tidy

fmt:
	go fmt

build-exe:
	@echo "Building a Windows executable"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/app.exe main.go
