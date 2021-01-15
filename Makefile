PROJECT_NAME    := "github.com/bo-er/todo-simper"
PKG             := "$(PROJECT_NAME)"
PKG_LIST        := $(shell go list ${PKG}/... | grep -v /vendor/)
NOW             = $(shell date -u '+%Y%m%d%I%M%S')
APP             = admin
RELEASE_VERSION = v1.0.0
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

start:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP}/main.go todosimpler -c ./configs/config.toml