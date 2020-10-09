.PHONY: build_java build clean run test test-go test-java
ANTLR_LIB=/usr/local/lib/antlr-4.8-complete.jar
PACKAGE_JAVA=org.letizi.sfz.parser

default: all

all: build-java build-go

build-go:
	antlr -Dlanguage=Go -o go/parser Sfz.g4

build-java:
	antlr -package $(PACKAGE_JAVA) -o src/main/java/org/letizi/sfz/parser Sfz.g4

test: test-java test-go

test-go: build-go
	go test ./go/... -v

test-java:
	mvn test

