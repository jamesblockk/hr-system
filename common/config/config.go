package config

import (
	"time"
)

type Config struct {
	Mysql MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	DSN                  string
	ParameterizedQueries bool
	MaxIdleConns         int
	MaxOpenConns         int
	ConnMaxLifetime      time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// var _config *Config

// func Init() error {
// 	viper.SetConfigFile("./config.yaml")
// 	if err := viper.ReadInConfig(); err != nil {
// 		return err
// 	}

// 	if err := viper.Unmarshal(&_config); err != nil {
// 		return err
// 	}
// 	fmt.Println("config init")
// 	return nil
// }

func Get() *Config {
	return &Config{
		Mysql: MySQLConfig{
			DSN:                  "root:password@tcp(mysql:3306)/hr_system?charset=utf8mb4&parseTime=True&loc=UTC",
			ParameterizedQueries: false,
			MaxIdleConns:         5,
			MaxOpenConns:         20,
			ConnMaxLifetime:      1 * time.Hour,
		},

		Redis: RedisConfig{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		},
	}
}

func GetLocal() *Config {
	return &Config{
		Mysql: MySQLConfig{
			DSN:                  "root:password@tcp(localhost:3306)/hr_system?charset=utf8mb4&parseTime=True&loc=UTC",
			ParameterizedQueries: false,
			MaxIdleConns:         5,
			MaxOpenConns:         20,
			ConnMaxLifetime:      1 * time.Hour,
		},

		Redis: RedisConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}
}
