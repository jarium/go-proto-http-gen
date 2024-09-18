# protoc-gen-http
Proto plugin that generates go http code from .proto files. With **lib** flag, you can specify http lib being used for code gen, currently supporting net/http,gin.

## Installation
```bash
go get github.com/jarium/protoc-gen-http
go install github.com/jarium/protoc-gen-http
``` 

## Generate Http Code
* Install Google Protocol Buffers compiler (protoc) on your system to be able to run protoc cli commands. 
* Get .proto files provided by google for enabling  option (google.api.http) in your proto files.

Then you can run the following command to generate http code:
```bash
protoc --go_out={path to generated go code} --http_out={path to generated http code} --http_opt=lib={net or gin} --proto_path={path to  option (google.api.http) files} --proto_path={path to other proto folders} {path to .proto file}
```

## Example
```bash
go generate
```
The generated code for example.proto will be located inside its directory (example/gen/example_pb)

## Simpler Way
If you would like a code generator tool that uses this with all dependencies handled internally, check out <a href="https://github.com/jarium/go-proto-cli" target="_blank">go-proto-cli</a>.