APP_NAME := openfeature-sample
DOCKER_REPO := quay.io/devguyio
VERSION := latest
DOCKER_IMAGE := $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE) -f Containerfile .

docker-push:
	docker push $(DOCKER_IMAGE)

docker-run: docker-build
	docker run -p 8080:8080 $(DOCKER_IMAGE)

k8s-apply:
	kubectl apply -f k8s-deployment.yaml

.PHONY: build run clean docker-build docker-push docker-run k8s-apply
