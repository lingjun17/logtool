package logtool

import (
	"github.com/sirupsen/logrus"
	"time"
)

var Logger = logrus.New()

func Init(logPath string, logFileName string, maxAge time.Duration, rotateTime time.Duration) {

	hook := NewHook()
	Logger.AddHook(hook)

	ConfigLocalFilesystemLogger(Logger, logPath, logFileName, maxAge, rotateTime)
	//logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
}

func GetLogger() *logrus.Logger {
	return Logger
}
