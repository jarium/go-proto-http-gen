# protoc-gen-http
Proto plugin that generates go http code from .proto files. 

## Installation
```bash
go get github.com/jarium/protoc-gen-http
``` 

## Generate Http Code
First, make sure you installed Google Protocol Buffers compiler (protoc) installed on your system to be able to run protoc cli commands. Then you can run the following command to generate http code:
```bash
protoc --http_out={path to generated code} --http_opt=paths=source_relative --proto_path=./proto/google {path to .proto file}
``` 

## Example
Example generated code can be found under proto/gen folder under **example_pb** folder