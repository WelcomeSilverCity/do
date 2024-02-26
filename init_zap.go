package do

import "go.uber.org/zap"

func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	//cfg.OutputPaths = []string{
	//	"./myproject.log",
	//}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}
