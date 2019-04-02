GOOS := linux
GOARCH := amd64
GOLDFLAGS := -ldflags

BPF_DOCKERFILE ?= bpf/Dockerfile
REGISTRY ?= dippynark
IMAGE ?= sockmap
TAG ?= $(shell uname -r)

DEBUG=1

# If you can use docker without being root, you can do "make SUDO="
SUDO=$(shell docker info >/dev/null 2>&1 || echo "sudo -E")

start:
	$(SUDO) docker run --name sockmap -d ${REGISTRY}/${IMAGE}:${TAG}

stop:
	$(SUDO) docker rm -f sockmap

build: bpf docker_build

docker_build:
	$(SUDO) docker build -t ${REGISTRY}/${IMAGE}:${TAG} .

docker_push:
	$(SUDO) docker push ${REGISTRY}/${IMAGE}:${TAG}

bpf: docker_build_bpf install_bpf

docker_build_bpf: docker_build_image
	$(SUDO) docker run --rm -e DEBUG=$(DEBUG) \
		-v $(CURDIR):/src:ro \
		-v $(CURDIR)/bpf:/dist/ \
		-v /usr/src:/usr/src \
		--workdir=/src/bpf \
		$(REGISTRY)/bpf-builder \
		make bpf

docker_build_image:
	$(SUDO) docker build -t $(REGISTRY)/bpf-builder -f $(BPF_DOCKERFILE) .

install_bpf:
	cp -a bpf/bpf_sockmap.go pkg/sockmap/bpf_sockmap.go
	mkdir -p pkg/sockmap/include
	cp -a bpf/include/*.h pkg/sockmap/include
