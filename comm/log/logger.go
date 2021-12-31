package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	logger = zerolog.New(output).With().Timestamp().Logger()

}

// Debugf 打印调试日志
func Debugf(format string, a ...interface{}) {
	logger.Debug().Msg(getCallerName() + fmt.Sprintf(format, a...))
}

// Infof 打印日志
func Infof(format string, a ...interface{}) {
	logger.Info().Msg(getCallerName() + fmt.Sprintf(format, a...))
}

// Errorf 打印错误日志
func Errorf(format string, a ...interface{}) {
	logger.Error().Msg(getCallerName() + fmt.Sprintf(format, a...))
}

func toString(args ...interface{}) string {
	return fmt.Sprint(args...)
}

// Debug 打印调试日志
func Debug(args ...interface{}) {
	logger.Debug().Msg(getCallerName() + toString(args...))
}

// Info 打印日志
func Info(args ...interface{}) {
	logger.Info().Msg(getCallerName() + toString(args...))
}

// Error 打印错误日志
func Error(args ...interface{}) {
	logger.Error().Msg(getCallerName() + toString(args...))
}

func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	fullpath, line := f.FileLine(pc)
	pkgname := "wxcloudrun-wxcomponent"
	name := fullpath[strings.LastIndex(fullpath, pkgname+"/")+len(pkgname+"/"):]
	return fmt.Sprintf("[%s:%d] ", name, line)
}
