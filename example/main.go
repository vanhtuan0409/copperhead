package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/vanhtuan0409/copperhead"
)

type Config struct {
	HttpPort int           `mapstructure:"http_port" cli:"port" default:"8080" description:"HTTP binding port"`
	Timeout  time.Duration `mapstructure:"timeout" default:"5s" description:"HTTP request timeout"`
}

func (c *Config) String() string {
	s, _ := json.MarshalIndent(c, "", "\t")
	return string(s)
}

func main() {
	cfg := Config{}
	copperhead.Unmarshal(&cfg, reflect.TypeOf(cfg), copperhead.ConfigOptions{})
	fmt.Println(cfg.String())
}
