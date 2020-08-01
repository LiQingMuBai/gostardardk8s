APP=gostardardk8s
CGO_ENABLED ?= 0
GOOS ?= linux
DEBUG_GOGCFLAGS = -gcflags='all=-N -l'
GOGCFLAGS = -ldflags '-s -w'

.PHONY: build
build: clean
	go build -o ${APP} main.go
.PHONY: run
run:
	go run -race main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning" \
		&& go clean

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'


.PHONY: docker-build
docker-build: build
	docker build -t ${APP} .
    docker tag stringifier stringifier:tag

.PHONY: docker-run
docker-run: docker-build
	docker run -itd -p 8000:8000 --name ${APP} ${APP}

clean-docker:
	# Remove gostardardk8s containers
	docker ps -f name=${APP}-* -aq | xargs docker stop
	docker ps -f name=${APP}-* -aq | xargs docker rm
	# Remove old gostardardk8s images
	docker images -a | grep "${APP}" | awk '{print $3}' | xargs docker rmi

	# remove unused volumes
	 docker volume ls -qf dangling=true | xargs -r docker volume rm
