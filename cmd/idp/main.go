package main

import (
	"log/slog"
	"math/rand/v2"
	"os"
	"time"

	"github.com/zenithplans/identity-provider-api/pkg/httpserver"
	"github.com/zenithplans/identity-provider-api/pkg/logger"
	"github.com/zenithplans/identity-provider-api/pkg/postgres"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	// LOGGER setup
	//
	slog.SetDefault(
		logger.New(
			os.Stdout,
			logger.DefaultConfig().
				WithLevel(slog.LevelDebug).
				WithAttr(
					slog.String("service-name", "ApplicationPlatformAPI"),
					slog.Int("service-session", rand.Int()),
				).
				WithMaskKeys("x-api-key", "authorization").
				WithSource().
				Build(),
		),
	)

	// DATABASE setup
	//
	db := postgres.New(
		slog.Default().With(slog.String("database-name", "identity-provider-local-db")),
		postgres.Defaultconfig().
			WithUser("dbadmin").
			WithPassword("superadminsecret").
			WithDbname("identity-provider-local-db").
			WithHost("localhost").
			WithPort(5432).
			WithSslmode(postgres.SslModeDisable).
			WithMaxConnCount(4).
			WithMinConnCount(1).
			WithMaxConnIdleTime(30*time.Minute).
			WithMaxConnLifeTime(60*time.Minute).
			WithMaxConnLifeTimeJitter(0).
			WithHealthCheckPeriod(1*time.Minute).
			Build(),
	)

	if err := db.Connect(); err != nil {
		return err
	}

	conn := db.GetConn()
	defer conn.Close()

	// SERVER setup
	//
	srv := httpserver.New(
		httpserver.Defaultconfig().
			WithAddr("0.0.0.0:4444").
			WithReadTimeout(30 * time.Second).
			WithWriteTimeout(30 * time.Second).
			WithIdleTimeout(30 * time.Second).
			Build(),
	)

	return srv.Start()
}
