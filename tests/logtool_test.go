package tests

import (
	"LogTool/src"
	"testing"
	"time"
)

func Test_LogTool_1(t *testing.T) {
	log := logtool.GetLogger("../log", "testlog", 7*24*time.Hour, time.Minute)
	log.Info("test log tool")
	t.Log("passed test")
}

