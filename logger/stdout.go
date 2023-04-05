package logger

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
)

var (
	red    = color.New(color.BgRed).Sprint
	yellow = color.New(color.BgHiYellow).Sprint
	green  = color.New(color.BgGreen).Sprint
	gray   = color.New(color.BgCyan).Sprint
	prefix = color.New(color.FgGreen).Sprint
)

func newStdOutLogger() Logger {
	return &stdOutLogger{
		logger: log.New(os.Stdout, prefix("LOG:")+" ", log.LstdFlags),
	}
}

type stdOutLogger struct {
	logger *log.Logger
}

func jsonEncode(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}

func (s stdOutLogger) Log(level Level, message string, keyValues ...interface{}) {
	if level < limitLevel {
		return
	}
	args := []interface{}{level.String(), "message", message}
	if withCaller {
		args = append(args, "caller", caller(3))
	}
	args = append(args, keyValues...)

	for i, v := range args {
		switch t := v.(type) {
		case []interface{}:
			args[i] = jsonEncode(t)
		}
	}

	s.logger.Println(args...)
}

func caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}
