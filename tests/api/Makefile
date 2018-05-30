GO_VERSION=$(shell go version)

.PHONY: help
help:
	@echo "Envs:"
	@echo "    $(GO_VERSION)"
	@echo ""
	@echo "Commands:"
	@echo "    make install               # install dependencies"
	@echo ""
	@echo "    make test/all              # run all tests"
	@echo "    make test/all/notruncate   # run all tests without truncating log output"
	@echo "    make test/all/verbose      # run all tests with verbose output"
	@echo ""
	@echo "    make tree                  # list directories"
	@echo "    make troubleshoot          # show how to troubleshoot"
	@echo ""

.PHONY: install
install:
	go get golang.org/x/vgo

##
## TEST
##
.PHONY: test/all
test/all:
	go test ./... -truncate 500

.PHONY: test/all/notruncate
test/all/notruncate:
	go test ./...

.PHONY: test/all/verbose
test/all/verbose:
	go test ./... -v -truncate 500


##
## TEST
##

.PHONY: tree
tree:
	tree -I node_modules


.PHONY: troubleshoot
troubleshoot:
	@echo "If you hit Githubs rate limit running vgo"
	@echo "Generate token @ https://github.com/settings/tokens - The token only needs 'public_repo' scope"
	@echo "echo 'machine api.github.com login <generated token> >> .netrc'"
