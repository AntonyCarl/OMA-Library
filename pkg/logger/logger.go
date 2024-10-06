package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   false,
	})
	Logger.SetLevel(logrus.InfoLevel)
	Logger.AddHook(&MethodHook{})

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Info("File cant be open, using terminal for logging")
	}
}

type MethodHook struct{}

func (hook *MethodHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *MethodHook) Fire(entry *logrus.Entry) error {
	if entry.HasCaller() {
		entry.Data["method"] = entry.Caller.Function
	}
	return nil
}
