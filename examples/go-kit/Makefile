SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)

TARGETS_GO :=	$(foreach source, $(SOURCES), $(source)_go)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)

.PHONY: build
build: $(TARGETS_GO) $(TARGETS_TMPL)

$(TARGETS_GO): %_go:
	protoc --gogo_out=. "$*"

$(TARGETS_TMPL): %_tmpl:
	@mkdir -p $(dir $*)gen
	protoc -I. --gotemplate_out=template_dir=templates:services "$*"
	gofmt -w $(dir $*)gen
