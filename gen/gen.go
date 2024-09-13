package gen

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type HttpService struct {
	S       *protogen.Service
	Methods []HttpMethod
}

type HttpMethod struct {
	M    *protogen.Method
	Rule *annotations.HttpRule
}
