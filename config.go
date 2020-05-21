package copperhead

import (
	"reflect"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ConfigOptions struct {
	EnvPrefix string
}

func Unmarshal(cfg interface{}, t reflect.Type, options ConfigOptions) (err error) {
	var cfgPath string
	pflag.StringVar(&cfgPath, "config", "config.yaml", "Path to config file")
	initViper(t)
	pflag.Parse()

	viper.SetConfigFile(cfgPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetEnvPrefix(options.EnvPrefix)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	err = viper.Unmarshal(&cfg)
	return
}

func initViper(rt reflect.Type, parts ...string) {
	for i := 0; i < rt.NumField(); i++ {
		t := rt.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch t.Type.Kind() {
		case reflect.Struct: // Handle nested struct
			if tv == ",squash" || tv == "" {
				initViper(t.Type, parts...)
			} else {
				initViper(t.Type, append(parts, tv)...)
			}

		default: // Handle leaf field
			keyPath := strings.Join(append(parts, tv), ".")
			viper.SetDefault(keyPath, t.Tag.Get("default"))
			cliFlag := t.Tag.Get("cli")
			if cliFlag == "" {
				cliFlag = strings.ReplaceAll(keyPath, "_", "-")
			}
			pflag.String(
				cliFlag,
				t.Tag.Get("default"),
				t.Tag.Get("description"),
			)
			viper.BindPFlag(
				keyPath,
				pflag.Lookup(cliFlag),
			)
		}
	}
}
