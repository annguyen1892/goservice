package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

// Config represents database server and credentials
type Config struct {
	Host           string
	Timeout        time.Duration
	Database       string
	Username       string
	Password       string
	Replicasetname string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
