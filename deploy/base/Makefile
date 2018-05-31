PROG=bates
VERSION=1
DOCKER_USER=unders
WEB=https://hub.docker.com/r/unders/bates/

##
## FULL
FULL_IMAGE=$(PROG):$(VERSION)-full
FULL_REMOTE_IMAGE=$(DOCKER_USER)/$(FULL_IMAGE)

##
## NOCHROME
NOCHROME_IMAGE=$(PROG):$(VERSION)-nochrome
NOCHROME_REMOTE_IMAGE=$(DOCKER_USER)/$(NOCHROME_IMAGE)

##
## USER
USER_IMAGE=$(PROG):$(VERSION)-user
USER_REMOTE_IMAGE=$(DOCKER_USER)/$(USER_IMAGE)

.PHONY: help
help:
	@echo ""
	@echo "Commands:"
	@echo "    make login                 # login to a Docker registry"
	@echo "    make logout                # logout from a Docker registry"
	@echo ""
	@echo "    make full/build            # build Docker image $(FULL_IMAGE)"
	@echo "    make full/check            # check that Docker image $(FULL_IMAGE) works"
	@echo "    make full/push             # push Docker image $(FULL_REMOTE_IMAGE) to registry"
	@echo "    make full/pull             # pull Docker image $(FULL_REMOTE_IMAGE) from registry"
	@echo ""
	@echo "    make nochrome/build        # build Docker image $(NOCHROME_IMAGE)"
	@echo "    make nochrome/check        # check that Docker image $(NOCHROME_IMAGE) works"
	@echo "    make nochrome/push         # push Docker image $(NOCHROME_REMOTE_IMAGE) to registry"
	@echo "    make nochrome/pull         # pull Docker image $(NOCHROME_REMOTE_IMAGE) from registry"
	@echo ""
	@echo "    make user/build            # build Docker image $(USER_IMAGE)"
	@echo "    make user/check            # check that Docker image $(USER_IMAGE) works"
	@echo "    make user/push             # push Docker image $(USER_REMOTE_IMAGE) to registry"
	@echo "    make user/pull             # pull Docker image $(USER_REMOTE_IMAGE) from registry"
	@echo ""
	@echo "    make open                  # open bates registry in browser"
	@echo "    make tree                  # list directories"
	@echo ""


.PHONY: login
login:
	docker login

.PHONY: logout
logout:
	docker logout


##
## FULL
.PHONY: full/build
full/build:
	docker build -f Dockerfile.full.docker -t $(FULL_IMAGE) .

.PHONY: full/check
full/check:
	docker run $(FULL_IMAGE) node -v
	docker run $(FULL_IMAGE) go version

.PHONY: full/push
full/push: full/build full/check
	docker tag $(FULL_IMAGE) $(FULL_REMOTE_IMAGE)
	docker push $(FULL_REMOTE_IMAGE)

.PHONY: full/pull
full/pull:
	docker pull $(FULL_REMOTE_IMAGE)
##
## FULL

##
## NOCHROME
.PHONY: nochrome/build
nochrome/build:
	docker build -f Dockerfile.full.nochrome.docker -t $(NOCHROME_IMAGE) .

.PHONY: nochrome/check
nochrome/check:
	docker run $(NOCHROME_IMAGE) node -v
	docker run $(NOCHROME_IMAGE) go version

.PHONY: nochrome/push
nochrome/push: nochrome/build nochrome/check
	docker tag $(NOCHROME_IMAGE) $(NOCHROME_REMOTE_IMAGE)
	docker push $(NOCHROME_REMOTE_IMAGE)

.PHONY: nochrome/pull
nochrome/pull:
	docker pull $(NOCHROME_REMOTE_IMAGE)
##
## NOCHROME

##
## USER
.PHONY: user/build
user/build:
	docker build -f Dockerfile.nochrome.user.docker -t $(USER_IMAGE) .

.PHONY: user/check
user/check:
	docker run $(USER_IMAGE) node -v
	docker run $(USER_IMAGE) go version

.PHONY: user/push
user/push: user/build user/check
	docker tag $(USER_IMAGE) $(USER_REMOTE_IMAGE)
	docker push $(USER_REMOTE_IMAGE)

.PHONY: user/pull
user/pull:
	docker pull $(USER_REMOTE_IMAGE)
##
## NOCHROME


.PHONY: open
open:
	open $(WEB)
