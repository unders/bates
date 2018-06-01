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
	@echo "    make test/all/show     # run all tests with visible broser"
	@echo "    make test/watch        # watch all tests"
	@echo ""
	@echo "    make tree              # list directories"
	@echo ""

.PHONY: install
install:
	npm install

##
## TEST
##
.PHONY: test/all
test/all:
	cp support/jest-puppeteer.config.js .
	node_modules/.bin/jest

.PHONY: test/all/show
test/all/show:
	cp support/show-chrome.jest-puppeteer.config.js jest-puppeteer.config.js
	node_modules/.bin/jest
	cp support/jest-puppeteer.config.js .

.PHONY: test/watch
test/watch:
	cp support/jest-puppeteer.config.js .
	node_modules/.bin/jest --watch

##
## TEST
##
.PHONY: tree
tree:
	tree -I node_modules
