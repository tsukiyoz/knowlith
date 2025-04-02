package log

import (
	"context"
	"time"

	"github.com/tsukiyoz/knowlith/internal/pkg/contextx"
	"github.com/tsukiyoz/knowlith/internal/pkg/known"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// 将 Options 中的日志级别（字符串）转换为 zapcore.Level 类型
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// 如果指定了非法的日志级别，则默认使用 info 级别
		zapLevel = zapcore.InfoLevel
	}

	// 创建 encoder 配置，用于控制日志的输出格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "msg"
	encoderConfig.TimeKey = "ts"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数，将 time.Duration 序列化为经过的毫秒数的浮点数
	// 毫秒数比默认的秒数更精确
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// 创建构建 zap.Logger 需要的配置
	cfg := &zap.Config{
		// 是否在日志中显示调用日志所在的文件和行号，例如：`"caller":"miniblog/miniblog.go:75"`
		DisableCaller: opts.DisableCaller,
		// 是否禁止在 panic 及以上级别打印堆栈信息
		DisableStacktrace: opts.DisableStacktrace,
		// 指定日志级别
		Level: zap.NewAtomicLevelAt(zapLevel),
		// 指定日志显示格式，可选值：console, json
		Encoding:      opts.Format,
		EncoderConfig: encoderConfig,
		// 指定日志输出位置
		OutputPaths: opts.OutputPaths,
		// 设置 zap 内部错误输出位置
		ErrorOutputPaths: []string{"stderr"},
	}

	// 使用 cfg 创建 *zap.Logger 对象
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	// 将标准库的 log 输出重定向到 zap.Logger
	zap.RedirectStdLog(z)

	return &zapLogger{z: z}
}

type zapLogger struct {
	z *zap.Logger
}

var _ Logger = (*zapLogger)(nil)

func (z *zapLogger) W(ctx context.Context) Logger {
	lc := z.clone()

	contextExtractors := map[string]func(context.Context) string{
		known.XRequestID: contextx.RequestID,
		known.XUserID:    contextx.UserID,
	}

	for field, extractor := range contextExtractors {
		if v := extractor(ctx); v != "" {
			lc.z = lc.z.With(zap.String(field, v))
		}
	}

	return lc
}

func (z *zapLogger) clone() *zapLogger {
	c := *z
	return &c
}

func (z *zapLogger) Debugw(msg string, kvs ...any) {
	z.z.Sugar().Debugw(msg, kvs)
}

func (z *zapLogger) Infow(msg string, kvs ...any) {
	z.z.Sugar().Infow(msg, kvs)
}

func (z *zapLogger) Warnw(msg string, kvs ...any) {
	z.z.Sugar().Warnw(msg, kvs)
}

func (z *zapLogger) Errorw(msg string, kvs ...any) {
	z.z.Sugar().Errorw(msg, kvs)
}

func (z *zapLogger) Panicw(msg string, kvs ...any) {
	z.z.Sugar().Panicw(msg, kvs)
}

func (z *zapLogger) Fatalw(msg string, kvs ...any) {
	z.z.Sugar().Fatalw(msg, kvs)
}

func (z *zapLogger) Sync() {
	_ = z.z.Sync()
}
