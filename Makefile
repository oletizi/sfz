.PHONY: get gen build test clean
COVER_FILE=./.coverage.txt

default: test

get:
	go get ./pkg/...

gen: get
	go generate ./pkg/...

build: get
	go build ./pkg/...

test: gen
	go install github.com/dave/courtney && courtney -o $(COVER_FILE) ./pkg/model/... && go tool cover -func $(COVER_FILE)

cover: test
	go tool cover -func $(COVER_FILE)

docker-build:
	(cd docker/builder-ubuntu && make run)