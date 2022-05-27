package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var singleLogger *zap.SugaredLogger
var once sync.Once

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		path := fmt.Sprintf("%s/logs", filepath.Dir(ex))
		logFile := fmt.Sprintf("%s/log.log", path)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
		f := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		}

		mr := io.MultiWriter(f, os.Stdout)

		w := zapcore.AddSync(mr)

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.DebugLevel,
		)
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0)) // skip : 0(current caller), 1 (current method caller)
		singleLogger = logger.Sugar()
	})
	return singleLogger
}
