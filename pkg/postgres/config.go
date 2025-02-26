package postgres

import "time"

type SslMode string

const (
	SslModeDisable SslMode = "disable"
)

type config struct {
	host                  string
	port                  int
	user                  string
	password              string
	dbname                string
	sslmode               SslMode
	maxConnCount          int32
	minConnCount          int32
	maxConnIdleTime       time.Duration
	maxConnLifeTimeJitter time.Duration
	maxConnLifeTime       time.Duration
	healthCheckPeriod     time.Duration
}

func (conf *config) WithHost(host string) *config {
	conf.host = host
	return conf
}

func (conf *config) WithPort(port int) *config {
	conf.port = port
	return conf
}

func (conf *config) WithUser(user string) *config {
	conf.user = user
	return conf
}

func (conf *config) WithPassword(secret string) *config {
	conf.password = secret
	return conf
}

func (conf *config) WithDbname(name string) *config {
	conf.dbname = name
	return conf
}

func (conf *config) WithSslmode(mode SslMode) *config {
	conf.sslmode = mode
	return conf
}

func (conf *config) WithMaxConnCount(count int32) *config {
	conf.maxConnCount = count
	return conf
}

func (conf *config) WithMinConnCount(count int32) *config {
	conf.minConnCount = count
	return conf
}

func (conf *config) WithMaxConnIdleTime(duration time.Duration) *config {
	conf.maxConnIdleTime = duration
	return conf
}

func (conf *config) WithMaxConnLifeTimeJitter(duration time.Duration) *config {
	conf.maxConnLifeTimeJitter = duration
	return conf
}

func (conf *config) WithMaxConnLifeTime(duration time.Duration) *config {
	conf.maxConnLifeTime = duration
	return conf
}

func (conf *config) WithHealthCheckPeriod(duration time.Duration) *config {
	conf.healthCheckPeriod = duration
	return conf
}

func (conf *config) Build() config {
	return *conf
}

func Defaultconfig() *config {
	return &config{
		sslmode:               SslModeDisable,
		host:                  "localhost",
		port:                  5432,
		maxConnCount:          4,
		minConnCount:          1,
		maxConnIdleTime:       30 * time.Minute,
		maxConnLifeTime:       60 * time.Minute,
		maxConnLifeTimeJitter: 0,
		healthCheckPeriod:     1 * time.Minute,
	}
}
