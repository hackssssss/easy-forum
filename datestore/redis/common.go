package redis

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive: 15,
		MaxIdle:   5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"127.0.0.1:6379",
			)
		},
	}
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("ping"); err != nil {
		panic(err)
	}
}
