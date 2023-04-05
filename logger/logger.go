package logger

var _logger = newStdOutLogger()
var limitLevel = LevelDebug
var withCaller = false

type Level int

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return gray(" DEBUG ")
	case LevelInfo:
		return green(" INFO ")
	case LevelWarn:
		return yellow(" WARNING ")
	case LevelError:
		return red(" ERROR ")
	}
	return gray(" DEBUG ")
}

const LevelDebug = Level(0)
const LevelInfo = Level(1)
const LevelWarn = Level(2)
const LevelError = Level(3)

type Logger interface {
	Log(level Level, message string, keyValues ...interface{})
}

func SetLoggerLevel(customLimitLevel Level) {
	limitLevel = customLimitLevel
}

func SetCaller(enable bool) {
	withCaller = enable
}

func SetLogger(customLogger Logger) {
	_logger = customLogger
}

func Default() Logger {
	return _logger
}
