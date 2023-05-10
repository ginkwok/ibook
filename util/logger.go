package util

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	encoder := getEncoder()

	var writerSyncer zapcore.WriteSyncer
	switch viper.GetString("log.output") {
	case "stdout":
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	case "file":
		writerSyncer = getLogWriter()
	default:
		panic("invalid logging output: " + viper.GetString("log.output"))
	}

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(time.FixedZone("Asia/Shanghai", 8*60*60)).Format(time.RFC3339Nano))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if viper.GetString("log.format") == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	} else {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func getLogWriter() zapcore.WriteSyncer {
	err := os.MkdirAll(viper.GetString("log.dir"), PERM_OF_DIR)
	if err != nil {
		panic(err)
	}

	fileWriter, _, err := zap.Open(
		filepath.Join(
			viper.GetString("log.dir"),
			viper.GetString("log.file"),
		),
	)
	if err != nil {
		panic(err)
	}
	return fileWriter
}

// func getLogWriter(fname string) zapcore.WriteSyncer {
// 	lumberJackLogger := &lumberjack.Logger{
// 		Filename:   fname,
// 		MaxSize:    10,
// 		MaxBackups: -1,
// 		MaxAge:     30,
// 		Compress:   false,
// 		LocalTime:  true,
// 	}
// 	return zapcore.AddSync(lumberJackLogger)
// }
