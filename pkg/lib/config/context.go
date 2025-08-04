package config

import "context"

type ctxKey struct{}

var ConfigKey = ctxKey{}

func FromContext(ctx context.Context) *Config {
	val := ctx.Value(ConfigKey)
	if cfg, ok := val.(*Config); ok {
		return cfg
	}
	return nil
}
