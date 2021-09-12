package logger

import (
	"context"
)

// Console Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Log Level
type Level int

const (
	Silent Level = iota
	Error
	Warn
	Info
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

// Logger interface
type Logger interface {
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	SetLevel(Level)
}

func NewLogger(writer Writer, level Level, colorful bool) Logger {
	infoStr := "[info] "
	warnStr := "[warn] "
	errStr := "[error] "

	if colorful {
		infoStr = Green + "[info] " + Reset
		warnStr = Magenta + "[warn] " + Reset
		errStr = Red + "[error] " + Reset
	}

	return &logger{
		Writer:   writer,
		Colorful: colorful,
		infoStr:  infoStr,
		warnStr:  warnStr,
		errStr:   errStr,
		Level:    level,
	}
}

type logger struct {
	Writer
	Colorful                 bool
	Level                    Level
	infoStr, warnStr, errStr string
}

// Info print info
func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Info {
		l.Printf(l.infoStr+msg, data...)
	}
}

// Warn print warn messages
func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Warn {
		l.Printf(l.warnStr+msg, data...)
	}
}

// Error print error messages
func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Error {
		l.Printf(l.errStr+msg, data...)
	}
}

func (l *logger) SetLevel(lv Level) {
	l.Level = lv
}
