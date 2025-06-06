package logging

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	mu      sync.Mutex
	loggers = map[string]*logrus.Logger{}
)

// GetLogger returns a logger for the given component. It reads the log level
// from the configuration key "log_levels.<component>". If not set, the global
// "log-level" is used.
func GetLogger(component string) *logrus.Entry {
	mu.Lock()
	defer mu.Unlock()
	if l, ok := loggers[component]; ok {
		return l.WithField("component", component)
	}
	l := logrus.New()
	levelStr := viper.GetString("log_levels." + component)
	if levelStr == "" {
		levelStr = viper.GetString("log-level")
	}
	lvl, err := logrus.ParseLevel(levelStr)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	l.SetLevel(lvl)
	loggers[component] = l
	return l.WithField("component", component)
}
