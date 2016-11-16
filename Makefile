setup:
	go get -u github.com/google/go-github/github
	go get -u github.com/jessevdk/go-flags
	go get -u github.com/joho/godotenv
	go get -u golang.org/x/oauth2

build:
	go build

run:
	go run *.go

test:
	go test

travis:
	go fmt
	go test
