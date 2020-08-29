package conf

import (
	"arhat.dev/pkg/log"
	"github.com/spf13/pflag"
)

// nolint:lll
type TemplateApplicationGoConfig struct {
	TemplateApplicationGo TemplateApplicationGoAppConfig `json:"templateApplicationGo" yaml:"templateApplicationGo"`
}

type TemplateApplicationGoAppConfig struct {
	Log log.ConfigSet `json:"log" yaml:"log"`

	Foo string `json:"foo" yaml:"foo"`
}

func FlagsForTemplateApplicationGo(prefix string, config *TemplateApplicationGoAppConfig) *pflag.FlagSet {
	fs := pflag.NewFlagSet("app", pflag.ExitOnError)

	fs.StringVar(&config.Foo, prefix+"foo", "bar", "set value of foo")

	return fs
}

func (c *TemplateApplicationGoConfig) GetLogConfig() log.ConfigSet {
	return c.TemplateApplicationGo.Log
}

func (c *TemplateApplicationGoConfig) SetLogConfig(config log.ConfigSet) {
	c.TemplateApplicationGo.Log = config
}
