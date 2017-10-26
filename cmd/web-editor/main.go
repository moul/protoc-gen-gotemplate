package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func generate(w http.ResponseWriter, r *http.Request) {
	// read input
	decoder := json.NewDecoder(r.Body)
	type Input struct {
		Protobuf string `json:"protobuf"`
		Template string `json:"template"`
	}
	var input Input
	if err := decoder.Decode(&input); err != nil {
		returnError(w, err)
		return
	}

	// create workspace
	dir, err := ioutil.TempDir("", "pggt")
	if err != nil {
		returnError(w, err)
	}
	defer os.RemoveAll(dir) // clean up
	if err := ioutil.WriteFile(filepath.Join(dir, "example.proto"), []byte(input.Protobuf), 0644); err != nil {
		returnError(w, err)
		return
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "example.output.tmpl"), []byte(input.Template), 0644); err != nil {
		returnError(w, err)
		return
	}

	// generate
	cmd := exec.Command("protoc", "-I"+dir, "--gotemplate_out=template_dir="+dir+",debug=true:"+dir, filepath.Join(dir, "example.proto"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		returnError(w, errors.New(string(out)))
		return
	}

	// read output
	content, err := ioutil.ReadFile(filepath.Join(dir, "example.output"))
	if err != nil {
		returnError(w, err)
		return
	}

	returnContent(w, content)
}

func returnContent(w http.ResponseWriter, output interface{}) {
	payload := map[string]interface{}{
		"output": fmt.Sprintf("%s", output),
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func returnError(w http.ResponseWriter, err error) {
	payload := map[string]interface{}{
		"error": fmt.Sprintf("%v", err),
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("static")))
	r.HandleFunc("/generate", generate)
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	fmt.Printf("Listening on %s...\n", addr)
	h := handlers.LoggingHandler(os.Stderr, r)
	h = handlers.CompressHandler(h)
	h = handlers.RecoveryHandler()(h)
	http.ListenAndServe(addr, r)
}
