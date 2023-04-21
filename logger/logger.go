package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/fatih/color"
)

var (
	Red         = color.New(color.FgRed)
	Green       = color.New(color.FgGreen)
	Yellow      = color.New(color.FgYellow)
	Blue        = color.New(color.FgBlue)
	Magenta     = color.New(color.FgMagenta)
	Cyan        = color.New(color.FgCyan)
	White       = color.New(color.FgWhite)
	BlueBold    = color.New(color.Bold, color.FgBlue)
	MagentaBold = color.New(color.Bold, color.FgMagenta)
	RedBold     = color.New(color.Bold, color.FgRed)
	YellowBold  = color.New(color.Bold, color.FgYellow)
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
type CustomLogger interface {
	Printf(string, ...interface{})
	Writer() io.Writer
}

// Logger interface
type Logger interface {
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	SetLevel(Level)
}

func NewLogger(customLogger CustomLogger, level Level, colorful bool) Logger {

	lg := logger{
		customLogger: customLogger,
		Colorful:     colorful,
		info:         make([]*content, 0),
		warn:         make([]*content, 0),
		err:          make([]*content, 0),
		Level:        level,
	}

	if colorful {
		lg.info = append(lg.info, &content{"[info] ", false, Green})
		lg.warn = append(lg.warn, &content{"[warn] ", false, Magenta})
		lg.err = append(lg.err, &content{"[error] ", false, Red})
	} else {
		lg.info = append(lg.info, &content{"[info] ", false, White})
		lg.warn = append(lg.warn, &content{"[warn] ", false, White})
		lg.err = append(lg.err, &content{"[error] ", false, White})
	}

	return &lg
}

type logger struct {
	customLogger    CustomLogger
	Colorful        bool
	Level           Level
	info, warn, err []*content
}

type content struct {
	text    string
	newLine bool
	color   *color.Color
}

// Info print info
func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Info {
		l.info = append(l.info, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.info {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
		l.info = l.info[:1]
	}
}

// Warn print warn messages
func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Warn {
		l.warn = append(l.warn, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.warn {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
		l.warn = l.warn[:1]
	}
}

// Error print error messages
func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Error {
		l.err = append(l.err, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.err {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
		l.err = l.err[:1]
	}
}

func (l *logger) SetLevel(lv Level) {
	l.Level = lv
}
