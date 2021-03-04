// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period   time.Duration `config:"period"`
	Api      string        `config:"api"`
	User     string        `config:"user"`
	Password string        `config:"password"`
	Cliente  string        `config:"cliente"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
}
