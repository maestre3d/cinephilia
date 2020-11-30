package configuration

import "github.com/spf13/viper"

type postgres struct {
	User       string
	Password   string
	Address    string
	Port       int
	Database   string
	SecureMode bool
}

func init() {
	viper.SetDefault("postgresql.user", "postgres")
	viper.SetDefault("postgresql.password", "postgres")
	viper.SetDefault("postgresql.address", "localhost")
	viper.SetDefault("postgresql.port", 5432)
	viper.SetDefault("postgresql.database", "watch_list")
	viper.SetDefault("postgresql.secure_mode", false)
}

func newPostgres() postgres {
	return postgres{
		User:       viper.GetString("postgresql.user"),
		Password:   viper.GetString("postgresql.password"),
		Address:    viper.GetString("postgresql.address"),
		Port:       viper.GetInt("postgresql.port"),
		Database:   viper.GetString("postgresql.database"),
		SecureMode: viper.GetBool("postgresql.secure_mode"),
	}
}
