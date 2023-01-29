package redis

import (
	"encoding/json"
	"fmt"
	"goadmin/lib/common"
	"goadmin/lib/config"
	"log"
	"testing"

	"github.com/gomodule/redigo/redis"
)

const (
	EXPIRE = 30
)

func init() {
	config.InitConfig("../../config/config.toml")
	var err error
	err = InitRedisPool()
	err = InitClassRoomRedisPool()
	if err != nil {
		panic(err)
	}
}

func TestNewRedisPool(t *testing.T) {
	pool := NewRedisPool("127.0.0.1:6379", "123456", 5, 240, 0)
	conn := pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("SET", "abc", "test"))
	if err != nil {
		panic(err)
	}
}

func TestGetStringFromRedis(t *testing.T) {
	key := "abc"
	result, err := GetStringFromRedis(key)
	if err != nil {
		panic(err)
	}
	log.Println("get string: ", result)
}

func TestSetStringToRedis(t *testing.T) {
	key := "settest"
	value := "valuetest"
	err := SetStringToRedis(key, value)
	if err != nil {
		panic(err)
	}
	result, err := GetStringFromRedis(key)
	if err != nil {
		panic(err)
	}

	if result == value {
		log.Printf("%s === %s  ok", value, result)
	} else {
		log.Println("error")
	}
}

func TestDelKeyFromRedis(t *testing.T) {
	key := "settest"
	result, err := DelKeyFromRedis(key)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}

func TestSetNXWithExpireToRedis(t *testing.T) {
	key := "testexpire"
	result, err := SetNXWithExpireToRedis(key, "testvalue", EXPIRE)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}

func TestSetNXToRedis(t *testing.T) {
	key := "testsetnx"
	result, err := SetNXToRedis(key, "nxvalue")
	if err != nil {
		panic(err)
	}
	log.Println(result)
}

func TestIncr(t *testing.T) {
	log.Println("test111")
	key := "test_incr"
	err := Incr(key)
	if err != nil {
		panic(err)
	}
}

func TestSetex(t *testing.T) {
	key := "test_setex"
	err := Setex(key, 60, "aaaaaaaaaa")
	if err != nil {
		panic(err)
	}
}

func TestHsetStringToRedis(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	cacheStr, _ := json.Marshal(Test{Name: "test", Age: 12})
	err := SetStringToClassRoomRedis(common.USER_COURSE_DATA+fmt.Sprint(1), string(cacheStr))
	fmt.Println("####reply, err: ", err)
}

func TestHgetStringToRedis(t *testing.T) {
	reply, err := GetStringFromClassRoomRedis(common.USER_COURSE_DATA + "1")
	fmt.Println("reply: ", reply, err)
}
