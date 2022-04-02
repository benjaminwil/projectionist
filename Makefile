DEPENDENCIES=\
  go \
  shelltest
K := $(foreach exec, $(DEPENDENCIES), \
				$(if $(shell which $(exec)),, $(error "No $(exec) in PATH")))

.PHONY: test-build test-go test-shelltest test

build:
	go build -o build/projectionist src/main.go

run:
	go run src/main.go

test-build:
	@printf "\n--- Compile Projectionist. ---\n"
	@rm -rf test/build/* && mkdir -p test/build
	go build -o test/build/projectionist src/main.go

test-go:
	@printf "\n--- Run Go unit tests. ---\n"
	go test ./...

test-shelltest:
	@printf "\n--- Run end-to-end commandline tests. ---\n"
	shelltest -ac src.shelltest

test: test-build test-shelltest test-go

all: build
