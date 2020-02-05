GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=rank-srv
    
all: build run
deps:
	go get \
		github.com/lucazulian/cryptocomparego \
		github.com/lib/pq 
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
run:
	./$(BINARY_NAME)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)