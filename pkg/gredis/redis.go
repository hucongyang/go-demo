package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"

	"github.com/hucongyang/go-demo/conf"
)

// redis连接池
var RedisConn *redis.Pool

// 初始化设置redis
func Setup() error {
	config := conf.Config()
	RedisConn = &redis.Pool{
		MaxIdle:     config.Redis.MaxIdle,
		MaxActive:   config.Redis.MaxActive,
		IdleTimeout: config.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", config.Redis.Addr)
			if err != nil {
				return nil, err
			}
			if config.Redis.Password != "" {
				if _, err := conn.Do("AUTH", config.Redis.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
	return nil
}

// 根据key设置缓存
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

// 缓存key是否存在
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

// 根据key获取缓存
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 根据key删除缓存
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		return false, err
	}
	return true, nil
}

// 根据key like删除缓存
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
