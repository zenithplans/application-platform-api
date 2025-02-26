package httpserver

import "time"

type config struct {
	listenAddr     string
	readTimeout    time.Duration
	writeTimeout   time.Duration
	idleTimeout    time.Duration
	maxHeaderBytes int
}

func (conf *config) WithAddr(addr string) *config {
	conf.listenAddr = addr
	return conf
}

func (conf *config) WithReadTimeout(duration time.Duration) *config {
	conf.readTimeout = duration
	return conf
}

func (conf *config) WithWriteTimeout(duration time.Duration) *config {
	conf.writeTimeout = duration
	return conf
}

func (conf *config) WithIdleTimeout(duration time.Duration) *config {
	conf.idleTimeout = duration
	return conf
}

func (conf *config) WithMaxHeaderBytes(size int) *config {
	conf.maxHeaderBytes = size
	return conf
}

func (conf *config) Build() config {
	return *conf
}

func Defaultconfig() *config {
	return &config{
		listenAddr:     "0.0.0.0:8080",
		readTimeout:    time.Second * 15,
		writeTimeout:   time.Second * 15,
		idleTimeout:    time.Minute * 1,
		maxHeaderBytes: 1 << 20, // 1 MB
	}
}
