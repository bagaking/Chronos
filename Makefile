GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: clean build all

clean:
	$(GOCLEAN)
	rm -rf ./bin

build:
	GOOS=darwin $(GOBUILD) -o bin/osx/chronos main.go
	GOOS=linux $(GOBUILD) -o bin/linux/chronos main.go
	GOOS=windows $(GOBUILD) -o bin/windows/chronos.exe main.go

all: test clean build