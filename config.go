package cert_plugin

import (
	"encoding/json"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
)

type Config struct {
	Uri        string `mapstructure:"uri"`
	DB         string `mapstructure:"db"`
	Collection string `mapstructure:"collection"`
}

const Namespace = "github_com/kuno989/cert_plugin"

func ParseConfig(e config.ExtraConfig, logger logging.Logger) *Config {
	v, ok := e[Namespace].(map[string]interface{})
	if !ok {
		return nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		//logger.Error("marshal error")
		return nil
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		//logger.Error("unmarshal error")
		return nil
	}
	return &cfg
}