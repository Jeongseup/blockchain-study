PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NameService \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=named \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=namecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: nsd nscli

named:
		go build -mod=readonly $(BUILD_FLAGS) ./cmd/named/

namecli:
		go build -mod=readonly $(BUILD_FLAGS) ./cmd/namecli/

install: go.sum
		@echo "--> Installing nsd & nscli"
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/named
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/namecli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

clean:
	rm -f named namecli

.PHONY: clean
