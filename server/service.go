package server

import (
	"context"
	"encoding/json"
	"fmt"

	adexp "github.com/florian74/assignement/adexp"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Fpl = adexp.Fpl

func pushToRedis(fpl Fpl) {

	fmt.Println("pushing to redis: " + fpl.IfplId)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	err := rdb.Set(ctx, fpl.IfplId, toJson(fpl), 0).Err()
	if err != nil {
		panic(err)
	}
}

func getFromRedis(fplId string) Fpl {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	val, err := rdb.Get(ctx, fplId).Result()
	if err == redis.Nil {
		panic(nil)
	} else if err != nil {
		panic(err)
	}
	return fromJson(val)

}

func deleteFromRedis(fplId string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	err := rdb.Set(ctx, fplId, "", 0).Err()
	if err == redis.Nil {
		panic(nil)
	} else if err != nil {
		panic(err)
	}

}

func getAllKeys() []Fpl {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	vals := rdb.Keys(ctx, "*").Val()
	ret := make([]Fpl, len(vals))
	for i := 0; i < len(vals); i++ {
		if vals[i] != "" {
			ret[i] = getFromRedis(vals[i])
		}
	}
	return ret

}

func toJson(fpl Fpl) []byte {
	msg, err := json.Marshal(fpl)
	if err != nil {
		panic(err)
	}
	return msg
}

func fromJson(data string) Fpl {
	var dat Fpl
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		panic(err)
	}
	return dat
}
