package redisProxy

import (
	"github.com/hoisie/redis"
	"global"
	. "tools"
	"tools/cfg"
)

var client redis.Client

//初始化
func InitClient(ip string, port string, pwd string) error {
	INFO(global.ServerName + " Connect RedisServer ...")
	addr := ip + ":" + port
	client.Addr = addr
	client.Password = pwd
	client.Db = cfg.GetInt("server_id")
	err := client.Ping()
	if err == nil {
		INFO(global.ServerName + " Connect RedisServer Success")
	}
	return err
}
