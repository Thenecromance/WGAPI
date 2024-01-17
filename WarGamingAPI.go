package wgapi

import "strings"

const (
	Host = "https://api.worldoftanks.*"
)

var (
	Logger ILogger     = &defaultLogger{}
	Client IHttpClient = &client{}
)

type ILogger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(msg ...interface{})
}

type IHttpClient interface {
	Request(url string, params map[string]string) ([]byte, error)
}

func applyRegion(param *Param) {
	param.popts.path = strings.Replace(param.popts.path, "*", param.popts.region, 1)
}

// for API usage only
func Request(opts ...ParamOption) ([]byte, error) {
	param := newParam(opts...)
	applyRegion(param)
	return Client.Request(param.popts.path, param.popts.args)
}

func SetLogger(logger ILogger) {
	Logger = logger
}

// SetClient will set the http client
func SetClient(cli IHttpClient) {
	Client = cli
}
