DEPENDENCIES=\
  go \
  shelltest
K := $(foreach exec, $(DEPENDENCIES), \
				$(if $(shell which $(exec)),, $(error "No $(exec) in PATH")))

build:
	go build -o build/projectionist src/main.go

run:
	go run src/main.go

test:
	shelltest -ac src.shelltest

all: build
