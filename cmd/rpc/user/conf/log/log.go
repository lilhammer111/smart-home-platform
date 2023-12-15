package log

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"github.com/cloudwego/kitex/pkg/klog"
	klogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitKitexLog() {
	switch binding.LoadEnv().GetEnv() {
	case "dev":
		log := logrus.New()
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
		logger := klogrus.NewLogger(
			klogrus.WithLogger(log),
		)
		klog.SetLogger(logger)
		klog.SetLevel(getLevelByConfig())
	case "pro":
		logger := klogrus.NewLogger()
		klog.SetLogger(logger)
		klog.SetLevel(getLevelByConfig())
		klog.SetOutput(&lumberjack.Logger{
			Filename:   binding.GetRemoteConf().Log.LogFileName,
			MaxSize:    binding.GetRemoteConf().Log.LogMaxSize,
			MaxBackups: binding.GetRemoteConf().Log.LogMaxBackups,
			MaxAge:     binding.GetRemoteConf().Log.LogMaxAge,
		})
	}
}

func getLevelByConfig() klog.Level {
	level := binding.GetRemoteConf().Log.LogLevel
	logrus.Infof("log level: %s", level)
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
