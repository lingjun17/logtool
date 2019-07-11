package logtool

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func ConfigLocalFilesystemLogger(logger *logrus.Logger, logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {

	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"_%Y%m%d"+".log",
		//rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationCount(7),
	)

	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	logger.SetOutput(writer)
}
