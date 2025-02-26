package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresql struct {
	driver *pgxpool.Pool
	logger *slog.Logger
	conf   config
}

func New(logger *slog.Logger, conf config) *postgresql {
	return &postgresql{
		logger: logger,
		conf:   conf,
	}
}

func (c *postgresql) parseConf() (*pgxpool.Config, error) {
	conf, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			c.conf.user,
			c.conf.password,
			c.conf.host,
			c.conf.port,
			c.conf.dbname,
			c.conf.sslmode,
		),
	)
	if err != nil {
		c.logger.Error("failed to parse database connection string")
		return nil, err
	}
	conf.MaxConns = c.conf.maxConnCount
	conf.MinConns = c.conf.minConnCount
	conf.MaxConnIdleTime = c.conf.maxConnIdleTime
	conf.MaxConnLifetime = c.conf.maxConnLifeTime
	conf.MaxConnLifetimeJitter = c.conf.maxConnLifeTimeJitter
	conf.HealthCheckPeriod = c.conf.healthCheckPeriod

	return conf, nil
}

func (p *postgresql) Close() error {
	if p.driver == nil {
		p.logger.Info("ignoring database connection end request, as instace is not created yet.")
	}

	p.logger.Info("closing database connection...")
	p.driver.Close()
	return nil
}

func (p *postgresql) Connect() error {

	if p.driver != nil {
		p.logger.Info("ignoring database re-connection")
		return nil
	}

	conf, err := p.parseConf()
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		p.logger.Error("failed to create database instance using specified configuration.")
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		p.logger.Error("failed to ping database connection.")
		return err
	}

	p.logger.Info("database connection established successfully.")
	p.driver = pool
	return nil
}

func (p *postgresql) GetConn() *pgxpool.Pool {
	return p.driver
}
