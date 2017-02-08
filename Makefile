build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app
docker:
	docker build -t `printenv DOCKER_IMAGE_NAME` .
