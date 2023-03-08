package boot

import (
	"context"
	"database/sql"
	"fmt"
	g "github.com/cqqqq777/simple_oauth2/app/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
)

func DatabaseInit() {
	MysqlInit()
	RedisInit()
}

func MysqlInit() {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		g.Config.Database.Mysql.Username,
		g.Config.Database.Mysql.Password,
		g.Config.Database.Mysql.Host,
		g.Config.Database.Mysql.Port,
		g.Config.Database.Mysql.DbName,
	))
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	g.Mdb = db
}

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", g.Config.Database.Redis.Host, g.Config.Database.Redis.Port),
		Password: g.Config.Database.Redis.Password,
		DB:       g.Config.Database.Redis.DB,
		PoolSize: g.Config.Database.Redis.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err.Error())
	}
	g.Rdb = client
}
