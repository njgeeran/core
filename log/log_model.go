package log

import (
	"go.uber.org/zap"
	"time"
)

type loger struct {
	zap 		*zap.SugaredLogger

	duration	*time.Duration
}

func (l *loger)WithDuration(duration time.Duration) *loger {
	tlog := log
	tlog.duration = &duration
	return &tlog
}
func (l *loger)handWith(args ...interface{}) []interface{} {
	if l.duration != nil {
		args = append(args, "duration",l.duration)
	}
	return args
}

func (l *loger)Infow(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues...)
	l.zap.Infow(msg,keysAndValues...)
}
func (l *loger)Info(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Info(args)
}
func (l *loger)Errorw(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues)
	l.zap.Errorw(msg,keysAndValues...)
}
func (l *loger)Error(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Error(args...)
}
func (l *loger)Debugw(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues)
	l.zap.Debugw(msg,keysAndValues...)
}
func (l *loger)Debug(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Debug(args...)
}