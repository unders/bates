NODE_VERSION=$(shell node -v)

.PHONY: help
help:
	@echo "Envs:"
	@echo "    node version=$(NODE_VERSION)"
	@echo ""
	@echo "Commands:"
	@echo "    make install           # install dependencies"
	@echo "    make test              # run tests"
	@echo "    make tree              # list directories"
	@echo ""

.PHONY: install
install:
	@npm install -g typescript
	@npm install


.PHONY: test
test:
	node_modules/.bin/jest

.PHONY: tree
tree:
	tree -I node_modules
