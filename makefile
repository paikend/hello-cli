# Makefile for hello-cli

# Variables
APP_NAME := hello-cli
GO_BUILD := go build
GO_RUN := go run

# Build Targets
.PHONY: build
build:
	$(GO_BUILD) -o $(APP_NAME)
.PHONY: run
run:
	$(GO_RUN) main.go

.PHONY: clean
clean:
	rm -f $(APP_NAME)

.PHONY: all
all: clean build run
