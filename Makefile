GIT_SHA=`git rev-parse --short HEAD || echo`
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=intTracker
all: build test
build:
				$(GOBUILD) -o bin/$(BINARY_NAME) -v
build-linux:
				CGO_ENABLED=0 GOOS=linux GOOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME) -v
test:
				$(GOTEST) -v ./...
clean:
				$(GOCLEAN)
				rm -rf bin/$(BINARY_NAME)
run: build
				./bin/$(BINARY_NAME)
docker-build:
				docker build -t dyk0/inttracker:$(GIT_SHA) -f Dockerfile $(CURDIR)
docker-push:
				docker tag dyk0/inttracker:$(GIT_SHA) dyk0/inttracker:latest
				docker push dyk0/inttracker:$(GIT_SHA)
				docker push dyk0/inttracker:latest
