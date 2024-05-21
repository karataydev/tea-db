NAME=tea-db

build:
	@go build -o bin/$(NAME) ./cmd/$(NAME)

run: build
	@./bin/$(NAME)
test:
	@go test -v ./...

compile:
	GOOS=freebsd GOARCH=386 go build -o bin/$(NAME)-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/$(NAME)-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/$(NAME)-windows-386 main.go

build-all: windows linux darwin
	@echo version: $(VERSION)

VERSION=$(shell git describe --tags)
WINDOWS=$(NAME)_windows_amd64_$(VERSION).exe
LINUX=$(NAME)_linux_amd64_$(VERSION)
DARWIN=$(NAME)_darwin_amd64_$(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o bin/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/db2db/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o bin/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/db2db/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/db2db/main.go

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
