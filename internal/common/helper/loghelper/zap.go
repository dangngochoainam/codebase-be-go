package loghelper

import (
	"context"
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/envhelper"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.SugaredLogger
	L *zap.Logger
}

var (
	Logger   *zapLogger = &zapLogger{}
	DBLogger *zapLogger = &zapLogger{}
)

func InitZap(app string, env string) error {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		NameKey:      "app",
		EncodeName:   zapcore.FullNameEncoder,
	}
	syncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
	)
	logLevel := configLogLevel(env)

	newCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), syncer, zap.NewAtomicLevelAt(logLevel))

	newLogger := zap.New(newCore, zap.AddCaller())
	defer func() {
		_ = newLogger.Sync()
	}()

	newLogger = newLogger.Named(app)
	zap.ReplaceGlobals(newLogger)

	Logger = &zapLogger{newLogger.Sugar(), newLogger}
	return nil
}

func InitZapWithSql(app string, env string, sqlOrmWriter io.Writer) error {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		NameKey:      "app",
		EncodeName:   zapcore.FullNameEncoder,
	}
	syncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(sqlOrmWriter),
	)
	logLevel := configLogLevel(env)

	newCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), syncer, zap.NewAtomicLevelAt(logLevel))

	newLogger := zap.New(newCore, zap.AddCaller())
	defer func() {
		_ = newLogger.Sync()
	}()

	newLogger = newLogger.Named(app)
	zap.ReplaceGlobals(newLogger)

	DBLogger = &zapLogger{newLogger.Sugar(), newLogger}

	return nil
}

func InitZapWithRotatingFile(app, env string) error {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		NameKey:      "app",
		EncodeName:   zapcore.FullNameEncoder,
	}
	syncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(&lumberjack.Logger{
			Filename: "logs/log",
		}),
	)
	logLevel := configLogLevel(env)

	newCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		syncer,
		zap.NewAtomicLevelAt(logLevel),
	)

	newLogger := zap.New(
		newCore,
		zap.AddCaller(),
	)
	defer func() {
		_ = newLogger.Sync()
	}()

	newLogger = newLogger.Named(app)
	zap.ReplaceGlobals(newLogger)
	Logger = &zapLogger{newLogger.Sugar(), newLogger}

	return nil
}

func configLogLevel(defaultEnv string) zapcore.Level {
	env := os.Getenv(envhelper.ENVIRONMENT)
	if env == "" {
		env = defaultEnv
	}

	var level zapcore.Level

	logLevelEnv := os.Getenv(envhelper.LOG_LEVEL)
	switch logLevelEnv {
	case string(commonhelper.LOG_LEVEL__WARN):
		level = zapcore.WarnLevel
	case string(commonhelper.LOG_LEVEL__DEBUG):
		if env == string(commonhelper.ENV__PRD) {
			level = zapcore.InfoLevel
		} else {
			level = zapcore.DebugLevel
		}
	case string(commonhelper.LOG_LEVEL__INFO):
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}

	return level
}

func (l *zapLogger) WithContext(ctx context.Context) *zap.SugaredLogger {
	return l.With(zap.Any("traceId", getTraceIdFromContext(ctx)))
}

func getTraceIdFromContext(ctx context.Context) any {
	return ctx.Value(commonhelper.HeaderKeyType_TraceId)
}
