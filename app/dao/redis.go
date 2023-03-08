package dao

import (
	"context"
	g "github.com/cqqqq777/simple_oauth2/app/global"
	"time"
)

func SetCode(id string, code int32) error {
	return g.Rdb.SetEx(context.Background(), id, code, time.Minute*10).Err()
}

func GetCode(id string) (int64, error) {
	cmd := g.Rdb.Get(context.Background(), id)
	return cmd.Int64()
}
