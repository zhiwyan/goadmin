package redis

/**
 * redis connnection pool
 * jun.guo@wenba100.com
 *
 */

import (
	"classroom/lib/config"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	PROTOCOL = "tcp" //connection protocol
)

var RedisPool *redis.Pool
var ClassRoomRedisPool *redis.Pool

func InitRedisPool() error {
	server := config.Config.Redis.Server
	password := config.Config.Redis.Password
	maxIdle := config.Config.Redis.RedisMaxIidleConn
	idleTimeout := config.Config.Redis.RedisIdleTimeoutSec
	selectDb := config.Config.Redis.SelectDb

	RedisPool = NewRedisPool(server, password, maxIdle, idleTimeout, selectDb)

	return nil
}

// 初始化 classroom redis
func InitClassRoomRedisPool() error {
	server := config.Config.ClassRoomRedis.Server
	password := config.Config.ClassRoomRedis.Password
	maxIdle := config.Config.ClassRoomRedis.RedisMaxIidleConn
	idleTimeout := config.Config.ClassRoomRedis.RedisIdleTimeoutSec
	selectDb := config.Config.Redis.SelectDb

	ClassRoomRedisPool = NewRedisPool(server, password, maxIdle, idleTimeout, selectDb)

	return nil
}

func CloseRedisPool() {
	if RedisPool != nil {
		RedisPool.Close()
	}
}

func CloseClassRoomRedisPool() {
	if ClassRoomRedisPool != nil {
		ClassRoomRedisPool.Close()
	}
}

/**
 * Redis Pool
 *
 * server serverAddress 127.0.0.1:6379
 * IdleTimeout  超时
 * MaxIdle 连接池最大容量
 * MaxActive 最大活跃数量
 * dbno 选择db127.0.0.1:6379:password:1
 *
 */
func NewRedisPool(server, password string, maxIdle, idleTimeout int, selectDb int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(PROTOCOL, server)
			if err != nil {
				return nil, err
			}

			if len(password) != 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if selectDb >= 0 {
				if _, err := c.Do("SELECT", selectDb); err != nil {
					c.Close()
					fmt.Println(err)
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if t.Add(time.Duration(maxIdle) * time.Second).After(time.Now()) {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetStringFromRedis(key string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}

	return value, nil
}

func SetStringToRedis(key, value string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("set", key, value))
	if err != nil {
		return err
	}

	return nil
}

func SetStringToRedisEX(key, value string, exTime int64) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("set", key, value, "EX", exTime))
	if err != nil {
		return err
	}

	return nil
}

func DelKeyFromRedis(key string) (int64, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("del", key))
}

func SetNXWithExpireToRedis(key, value string, expire int) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	r, err := redis.String(conn.Do("set", key, value, "EX", expire, "NX"))
	if err != nil {
		return r, err
	}

	return r, nil
}

func SetNXToRedis(key, value string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	r, err := redis.String(conn.Do("set", key, value, "NX"))
	if err != nil {
		return r, err
	}

	return r, nil
}

//递增
func Incr(key string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("incr", key)
	return err
}

func Setex(key string, seconds int64, value string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", key, seconds, value)
	return err
}

func HsetStringToRedis(key, field, value string) (int, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("hset", key, field, value))
}

func ExpireKeyToRedis(key string, expire int) (int, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("expire", key, expire))
}

func HgetStringToRedis(key, field string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("hget", key, field))
}

func HdelStringToRedis(key, field string) (int, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("hdel", key, field))
}

func GetStringFromClassRoomRedis(key string) (string, error) {
	conn := ClassRoomRedisPool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}

	return value, nil
}

func SetStringToClassRoomRedis(key, value string) error {
	conn := ClassRoomRedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("set", key, value))
	if err != nil {
		return err
	}

	return nil
}

func HsetStringToClassRoomRedis(key, field, value string) (int, error) {
	conn := ClassRoomRedisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("hset", key, field, value))
}

func HgetStringToClassRoomRedis(key, field string) (string, error) {
	conn := ClassRoomRedisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("hget", key, field))
}

func HdelStringToClassRoomRedis(key, field string) (reply int, err error) {
	conn := ClassRoomRedisPool.Get()
	defer conn.Close()
	reply, err = redis.Int(conn.Do("HDEL", key, field))
	return
}
