package wgapi

// create a empty struct to implement the logger interface
type defaultLogger struct {
}

func (dl *defaultLogger) Debugf(format string, args ...interface{}) {}
func (dl *defaultLogger) Errorf(format string, args ...interface{}) {}
func (dl *defaultLogger) Error(msg ...interface{})                  {}
