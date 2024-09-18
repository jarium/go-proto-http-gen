package gen

type TemplateData struct {
	Entrance string
	GoMod    string
	Version  string
	Package  string
	Services []Service
}

type Service struct {
	Name    string
	Methods []Method
}

type Method struct {
	Name          string
	Uri           string
	RequestMethod string
	In            string
	Out           string
}
