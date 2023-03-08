package global

import (
	"database/sql"
	"github.com/cqqqq777/simple_oauth2/app/config"
	"github.com/redis/go-redis/v9"
)

var (
	Config *config.Config
	Mdb    *sql.DB
	Rdb    *redis.Client
)
