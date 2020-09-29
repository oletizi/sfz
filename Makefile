.PHONY: build_java build clean run
ANTLR_LIB=/usr/local/lib/antlr-4.8-complete.jar
PACKAGE_JAVA=org.letizi.sfz.parser

build-go:
	antlr -Dlanguage=Go -o go/parser Sfz.g4

build-java:
	antlr -package $(PACKAGE_JAVA) -o src/main/java/org/letizi/sfz/parser Sfz.g4