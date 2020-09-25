.PHONY: build-antlr build clean run
SRC_JAVA=src/main/java
SFZ_G4=$(SRC_JAVA)/org/letizi/sfz/parser/Sfz.g4
ANTLR_LIB=/usr/local/lib/antlr-4.8-complete.jar
PACKAGE_JAVA=org.letizi.sfz.parser

default: build

build: build-antlr

build-antlr:
	antlr -visitor -package $(PACKAGE_JAVA) $(SFZ_G4)