package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Env string

var (
	EnvLocal   = Env("local")
	EnvDev     = Env("dev")
	EnvStaging = Env("staging")
	EnvProd    = Env("prod")
)

// TODO: Read sync.Once
var once sync.Once

var cfg *Config

func New(env Env) *Config {
	return NewWithPath(env, "configs")
}

func NewWithPath(env Env, cfgPath string) *Config {
	if cfg == nil {
		once.Do(func() {
			viper.AutomaticEnv()
			viper.AddConfigPath(cfgPath)
			viper.SetConfigType("yaml")

			var cfgName string

			switch env {
			case EnvLocal:
				cfgName = "config.local"
			case EnvDev:
				cfgName = "config.dev"
			case EnvStaging:
				cfgName = "config.staging"
			case EnvProd:
				cfgName = "config.prod"
			default:
				fmt.Fprintln(os.Stderr, "please specified env")
				os.Exit(1)
			}

			// Read Config
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("error reading config file: %v", err)
			}

			// unmarshal to value of struct
			if err := viper.Unmarshal(&cfg); err != nil {
				log.Fatalf("unable to decode to config struct: %v", err)
			}

			viper.SetConfigName(cfgName)
			switch env {
			case EnvLocal:
				cfg.Env = EnvLocal
			case EnvDev:
				cfg.Env = EnvDev
			case EnvStaging:
				cfg.Env = EnvStaging
			case EnvProd:
				cfg.Env = EnvProd
			default:
				fmt.Fprintln(os.Stderr, "please specified env")
				os.Exit(1)
			}

			_ = viper.BindEnv("database.user_dsn", "USER_DB_DSN")
		})
	}

	return cfg
}

type Config struct {
	Env  Env
	User UserConfig `mapstructure:"usercfg"`
	Jwt  JWTConfig  `mapstructure:"jwt"`
}

type UserConfig struct {
	// TODO: Understand this
	CacheAddrPort string `mapstructure:"cache_addr_port"`
}

type Database struct {
	UserDSN string `mapstructure:"user_db_dsn"`
}

type JWTConfig struct {
	PublicKey  string `mapstructure:"public_key"`
	PrivateKey string `mapstructure:"private_key"`

	AccessTokenExpiration  int `mapstructure:"access_token_expiration"`
	RefreshTokenExpiration int `mapstructure:"refresh_token_expiration"`
	// TODO: Add Issuer, Audience
}
