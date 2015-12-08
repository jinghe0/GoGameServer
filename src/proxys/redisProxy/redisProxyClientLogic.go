package redisProxy

import (
//	."tools"
)

const (
	DB_Write_Msgs = "DB_Write_Msgs"
)

//增加DB的写操作
func PushDBWriteMsg(msg []byte) {
	client.Rpush(DB_Write_Msgs, msg)
}

//获取所有未处理的写操作
func PullDBWriteMsg() [][]byte{
	datas, err := client.Lrange(DB_Write_Msgs, 0, -1)
	if err != nil{
		return nil
	}

	client.Ltrim(DB_Write_Msgs, len(datas), -1)
	return datas
}