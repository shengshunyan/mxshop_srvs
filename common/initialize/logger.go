package initialize

import "go.uber.org/zap"

var zapLogger *zap.Logger

func InitLogger() {
	zapLogger, _ = zap.NewDevelopment()

	zap.ReplaceGlobals(zapLogger)
}

func CloseLogger() {
	_ = zapLogger.Sync()
}
