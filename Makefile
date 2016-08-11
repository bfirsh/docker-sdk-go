IMAGE_NAME=docker-sdk-go

build:
	docker build -t $(IMAGE_NAME) .

test: build
	docker run -v /var/run/docker.sock:/var/run/docker.sock $(IMAGE_NAME) go test ./

