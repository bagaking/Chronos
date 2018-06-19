
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: clean build all

deps:
    $(GOGET)

test:
		$(GOTEST)

clean:
		$(GOCLEAN)
		find bin -type f -delete

build:
		GOOS=darwin $(GOBUILD) -o bin/osx/chronos main.go
		GOOS=linux $(GOBUILD) -o bin/linux/chronos main.go
		GOOS=windows $(GOBUILD) -o bin/windows/chronos.exe main.go

all: deps test clean build