# protoc-gen-gotemplate
:open_file_folder: protocol generator + golang text/template (protobuf)

Generic protocol buffer generator backed by Golang's [text/template](https://golang.org/pkg/text/template).

---

This is a generator plugin for the Google Protocol Buffers compiler (`protoc`).

The plugin can generate files based on a template directory using the [Golang's `text/template`](https://golang.org/pkg/text/template/) engine.

## Usage

`protoc-gen-gotemplate` requires a **template_dir** directory *(by default `./templates`)*.

Every files ending with `.tmpl` will be processed and written in the destination folder, following the file hierarchy of the `template_dir`, and removing the `.tmpl` extension.

---

```console
$> ls -R
input.proto     templates/doc.txt.tmpl      templates/config.json.tmpl
$> protoc --gotemplate_out=. input.proto
$> ls -R
input.proto     templates/doc.txt.tmpl      templates/config.json.tmpl
doc.txt         config.json
```

---

You can specify a custom `template_dir` or enable `debug`:

```console
$> protoc --gotemplate_out=debug=true,template_dir=/path/to/template/directory:. input.proto
```

---

See [examples](./examples).

## Funcmap

This project uses [moul/funcmap](https://github.com/moul/funcmap) library to extend the builtin [text/template](https://golang.org/pkg/text/template) helpers.

Non-exhaustive list of new helpers:

* `json`
* `prettyjson`
* `indent`
* `split`
* `join`
* `title`
* `unexport`
* `add`
* `trimspace`
* `lower`
* `upper`
* `rev`
* `int`

See the project for the complete list.

## Install

* Install the **Go** compiler and tools from https://golang.org/doc/install
* Install **protobuf**: `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
* Install **protoc-gen-gotemplate**: `go get -u github.com/moul/protoc-gen-gotemplate`

## License

MIT
