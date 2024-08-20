GOAPP=totp
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test


build: clean
	cd tests && $(GOBUILD) -o ../bin/$(GOAPP)

clean:
	cd tests && $(GOCLEAN)
	rm -rf ./bin

.PHONY: test
test:
	cd tests && $(GOTEST)