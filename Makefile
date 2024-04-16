GOCMD=go 
GOBUILD=$(GOCMD) build 
GOINSTALL=$(GOCMD) install

all: build install 

build:
	$(GOBUILD) -o bin/ansiblecli main.go 

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 $(GOBUILD) -o bin/linuxarm64-ansiblecli main.go

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linuxamd64-ansiblecli main.go

build-linux-386:
	env GOOS=linux GOARCH=386 $(GOBUILD) -o bin/linux386-ansiblecli main.go

build-linux-arm:
	env GOOS=linux GOARCH=arm $(GOBUILD) -o bin/linuxarm-ansiblecli main.go

build-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 $(GOBUILD) -o bin/darwinarm64-ansiblecli main.go

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/darwinamd64-ansiblecli main.go

build-windows-arm64:
	env GOOS=windows GOARCH=arm64 $(GOBUILD) -o bin/windowsarm64-ansiblecli main.go

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/windowsamd64-ansiblecli main.go

build-windows-386:
	env GOOS=windows GOARCH=386 $(GOBUILD) -o bin/windows386-ansiblecli main.go

build-windows-arm:
	env GOOS=windows GOARCH=arm $(GOBUILD) -o bin/windowsarm-ansiblecli main.go

install:
	$(GOINSTALL)