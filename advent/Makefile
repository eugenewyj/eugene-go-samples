APP?=advent
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PROJECT?=github.com/eugenewyj/go-samples/advent

GOOS?=linux
GOARCH?=amd64

CONTAINER_IMAGE?=docker.io/eugenewyj/${APP}

clean:
	rm -f bin/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o bin/${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true 
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
	-e "PORT=${PORT}" \
	$(APP):$(RELEASE)

test:
	go test -v -race ./...

minikube: push
	for t in $(shell find ./kubernetes/advent -type f -name "*.yaml"); do \
		cat $$t | \
		sed -e 's/{{[ ]*.Release[ ]*}}/$(RELEASE)/g' | \
		sed -e 's/{{[ ]*.ServiceName[ ]*}}/$(APP)/g'; \
		echo ---; \
	done > tmp.yaml
	kubectl apply -f tmp.yaml
