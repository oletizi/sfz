.PHONY: build clean run

BUILD=build
ANTLR_LIB=/usr/local/lib/antlr-4.8-complete.jar

default: build

build: $(BUILD)/Sfz.g4
	echo "antlr lib: $(ANTLR_LIB)"
	mkdir -p $(BUILD) && cp Sfz.g4 $(BUILD) && cd $(BUILD) && antlr Sfz.g4 && javac -cp $(ANTLR_LIB) *.java

test:  build
	cd $(BUILD) && cp ../test/resources/test.sfz . && grun Sfz r test.sfz

clean:
	rm -rf $(BUILD)