NODE_VERSION=$(shell node -v)

.PHONY: help
help:
	@echo "Envs:"
	@echo "    node version=$(NODE_VERSION)"
	@echo ""
	@echo "Commands:"
	@echo "    make install           # install dependencies"
	@echo "    make api/install       # install api dependencies"
	@echo ""
	@echo "    make test/all          # run all tests"
	@echo "    make test/e2e          # run all e2e tests"
	@echo "    make test/e2e/show     # run all e2e tests in a visible browser"
	@echo "    make test/api          # run all api tests"
	@echo "    make test/contracts    # run all contract tests"
	@echo ""
	@echo "    make tree              # list directories"
	@echo "    make log               # show git log"
	@echo ""

.PHONY: install
install:
	@cd tests/e2e && make install
	@cd tests/api && make install

.PHONY: api/install
api/install:
	@cd tests/api && make install

.PHONY: test/all
test/all: test/e2e test/api

.PHONY: test/e2e
test/e2e:
	@cd tests/e2e && make test/all

.PHONY: test/e2e/show
test/e2e/show:
	@cd tests/e2e && make test/all/show

.PHONY: test/api
test/api:
	@cd tests/api && make test/all

.PHONY: tree
tree:
	tree -I node_modules

.PHONY: log
log: ## show git log
	@git log --graph --oneline --decorate

