package wgapi

const (
	Host = "https://api.worldoftanks.*"
)

var Logger ILogger = &defaultLogger{}
var Client IHttpClient = &client{}

type ILogger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(msg ...interface{})
}

type IHttpClient interface {
	Request(url string, params map[string]string) ([]byte, error)
}

type Params struct {
	region string `default:"asia"`
	args   map[string]string
}

// for API usage only
func Request(opts ...ParamOption) ([]byte, error) {
	param := NewParam(opts...)
	return Client.Request(param.path, param.popts.args)
}

func SetLogger(logger ILogger) {
	Logger = logger
}
func SetClient(cli IHttpClient) {
	Client = cli
}
