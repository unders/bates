PROG=bates
VERSION=1
IMAGE=$(PROG):$(VERSION)
DOCKER_USER=unders
REMOTE_IMAGE=$(DOCKER_USER)/$(IMAGE)
WEB=https://hub.docker.com/r/unders/bates/

.PHONY: help
help:
	@echo ""
	@echo "Commands:"
	@echo "    make login             # login to a Docker registry"
	@echo "    make logout            # logout from a Docker registry"
	@echo ""
	@echo "    make build             # build Docker file"
	@echo "    make push              # push Docker image $(REMOTE_IMAGE) to registry"
	@echo "    make open              # open $(REMOTE_IMAGE) registry in browser"
	@echo "    make pull              # pull Docker image $(REMOTE_IMAGE) from registry"
	@echo "    make check             # check that Docker image $(IMAGE) works"
	@echo "    make tree              # list directories"
	@echo ""

.PHONY: login
login:
	docker login

.PHONY: logout
logout:
	docker logout

.PHONY: build
build:
	docker build -f Dockerfile -t $(IMAGE) .

.PHONY: check
check:
	@docker run -d $(IMAGE) node -v
	@echo -n output:
	@docker ps -lq | xargs docker logs

.PHONY: push
push: build check
	docker tag $(IMAGE) $(REMOTE_IMAGE)
	docker push $(REMOTE_IMAGE)

.PHONY: open
open:
	open $(WEB)

.PHONY: pull
pull:
	docker pull $(REMOTE_IMAGE)

.PHONY: tree
tree:
	tree -I node_modules
