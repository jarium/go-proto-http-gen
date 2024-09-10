package apierror

type IError interface {
	GetStatusCode() int
	GetMessage() string
	Error() string
	Unwrap() error
}
