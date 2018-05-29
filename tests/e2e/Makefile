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
	@echo "    make test/watch        # watch all tests"
	@echo ""
	@echo "    make tree              # list directories"
	@echo ""

.PHONY: install
install:
	@npm install

##
## TEST
##
.PHONY: test/all
test/all:
	node_modules/.bin/jest

.PHONY: test/watch
test/watch:
	node_modules/.bin/jest --watch

##
## TEST
##
.PHONY: tree
tree:
	tree -I node_modules
