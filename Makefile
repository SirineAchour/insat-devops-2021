
IMAGE ?= "exampleimage"

install:
	go get ./...

test: install
	go test ./...

start_dev: install
	echo "Server started"
	go run main.go

build:
	go build -o project1 main.go 

image:
	docker build . -t ${IMAGE}

start_docker: image
	docker run -it -p 8080:8080 ${IMAGE}