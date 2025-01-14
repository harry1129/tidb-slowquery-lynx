package log

import (
	"fmt"
	"os"
	"time"

	gormzap "github.com/things-go/gormzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

var mylogger *zap.Logger
var gormlogger logger.Interface

func InitLogger(logfile string) {
	logFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error open log file %s\n", logfile)
		os.Exit(1)
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(logFile), zapcore.InfoLevel)
	mylogger = zap.New(core, zap.AddCaller())
	gormlogger = gormzap.New(mylogger,
		gormzap.WithConfig(logger.Config{
			SlowThreshold:             120 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Error,
		}),
	)
}

func GetLogger() *zap.Logger {
	return mylogger
}

func GetgormLogger() logger.Interface {
	return gormlogger
}
