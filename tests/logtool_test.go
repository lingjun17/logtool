package tests

import (
	"github.com/verylucky/logtool"
	"testing"
	"time"
)

func Test_LogTool_1(t *testing.T) {
	logtool.Init("../log", "testlog", 7*24*time.Hour, time.Minute)
	logtool.GetLogger().Info("test log tool")
	t.Log("passed test")
}
