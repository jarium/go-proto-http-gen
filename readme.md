# protoc-gen-http
Proto plugin that generates go http code (using gin library) from .proto files. 

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
protoc --http_out={path to generated code} --http_opt=paths=source_relative --proto_path={path to  option (google.api.http) files} {path to .proto file}
```