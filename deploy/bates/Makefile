PROG=bates
VERSION=1
DOCKER_USER=unders
WEB=https://hub.docker.com/r/unders/bates/

##
## BASE
BASE_IMAGE=$(PROG):$(VERSION)-base
BASE_REMOTE_IMAGE=$(DOCKER_USER)/$(BASE_IMAGE)

.PHONY: help
help:
	@echo ""
	@echo "Commands:"
	@echo "    make login             # login to a Docker registry"
	@echo "    make logout            # logout from a Docker registry"
	@echo ""
	@echo "    make base/build        # build Docker file"
	@echo "    make base/push         # push Docker image $(BASE_REMOTE_IMAGE) to registry"
	@echo "    make base/pull         # pull Docker image $(BASE_REMOTE_IMAGE) from registry"
	@echo "    make base/check        # check that Docker image $(BASE_IMAGE) works"
	@echo ""
	@echo "    make open              # open bates registry in browser"
	@echo "    make tree              # list directories"
	@echo ""


.PHONY: login
login:
	docker login

.PHONY: logout
logout:
	docker logout


##
## BASE
.PHONY: base/build
base/build:
	mkdir -p dist
	cp ../../Makefile dist
	cp -R ../../tests dist
	rm -rf dist/tests/e2e/node_modules
	docker build -f Dockerfile.base.embedded.docker -t $(BASE_IMAGE) .
	rm -rf dist

.PHONY: base/check
base/check:
	docker run $(BASE_IMAGE) make test/api
	docker run $(BASE_IMAGE) make test/e2e

.PHONY: base/push
base/push: base/build base/check
	docker tag $(BASE_IMAGE) $(BASE_REMOTE_IMAGE)
	docker push $(BASE_REMOTE_IMAGE)


.PHONY: base/pull
base/pull:
	docker pull $(BASE_REMOTE_IMAGE)
##
## BASE


.PHONY: open
open:
	open $(WEB)

.PHONY: tree
tree:
	tree -I node_modules
