PROJECT_NAME=$(shell basename "$(PWD)")
IMAGE_TAG=covid-tracking-api
IMAGE_NAME=covid-tracking-api-app

.PHONY: all dep test build deploy clean

all: build

dep: ## Get the dependencies
	@echo "Installing the dependencies...\n"
	@go mod download
	@echo "\nDone"

test: ## Run unit tests
	@echo "Running unit tests...\n"
	@go test -v
	@echo "\nDone"

build: ## Build the docker image
	@echo "Creating the docker image...\n"
	docker build -t=$(IMAGE_TAG) .
	@echo "\nDone"

deploy:
	@echo "Deploy docker image...\n"
	docker run -p 3000:3000 --rm --name $(IMAGE_NAME) -t $(IMAGE_TAG) -d
	@echo "\nDone"

clean: ## Remove previous build
	@echo "Cleaning build files...\n"
	@rm -f $(PROJECT_NAME)/build
	@echo "Done\n"

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'