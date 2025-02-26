package logger

import (
	"io"
	"log/slog"
	"strings"
)

// An abstraction over the default [slog.Logger].
// It sets up a new logger instance with the defined configuration
// and writes to io.
func New(w io.Writer, conf config) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			w,
			&slog.HandlerOptions{
				Level:       conf.level,
				AddSource:   conf.addSource,
				ReplaceAttr: replaceAttr(conf.maskKeys, conf.maskVal),
			},
		).WithAttrs(conf.attr),
	)
}

// replaces all `maskKeys` present in the current log attribute key (case-insensitive)
// with `maskVal`
func replaceAttr(maskKeys []string, maskVal string) func([]string, slog.Attr) slog.Attr {

	return func(groups []string, attr slog.Attr) slog.Attr {

		for _, maskKey := range maskKeys {
			if strings.EqualFold(maskKey, attr.Key) {
				attr.Value = slog.StringValue(maskVal)
				break
			}
		}

		return attr
	}
}
