SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)

TARGETS_GO :=	$(foreach source, $(SOURCES), $(source)_go)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)

service_name =	$(word 2,$(subst /, ,$1))

.PHONY: build
build: server

server: $(TARGETS_GO) $(TARGETS_TMPL)
	go build -o server .

$(TARGETS_GO): %_go:
	protoc --gogo_out=plugins=grpc:. "$*"
	@mkdir -p services/$(call service_name,$*)/gen/pb
	@mv ./services/$(call service_name,$*)/$(call service_name,$*).pb.go ./services/$(call service_name,$*)/gen/pb/pb.go

$(TARGETS_TMPL): %_tmpl:
	@mkdir -p $(dir $*)gen
	protoc -I. --gotemplate_out=destination_dir=services/$(call service_name,$*)/gen,template_dir=templates:services "$*"
	@rm -rf services/services  # need to investigate why this directory is created
	gofmt -w $(dir $*)gen
