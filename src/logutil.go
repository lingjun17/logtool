package logtool

import (
	"github.com/sirupsen/logrus"
	"time"
)

var logger = logrus.New()

func GetLogger(logPath string, logFileName string, maxAge time.Duration, rotateTime time.Duration) *logrus.Logger{

	hook := NewHook()
	logger.AddHook(hook)

	ConfigLocalFilesystemLogger(logger, logPath, logFileName, maxAge, rotateTime)
	//logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: time.RFC3339Nano,
	})

	return logger
}
