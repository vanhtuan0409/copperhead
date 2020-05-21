# Copperhead

A wrapper around [spf13/viper](https://github.com/spf13/viper) and [spf13/pflag](https://github.com/spf13/pflag) with opinionted interface.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/vanhtuan0409/copperhead)

### Installation

Install by running:

```
go get -u github.com/vanhtuan0409/copperhead
```

### Usage

Copperhead support unmarshal struct using values from

- Command line args
- Environment variables
- Config file

Example usage:

```go
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
```

CLI Help:

```
$ go run main.go -h
Usage of /tmp/go-build191072962/b001/exe/main:
      --config string    Path to config file (default "config.yaml")
      --port string      HTTP binding port (default "8080")
      --timeout string   HTTP request timeout (default "5s")
pflag: help requested
exit status 2
```

### Config

Copperhead provide 3 struct tags:

- **mapstructure**: key for configuration. Example: `mapstructure:"http_port"` will be intepreted as env var `HTTP_PORT` and cli args `--http-port`
  - Enviroment variable will be translated as upper case
  - CLI args will be translated as `strings.ReplaceAll(keyPath, "_", "-")`
- **default**: default value of this field. If cannot be coerce to target type, copperhead will panic
- **cli**: name of CLI args. If not defined, CLI args will be intepreted from mapstructure
- **description**: CLI args description

#### Nested struct

Nested struct will be key path will be intepreted with `.` for cli args and `__` for env var

Example:
```
type Config struct {
	NestedStruct struct {
		MyNumber int `mapstructure:"my_number" default:"10"`
	} `mapstructure:"nested_struct"`
}
```

Will be intepreted as:

```yaml
# Yaml file
nested_struct:
  my_number: 10
```

```sh
$ go run main.go -h
Usage of /tmp/go-build793598133/b001/exe/main:
      --config string                    Path to config file (default "config.yaml")
      --nested-struct.my-number string    (default "10")
pflag: help requested
exit status 2
```

And environment variable `NESTED_STRUCT__MY_NUMBER`
