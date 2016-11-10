package common

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var (
	RedisClient *redis.Pool
)

func InitRedis() (err error) {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := beego.AppConfig.String("redis.host")
	REDIS_DB, _ := beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	return
}

func SetRedis(key, value string) (err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err = rc.Do("SET", key, value)
	if err != nil {
		return
	}
	_, err = rc.Do("EXPIRE", key, 3600)
	return
}
func GetRedis(key string) (value string, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	reply, err := rc.Do("GET", key)

	if err != nil {
		value = ""
		return
	}
	if reply == nil {
		value = ""
		err = errors.New("当前KEY不存在！")
		return
	}

	value = string(reply.([]byte))
	_, err = rc.Do("EXPIRE", key, 3600)

	return
}
