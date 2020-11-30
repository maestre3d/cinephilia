package configuration

import (
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Service  string
	Stage    string
	Version  string
	Postgres postgres
}

var stageEnv = os.Getenv("WATCH_LIST_STAGE")

func init() {
	viper.SetDefault("cinephilia.service", "watch_list")
	viper.SetDefault("cinephilia.stage", DevStage)
	viper.SetDefault("cinephilia.version", "0.1.0-alpha")
}

func NewConfiguration() (Configuration, error) {
	if err := load(); err != nil {
		return Configuration{}, err
	}

	return Configuration{
		Service:  viper.GetString("cinephilia.service"),
		Stage:    viper.GetString("cinephilia.stage"),
		Version:  viper.GetString("cinephilia.version"),
		Postgres: newPostgres(),
	}, nil
}

func load() error {
	if stageEnv == ProdStage || stageEnv == StagingStage {
		// viper.SetEnvPrefix("service")
		viper.AutomaticEnv()
		return nil
	}

	return readFromFile()
}

func readFromFile() error {
	viper.SetConfigName("cinephilia")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.cinephilia/")
	viper.AddConfigPath("/etc/cinephilia/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return viper.SafeWriteConfig()
		}

		return err
	}

	viper.WatchConfig()
	return nil
}

func (c Configuration) IsDevEnv() bool {
	return c.Stage == DevStage
}

func (c Configuration) IsTestEnv() bool {
	return c.Stage == TestStage
}

func (c Configuration) IsStagingEnv() bool {
	return c.Stage == StagingStage
}

func (c Configuration) IsProdEnv() bool {
	return c.Stage == ProdStage
}
