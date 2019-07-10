# LogTool
Golang log tool, based on logrus, integrated with logrotate and caller information, such as function, filename, line...

## Usage
```
log := logtool.GetLogger("./log", "testlog", 7*24*time.Hour, time.Minute)
log.Info("test log tool")
```
