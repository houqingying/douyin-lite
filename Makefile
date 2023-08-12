# Obtain an absolute path to the directory of the Makefile.
# Assume the Makefile is in the root of the repository.
REPODIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
GOLANGCILINT_VERSION ?= v1.52.2
GLOBAL_GOLANGCILINT := $(shell which golangci-lint)
GOBIN := $(shell go env GOPATH)/bin
GOBIN_GOLANGCILINT := $(shell which $(GOBIN)/golangci-lint)

##@ Help
.PHONY: help
help: ## Display this help screen
	@awk -v target="$(filter-out $@,$(MAKECMDGOALS))" 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { if (!target || index($$1, target) == 1) printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { getline nextLine; if (!target || index(nextLine, target) == 1) printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Lint
.PHONY: lint
.PHONY: install-golint

# verify will verify the code.
verify: verify-mod

# verify-mod will check if go.mod has beed tidied.
verify-mod:
	hack/make-rules/verify_mod.sh

install-golint: ## check golint if not exist install golint tools
ifeq ($(shell $(GLOBAL_GOLANGCILINT) version --format short), $(GOLANGCILINT_VERSION))
GOLINT_BIN=$(GLOBAL_GOLANGCILINT)
else ifeq ($(shell $(GOBIN_GOLANGCILINT) version --format short), $(GOLANGCILINT_VERSION))
GOLINT_BIN=$(GOBIN_GOLANGCILINT)
else
	@{ \
    set -e ;\
    echo 'installing golangci-lint-$(GOLANGCILINT_VERSION)' ;\
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCILINT_VERSION) ;\
    echo 'Successfully installed' ;\
    }
GOLINT_BIN=$(GOBIN)/golangci-lint
endif

lint: install-golint ## Run go lint against code.
	GOOS=linux $(GOLINT_BIN) run -v


.PHONY: test
test:
	go test -v -short -coverprofile cover.out



# Add this line at the end of the Makefile
%: ;