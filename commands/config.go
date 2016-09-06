package commands

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/bndw/pick/backends"
	"github.com/bndw/pick/crypto"
	"github.com/mitchellh/go-homedir"
)

var (
	config                Config // holds the global pick config.
	defaultConfigFileTmpl = "%s/.pick/config.toml"
)

func Execute() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	configFile := fmt.Sprintf(defaultConfigFileTmpl, home)
	if err := initConfig(configFile); err != nil {
		return err
	}

	return RootCmd.Execute()
}

type Config struct {
	Encryption crypto.Config
	Storage    backends.Config
}

func initConfig(configFile string) error {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// TODO(): No config found, should we create one?
	} else {
		if _, err := toml.DecodeFile(configFile, &config); err != nil {
			return err
		}
	}

	return nil
}
