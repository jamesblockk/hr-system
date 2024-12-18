package database

import (
	"hr-system/common/config"
	"log"

	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func Init(cfg *config.Config) {
	db, err := gorm.Open(mysql.Open(cfg.Mysql.DSN), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: &redisCacher{
			rdb: redis.NewClient(&redis.Options{
				Addr:     cfg.Redis.Addr,
				Password: cfg.Redis.Password,
				DB:       cfg.Redis.DB,
			}),
		},
	}}

	if err = db.Use(cachesPlugin); err != nil {
		log.Fatalf("Failed to use to cachesPlugin: %v", err)
	}

	_db = db
}

func GetDB() *gorm.DB {
	return _db
}
