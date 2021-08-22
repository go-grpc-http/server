package logginghelper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger
var isProduction = false

// Mix up foreground and background colors, create new mixes!
// var red = color.New(color.BgRed).Add(color.FgWhite)
// var normalRed = color.New(color.FgRed)
// var yellow = color.New(color.FgYellow)

// var blue = color.New(color.BgBlue).Add(color.FgWhite)
// var magenta = color.New(color.BgMagenta).Add(color.FgWhite)
// var cyan = color.New(color.BgCyan).Add(color.FgBlack)
// Init  Init Logger
// maxBackupFileSize,  megabytes
// maxAgeForBackupFile,  days
func Init(fileName string, isProd bool, maxBackupCount int, maxBackupFileSize int, maxAgeForBackupFile int, isSafeMode bool) {

	parentDir := filepath.Dir(fileName)
	os.MkdirAll(parentDir, os.ModePerm)

	isProduction = isProd
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxBackupFileSize, // megabytes
		MaxBackups: maxBackupCount,
		MaxAge:     maxAgeForBackupFile, // days
	})

	// Default make prod
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		zap.ErrorLevel,
	)

	if isSafeMode {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
			w,
			zap.DebugLevel,
		)
	}

	logger = zap.New(core)
	defer logger.Sync()
	sugar = logger.Sugar()
}

// LogDebug logs a message at level Debug on the standard logger.
func LogDebug(args ...interface{}) {
	if isProduction {
		sugar.Debug(fileInfo(2), args)
	} else {
		fmt.Print(" DEBUG - ")
		fmt.Print(" " + fileInfo(2) + " ")
		fmt.Println(args)
		// yellow.DisableColor()
	}
}

// LogInfo logs a message at level Info on the standard logger.
func LogInfo(args ...interface{}) {
	if isProduction {
		sugar.Info(fileInfo(2), args)
	} else {
		fmt.Print(" INFO - ")
		fmt.Print(" " + fileInfo(2) + " ")
		fmt.Println(args)
		// yellow.DisableColor()
	}
}

// LogError logs a message at level Error on the standard logger.
func LogError(args ...interface{}) {
	if isProduction {
		sugar.Error(fileInfo(2), args)
	} else {
		fmt.Println("")
		fmt.Print(" ERROR - ")
		fmt.Println("")
		fmt.Println(" " + fileInfo(2) + " ")
		fmt.Println(args)
		fmt.Println("")
	}
}

// Panic logs a message at level Panic on the standard logger.
func LogPanic(args ...interface{}) {
	if isProduction {
		sugar.Panic(args)
	} else {
		log.Panic(args)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
