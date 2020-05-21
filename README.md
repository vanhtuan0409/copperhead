# Copperhead

A wrapper around [spf13/viper](https://github.com/spf13/viper) and [spf13/pflag](https://github.com/spf13/pflag) with opinionted interface.

### Installation

Install by running:

```
go get -u github.com/vanhtuan0409/copperhead
```

### Usage

Copperhead support unmarshal struct using values from

1/ Command line args
2/ Environment variables
3/ Config file

Example usage:

```go
import (
	"reflect"
	"time"

	"github.com/vanhtuan0409/copperhead"
)

type Config struct {
	HttpPort int           `mapstructure:"http_port" cli:"port" default:"8080" description:"HTTP binding port"`
	Timeout  time.Duration `mapstructure:"timeout" default:"5s" description:"HTTP request timeout"`
}

func main() {
	cfg := Config{}
	copperhead.Unmarshal(&cfg, reflect.TypeOf(cfg), copperhead.ConfigOptions{})
}
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

Nested struct will be key path will be intepreted with `.`
