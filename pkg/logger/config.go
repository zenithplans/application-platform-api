package logger

import "log/slog"

type config struct {
	level slog.Level
	// additional log attributes for each log
	attr []slog.Attr
	// sensitive data attribute keys that need to be masked before logging.
	maskKeys []string
	// value to replace sensite data with.
	maskVal string
	// adds stack trace of the log as a log attribute.
	addSource bool
}

func (conf *config) WithLevel(level slog.Level) *config {
	conf.level = level
	return conf
}

func (conf *config) WithAttr(attr ...slog.Attr) *config {
	conf.attr = attr
	return conf
}

func (conf *config) WithMaskKeys(keys ...string) *config {
	conf.maskKeys = keys
	return conf
}

func (conf *config) WithMaskVal(maskVal string) *config {
	conf.maskVal = maskVal
	return conf
}

func (conf *config) WithSource() *config {
	conf.addSource = true
	return conf
}

func (conf *config) Build() config {
	return *conf
}

func DefaultConfig() *config {
	return &config{
		level:   slog.LevelInfo,
		maskVal: "***",
	}
}
