package wgapi

const (
	Host = "https://api.worldoftanks."
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

type defaultLogger struct {
}

func (dl *defaultLogger) Debugf(format string, args ...interface{}) {}
func (dl *defaultLogger) Errorf(format string, args ...interface{}) {}
func (dl *defaultLogger) Error(msg ...interface{})                  {}

func Request(url string, params map[string]string) ([]byte, error) {
	return Client.Request(url, params)
}

func SetLogger(logger ILogger) {
	Logger = logger
}
func SetClient(cli IHttpClient) {
	Client = cli
}
