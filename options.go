package gotel

import "time"

type Config struct {
	TraceBatchTimeout time.Duration
	MetricInterval    time.Duration
}

var baseCfg = Config{
	TraceBatchTimeout: 5 * time.Second,
	MetricInterval:    time.Minute,
}

type CfgOptionFunc func(c *Config) error

func WithTraceBatchTimeout(timeout time.Duration) CfgOptionFunc {
	return func(c *Config) error {
		c.TraceBatchTimeout = timeout
		return nil
	}
}

func WithMetricInterval(timeout time.Duration) CfgOptionFunc {
	return func(c *Config) error {
		c.MetricInterval = timeout
		return nil
	}
}
