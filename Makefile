NODE_VERSION=$(shell node -v)

.PHONY: help
help:
	@echo "Envs:"
	@echo "    node version=$(NODE_VERSION)"
	@echo ""
	@echo "Commands:"
	@echo "    make install           # install dependencies"
	@echo ""
	@echo "    make test/all          # run all tests"
	@echo "    make test/e2e          # run all api tests"
	@echo "    make test/api          # run all api tests"
	@echo "    make test/contracts    # run all contract tests"
	@echo ""
	@echo "    make tree              # list directories"
	@echo ""

.PHONY: install
install:
	@cd tests/e2e && make install
	@cd tests/api && make install

.PHONY: test/all
test/all: test/e2e test/api

.PHONY: test/e2e
test/e2e:
	@cd tests/e2e && make test/all

.PHONY: test/api
test/api:
	@cd tests/api && make test/all

.PHONY: tree
tree:
	tree -I node_modules

