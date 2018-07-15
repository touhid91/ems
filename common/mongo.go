package common

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/go-redis/redis"
	"fmt"
	"github.com/touhid91/ems/conf"
)

func Db(tenant string) (*mongo.Database, error) {
	env := conf.New()

	r := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("localhost:%d", env.Ports.Redis),
	})

	db := r.HGet(tenant, "db").Val()

	c, e := mongo.NewClient(fmt.Sprintf("mongodb://localhost:%d", env.Ports.Mongo))

	if nil != e {
		log.Fatal(e)
		return nil, e
	}

	return c.Database(db), nil
}

