package usecases

import (
	"default-repo/libs/log/models"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var LogEntry LogEntryInterface

type LogEntryInterface interface {
	OnInfo(args models.LogInfo)
	OnDebug(args models.LogDebug)
	OnError(args models.LogError)
	OnExternal(args models.LogExternal)
	OnQuery(args models.LogQuery)
}

type logEntry struct {
	log        *logrus.Logger
	defaultLog models.LogDefault
}

func NewLogEntry(level logrus.Level, defaultLog models.LogDefault) LogEntryInterface {
	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableTimestamp:  true,
		PrettyPrint:       viper.GetBool("logRus.jsonFormatter"),
		DisableHTMLEscape: viper.GetBool("logRus.htmlEscape"),
	})
	log.SetOutput(os.Stdout)
	LogEntry = &logEntry{
		log:        log,
		defaultLog: defaultLog,
	}

	return LogEntry
}

func (l *logEntry) OnInfo(args models.LogInfo) {
	args.LogDefault = l.defaultLog
	l.log.Info(args)
}

func (l *logEntry) OnDebug(args models.LogDebug) {
	args.LogDefault = l.defaultLog
	l.log.Debug(args)
}

func (l *logEntry) OnError(args models.LogError) {
	args.LogDefault = l.defaultLog
	l.log.Error(args)
}

func (l *logEntry) OnExternal(args models.LogExternal) {
	args.LogDefault = l.defaultLog
	l.log.Info(args)
}

func (l *logEntry) OnQuery(args models.LogQuery) {
	args.LogDefault = l.defaultLog
	l.log.Info(args)
}
