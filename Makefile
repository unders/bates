NODE_VERSION=$(shell node -v)

.PHONY: help
help:
	@echo "Envs:"
	@echo "    node version=$(NODE_VERSION)"
	@echo ""
	@echo "Commands:"
	@echo "    make install           # install dependencies"
	@echo "    make target            # some information"
	@echo ""

.PHONY: install
	@echo "    installing..."

