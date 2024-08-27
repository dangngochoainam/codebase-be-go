package confighelper

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

func Load(cfgMap any, defaultConfig []byte) error {
	viper.SetConfigType("env")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfgMap)
	if err != nil {
		return err
	}
	return nil
}
