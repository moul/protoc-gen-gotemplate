.PHONY: build
build:
	mkdir -p output
	protoc -I. --gotemplate_out=output *.proto


.PHONY: re
re: clean build


.PHONY: clean
clean:
	rm -rf output
