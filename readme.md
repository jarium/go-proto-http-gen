# protoc-gen-http
Proto plugin that generates go http code (using <a href="https://github.com/gin-gonic/gin" target="_blank">gin library</a>) from .proto files. 

## Installation
```bash
go install github.com/jarium/protoc-gen-http@latest
``` 

## Generate Http Code
* Install Google Protocol Buffers compiler (protoc) on your system to be able to run protoc cli commands. 
* Get .proto files provided by google for enabling  option (google.api.http) in your proto files.

Then you can run the following command to generate http code:
```bash
protoc --http_out={path to generated code} --http_opt=paths=source_relative --proto_path={path to  option (google.api.http) files} {path to .proto file}
```

## Simpler Way
If you would like a code generator tool that uses this with all dependencies handled internally, check out <a href="https://github.com/jarium/go-proto-cli" target="_blank">go-proto-cli</a>.