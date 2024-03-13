package log

import "github.com/sirupsen/logrus"

// Log 日志记录器
var Log = logrus.WithFields(logrus.Fields{})

// 初始化
func init() {
	Log.Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	Log.Logger.SetLevel(logrus.DebugLevel)
}
