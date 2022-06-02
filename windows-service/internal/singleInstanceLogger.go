package internal

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var singleLogger *zap.SugaredLogger
var once sync.Once

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func AbsPathify(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		pwd := AbsPathify("./logs")
		if err := os.MkdirAll(pwd, os.ModePerm); err != nil {
			panic(err)
		}

		logFile := path.Join(pwd, "scrape-wsvc.log")

		f := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		}

		singleMultiWriter = io.MultiWriter(f, os.Stdout)

		w := zapcore.AddSync(singleMultiWriter)
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			zap.DebugLevel,
		)
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0)) // skip : 0(current caller), 1 (current method caller)
		singleLogger = logger.Sugar()
	})
	return singleLogger
}
