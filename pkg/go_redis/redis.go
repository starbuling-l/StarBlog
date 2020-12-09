package go_redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,     //最大空闲连接数
		MaxActive:   setting.RedisSetting.MaxActive,   //在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		IdleTimeout: setting.RedisSetting.IdleTimeout, //在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := conn.Do("AUTH", setting.RedisSetting.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //可选的应用程序检查健康功能
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get() //在连接池中获取一个活跃连接
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value) //向 Redis 服务器发送命令并返回收到的答复
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key)) //将命令返回转为布尔值
	if err != nil {
		return false
	}
	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DELETE", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		if _, err = Delete(key); err != nil {
			return err
		}
	}

	return nil
}
