package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger
var SimpleLogger *zap.Logger

func InitLogger() {

	core := zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	SimpleLogger = zap.New(core, zap.AddCaller())
	fmt.Println("Hello InitLogger")

	SugarLogger = SimpleLogger.Sugar()
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func NonSugaredLogger() *zap.Logger {
	return SimpleLogger
}

func SugaredLogger() *zap.SugaredLogger {
	fmt.Println("SugaredLogger")
	return SugarLogger
}
