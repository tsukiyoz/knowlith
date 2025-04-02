package log

import (
	"context"
	"sync"
)

type Logger interface {
	W(ctx context.Context) Logger
	Debugw(msg string, kvs ...any)
	Infow(msg string, kvs ...any)
	Warnw(msg string, kvs ...any)
	Errorw(msg string, kvs ...any)
	Panicw(msg string, kvs ...any)
	Fatalw(msg string, kvs ...any)
	Sync()
}

var (
	mu  sync.Mutex
	std = New(NewOptions())
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = New(opts)
}

func New(opts *Options) Logger {
	return newZapLogger(opts)
}

func Debugw(msg string, kvs ...any) {
	std.Debugw(msg, kvs)
}

func Infow(msg string, kvs ...any) {
	std.Infow(msg, kvs)
}

func Warnw(msg string, kvs ...any) {
	std.Warnw(msg, kvs)
}

func Errorw(msg string, kvs ...any) {
	std.Errorw(msg, kvs)
}

func Panicw(msg string, kvs ...any) {
	std.Panicw(msg, kvs)
}

func Fatalw(msg string, kvs ...any) {
	std.Fatalw(msg, kvs)
}

func W(ctx context.Context) Logger {
	return std
}

func Sync() {
	std.Sync()
}
