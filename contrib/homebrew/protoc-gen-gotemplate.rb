require "language/go"

class ProtocGenGotemplate < Formula
  desc "protocol generator + golang text/template (protobuf)"
  homepage "https://github.com/moul/protoc-gen-gotemplate"
  url "https://github.com/moul/protoc-gen-gotemplate/archive/v1.0.0.tar.gz"
  sha256 "1ff57cd8513f1e871cf71dc8f2099bf64204af0df1b7397370827083e95bbb82"
  head "https://github.com/moul/protoc-gen-gotemplate.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    ENV["GO15VENDOREXPERIMENT"] = "1"
    (buildpath/"src/github.com/moul/protoc-gen-gotemplate").install Dir["*"]

    system "go", "build", "-o", "#{bin}/protoc-gen-gotemplate", "-v", "github.com/moul/protoc-gen-gotemplate"
  end
end
