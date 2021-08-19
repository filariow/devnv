.PHONY: build install clean

GO = go
DEVNV_PATH = $(shell go env GOPATH)/bin/devnv
PM_PATH = ${HOME}/.devnv/scripts/pm.sh

all: build

test:
	@$(GO) test ./...

build:
	@$(GO) build -ldflags="-s -w" -trimpath  -o bin/devnv ./cmd/devnv/main.go

install: build
	@echo "Installing devnv binary into ${DEVNV_PATH}"
	@install ./bin/devnv ${DEVNV_PATH}
	@echo "Installing pm.sh script into ${PM_PATH}"
	@mkdir -p $(shell dirname ${PM_PATH})
	@install ./scripts/pm/pm.sh ${PM_PATH}
	@echo ""
	@echo "Please add the following lines to your .zshrc file:"
	@echo '```'
	@echo "#####################################"
	@echo "## devnv"
	@echo "#####################################"
	@echo "source <(devnv completion oh-my-zsh)"
	@echo "source $(PM_PATH)"
	@echo '```'


uninstall:
	rm -f ${DEVNV_PATH}
	rm -f ${PM_PATH}

clean:
	@rm -rf ./bin

