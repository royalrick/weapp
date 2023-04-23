package logger

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), Info, true)
	for i := 0; i < 4; i++ {
		l.Info(context.Background(), "request url: %d", i)
	}
}
