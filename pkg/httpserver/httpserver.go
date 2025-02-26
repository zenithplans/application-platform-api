package httpserver

import (
	"log/slog"
	"net/http"
)

type apiserver struct {
	driver *http.Server
}

func New(conf config) *apiserver {

	slog.Debug(
		"setting server with following configuration ...",
		slog.Float64("readTimeout", conf.readTimeout.Seconds()),
		slog.Float64("writeTimeout", conf.writeTimeout.Seconds()),
		slog.Float64("idleTimeout", conf.idleTimeout.Seconds()),
		slog.Int("maxHeaderBytes", conf.maxHeaderBytes),
	)

	return &apiserver{
		driver: &http.Server{
			Addr:           conf.listenAddr,
			Handler:        http.NewServeMux(),
			ReadTimeout:    conf.readTimeout,
			WriteTimeout:   conf.writeTimeout,
			IdleTimeout:    conf.idleTimeout,
			MaxHeaderBytes: conf.maxHeaderBytes,
		},
	}
}

func (s *apiserver) Start() error {
	slog.Info("listening and serving on the specified server address ...")
	return s.driver.ListenAndServe()
}
