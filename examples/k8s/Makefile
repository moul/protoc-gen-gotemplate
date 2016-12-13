SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)

TARGETS_K8S :=	$(foreach source, $(SOURCES), $(source)_k8s)

.PHONY: build
build: $(TARGETS_K8S)

$(TARGETS_K8S): %_k8s:
	@mkdir -p $(dir $*)gen
	protoc $(PROTOC_OPTS) --gotemplate_out=debug=true,template_dir=./templates:$(dir $*)gen "$*"