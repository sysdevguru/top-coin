GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=web-srv
    
all: build run
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
run:
	./$(BINARY_NAME)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)