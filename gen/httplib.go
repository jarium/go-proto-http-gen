package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type HttpLibGen interface {
	Imports(g *protogen.GeneratedFile)
	ServerRegisterFunc(g *protogen.GeneratedFile, hs HttpService)
	HandlerFunc(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method)
}

type Net struct {
}

func (ng Net) Imports(g *protogen.GeneratedFile) {
	g.P(`"github.com/gorilla/mux"`)
	g.P(`"encoding/json"`)
}

func (ng Net) ServerRegisterFunc(g *protogen.GeneratedFile, hs HttpService) {
	g.P("func Register", hs.S.GoName, "HTTPServer(r *mux.Router, srv ", hs.S.GoName, "HTTPServer) {")
	for _, method := range hs.Methods {
		if getPattern := method.Rule.GetGet(); getPattern != "" {
			g.P(`r.HandleFunc("`, getPattern, `", _`, hs.S.GoName, "_", method.M.GoName, `_HTTP_Handler(srv)).Methods("GET")`)
		} else if postPattern := method.Rule.GetPost(); postPattern != "" {
			g.P(`r.HandleFunc("`, postPattern, `", _`, hs.S.GoName, "_", method.M.GoName, `_HTTP_Handler(srv)).Methods("POST")`)
		}
	}
	g.P("}")
}

func (ng Net) HandlerFunc(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	g.P("func _", service.GoName, "_", method.GoName, `_HTTP_Handler(srv `, service.GoName, `HTTPServer) http.HandlerFunc {`)
	g.P("return func(w http.ResponseWriter, r *http.Request) {")
	g.P(`w.Header().Set("Content-Type", "application/json")`)
	g.P("var in ", method.Input.GoIdent)
	g.P("if err := json.NewDecoder(r.Body).Decode(&in); err != nil {")
	g.P(`http.Error(w, err.Error(), http.StatusBadRequest)`)
	g.P("return")
	g.P("}")
	g.P("out, err := srv.", method.GoName, `(r.Context(), &in)`)
	g.P("if err != nil {")
	g.P("var apiErr apierror.IError")
	g.P("if errors.As(err, &apiErr) {")
	g.P(`r = r.WithContext(context.WithValue(r.Context(), "error", apiErr.Unwrap()))`)
	g.P(`http.Error(w, apiErr.GetMessage(), apiErr.GetStatusCode())`)
	g.P("return")
	g.P("}")
	g.P(`r = r.WithContext(context.WithValue(r.Context(), "error", err))`)
	g.P(`http.Error(w, "internal server error", http.StatusInternalServerError)`)
	g.P("return")
	g.P("}")
	g.P("json.NewEncoder(w).Encode(out)")
	g.P("}")
	g.P("}")
	g.P()
}

type Gin struct {
}

func (gg Gin) Imports(g *protogen.GeneratedFile) {
	g.P(`"github.com/gin-gonic/gin"`)
}

func (gg Gin) ServerRegisterFunc(g *protogen.GeneratedFile, hs HttpService) {
	g.P("func Register", hs.S.GoName, "HTTPServer(r *gin.Engine, srv ", hs.S.GoName, "HTTPServer) {")
	for _, method := range hs.Methods {
		if getPattern := method.Rule.GetGet(); getPattern != "" {
			g.P(`r.GET("`, getPattern, `", _`, hs.S.GoName, "_", method.M.GoName, `_HTTP_Handler(srv))`)
		} else if postPattern := method.Rule.GetPost(); postPattern != "" {
			g.P(`r.POST("`, postPattern, `", _`, hs.S.GoName, "_", method.M.GoName, `_HTTP_Handler(srv))`)
		}
	}
	g.P("}")
}

func (gg Gin) HandlerFunc(g *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	g.P("func _", service.GoName, "_", method.GoName, `_HTTP_Handler(srv `, service.GoName, `HTTPServer) func(c *gin.Context) {`)
	g.P("return func(c *gin.Context) {")
	g.P("var in ", method.Input.GoIdent)
	g.P("if err := c.ShouldBind(&in); err != nil {")
	g.P(`c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})`)
	g.P("return")
	g.P("}")
	g.P("out, err := srv.", method.GoName, `(c.Request.Context(), &in)`)
	g.P("if err != nil {")
	g.P("var apiErr apierror.IError")
	g.P("if errors.As(err, &apiErr) {")
	g.P("c.Error(apiErr.Unwrap())")
	g.P(`c.JSON(apiErr.GetStatusCode(), gin.H{"error": apiErr.GetMessage()})`)
	g.P("return")
	g.P("}")
	g.P("c.Error(err)")
	g.P(`c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})`)
	g.P("return")
	g.P("}")
	g.P("c.JSON(http.StatusOK, out)")
	g.P("}")
	g.P("}")
	g.P()
}
