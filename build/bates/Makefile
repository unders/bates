PROG=bates
VERSION=2
DOCKER_USER=unders
WEB=https://hub.docker.com/r/unders/bates/
GIT_REPO=https://github.com/unders/bates

##
## BASE
BASE_IMAGE=$(PROG):$(VERSION)-tests-embedded-chrome
BASE_REMOTE_IMAGE=$(DOCKER_USER)/$(BASE_IMAGE)

##
## EXT
EXT_IMAGE=$(PROG):$(VERSION)-tests-external-chrome
EXT_REMOTE_IMAGE=$(DOCKER_USER)/$(EXT_IMAGE)

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
	@echo "    make ext/build         # build Docker file"
	@echo "    make ext/push          # push Docker image $(EXT_REMOTE_IMAGE) to registry"
	@echo "    make ext/pull          # pull Docker image $(EXT_REMOTE_IMAGE) from registry"
	@echo "    make ext/check         # check that Docker image $(EXT_IMAGE) works"
	@echo ""
	@echo "    make git/tag           # set git tag v$(VERSION) and push to repo"
	@echo "    make git/open          # sopen bates git repo in browser: $(GIT_REPO)"
	@echo "    make open              # open bates registry in browser: $(WEB)"
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
	cp -f  dist/tests/e2e/support/jest-puppeteer.config.js dist/tests/e2e/jest-puppeteer.config.js
	rm -rf dist/tests/e2e/node_modules
	docker build -f Dockerfile.embedded.docker -t $(BASE_IMAGE) .
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

##
## EXT
.PHONY: ext/build
ext/build:
	mkdir -p dist
	cp ../../Makefile dist
	cp -R ../../tests dist
	cp -f jest-puppeteer.config.js.external dist/tests/e2e/jest-puppeteer.config.js
	rm -rf dist/tests/e2e/node_modules
	docker build -f Dockerfile.external.docker -t $(EXT_IMAGE) .
	rm -rf dist

.PHONY: ext/check
ext/check:
	docker run $(EXT_IMAGE) make test/api
	docker run $(EXT_IMAGE) make test/e2e

.PHONY: ext/push
ext/push: ext/build ext/check
	docker tag $(EXT_IMAGE) $(EXT_REMOTE_IMAGE)
	docker push $(EXT_REMOTE_IMAGE)


.PHONY: ext/pull
ext/pull:
	docker pull $(EXT_REMOTE_IMAGE)
##
## EXT


.PHONY: git/tag
git/tag:
	git tag v$(VERSION) master
	git push origin v$(VERSION)

.PHONY: git/open
git/open:
	open $(GIT_REPO)

.PHONY: open
open:
	open $(WEB)

