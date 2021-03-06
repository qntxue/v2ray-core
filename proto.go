package core

import "path/filepath"

//go:generate go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go:generate go install -v github.com/gogo/protobuf/protoc-gen-gofast@latest
//go:generate go run ./infra/vprotogen/

// ProtoFilesUsingProtocGenGoFast is the map of Proto files
// that use `protoc-gen-gofast` to generate pb.go files
var ProtoFilesUsingProtocGenGoFast = map[string]bool{filepath.Join("proxy", "vless", "encoding", "addons.proto"): true}
