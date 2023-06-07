ABIGEN_TAG = 1.12.0
# Path to the binary
LOCAL_BIN:=$(CURDIR)/bin
# Путь до бинарника abigen
ABIGEN_BIN:=$(LOCAL_BIN)/abigen

# Install tools
.PHONY: install-tools
install-tools:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info "Downloading abigen v$(GOLANGCI_TAG)")
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/ethereum/go-ethereum/cmd/abigen@v$(ABIGEN_TAG) && \
		go build -ldflags "-X 'main.version=$(ABIGEN_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/abigen github.com/ethereum/go-ethereum/cmd/abigen && \
		rm -rf $$tmp
ABIGEN_BIN:=$(LOCAL_BIN)/abigen
endif

##################### Checks to run golang-ci #####################
# Local bin version check
ifneq ($(wildcard $(ABIGEN_BIN)),)
ABIGEN_BIN_VERSION:=$(shell $(ABIGEN_BIN) --version)
ifneq ($(ABIGEN_BIN_VERSION),)
ABIGEN_BIN_VERSION_SHORT:=$(shell echo "$(ABIGEN_BIN_VERSION)" | sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
ABIGEN_BIN_VERSION_SHORT:=0
endif
ifneq "$(ABIGEN_TAG)" "$(word 1, $(sort $(ABIGEN_TAG) $(ABIGEN_BIN_VERSION_SHORT)))"
ABIGEN_BIN:=
endif
endif

.PHONY: generate
generate:
	PATH="${LOCAL_BIN}:${PATH}" go generate ./...

