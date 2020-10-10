.PHONY: build_java build clean docker docker-intermediate docker-ci run test test-go test-java
ANTLR_LIB=/usr/local/lib/antlr-4.8-complete.jar
PACKAGE_JAVA=org.letizi.sfz.parser

default: all

all: build-java build-go


build-go: install-go
	antlr -Dlanguage=Go -o go/parser Sfz.g4

build-java:
	antlr -package $(PACKAGE_JAVA) -o src/main/java/org/letizi/sfz/parser Sfz.g4

test: test-java test-go

test-go: build-go
	cd go && go test ./... -v

test-java:
	mvn test

install-go:
	cd ./go && go get ./...

ci-local:
	circleci local execute

docker: docker-intermediate docker-ci

docker-intermediate:
	cd docker/golang-java-intermediate && docker build -t oletizi/golang-java .

docker-ci:
	cd docker/golang-java-antlr && docker build -t oletizi/golang-java-antlr .
	docker push oletizi/golang-java-antlr
