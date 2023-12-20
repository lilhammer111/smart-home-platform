package log

import (
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"github.com/cloudwego/kitex/pkg/klog"
	klogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

func InitKitexLog() {
	klog.Infof("get env success and env :%s", binding.LoadEnv().GetEnv())
	switch binding.LoadEnv().GetEnv() {
	case binding.DevelopmentEnv:
		klog.SetLevel(klog.LevelDebug)
	case binding.ProductionEnv:
		//logger := logrus.New()
		//logger.SetFormatter(&logrus.TextFormatter{
		//	ForceColors:     true,
		//	FullTimestamp:   true,
		//	TimestampFormat: "2006-01-02 15:04:05",
		//	PadLevelText:    true,
		//})
		klogger := klogrus.NewLogger(
		//klogrus.WithLogger(logger),
		)
		klog.SetLogger(klogger)
		klog.SetLevel(getLevelByConfig())
		klog.SetOutput(&lumberjack.Logger{
			Filename:   binding.GetRemoteConf().Log.LogFileName,
			MaxSize:    binding.GetRemoteConf().Log.LogMaxSize,
			MaxBackups: binding.GetRemoteConf().Log.LogMaxBackups,
			MaxAge:     binding.GetRemoteConf().Log.LogMaxAge,
		})
	default:
		klog.Info("entry default switch...")
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

type customFormatter struct{}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logLevel := strings.ToUpper(entry.Level.String())
	message := entry.Message
	caller := entry.Caller

	// 自定义日志格式
	return []byte(fmt.Sprintf("%s\t[ %s ]\n%s\t%s\n\n", logLevel, timestamp, caller, message)), nil
}
