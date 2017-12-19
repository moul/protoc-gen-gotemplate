.PHONY: build
build:
	mkdir -p output
	protoc -I. --gotemplate_out=template_dir=.,debug=true:. *.proto


.PHONY: re
re: clean build


.PHONY: clean
clean:
	rm -rf output
