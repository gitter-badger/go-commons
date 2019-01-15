package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
)

var (
	//default settings
	settings = LogSettings{
		Output:       "std",
		Format:       "text",
		Level:        "info",
		ReportCaller: true,
	}
	lock   sync.Mutex
	logger *logrus.Logger
)

// supporting ini/yaml/json
type LogSettings struct {
	Output       string `json:"output" yaml:"output" ini:"output"`
	Format       string `json:"format" yaml:"format" ini:"format"`
	Level        string `json:"level" yaml:"level" ini:"level"`
	ReportCaller bool   `json:"reportCaller" yaml:"report-caller" ini:"report-caller"`
}

func initLogger(c interface{}) {
	var conf = settings
	if c != nil {
		conf = getConf(c)
	}
	l := logrus.New()
	l.SetOutput(getOutput(conf))
	l.SetFormatter(getFormatter(conf))
	l.SetLevel(getLogLevel(conf))
	l.SetReportCaller(conf.ReportCaller)
	logger = l
}

func GetLogger(c interface{}) *logrus.Logger {
	if logger != nil {
		return logger
	} else {
		lock.Lock()
		initLogger(c)
		lock.Unlock()
	}
	return logger
}

func getConf(raw interface{}) LogSettings {
	if v, ok := raw.(LogSettings); ok {
		return v
	}
	if v, ok := raw.(*LogSettings); ok && v != nil {
		return *v
	}
	getType := reflect.TypeOf(raw)
	getValue := reflect.ValueOf(raw)
	if getType.Kind() == reflect.Struct {
		for i := 0; i < getType.NumField(); i++ {
			value := getValue.Field(i).Interface()
			if reflect.TypeOf(value).Kind() != reflect.Struct {
				continue
			}
			return getConf(value)
		}
		return settings
	} else {
		return settings
	}
}

// get log level, default level info
func getLogLevel(settings LogSettings) logrus.Level {
	switch strings.ToLower(settings.Level) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func getFormatter(c LogSettings) logrus.Formatter {
	switch c.Format {
	case "text":
		return &logrus.TextFormatter{}
	case "json":
		return &logrus.JSONFormatter{}
	default:
		return &logrus.TextFormatter{}
	}
}

func getOutput(c LogSettings) io.Writer {
	switch c.Output {
	case "std":
		return os.Stdout

	default:
		return os.Stdout
	}
}