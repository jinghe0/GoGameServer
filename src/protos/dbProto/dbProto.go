package dbProto

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/funny/binary"
)
import (
	"protos"
	//	. "tools"
)

//初始化消息ID和消息类型的对应关系
func init() {
	protos.SetMsg(ID_DB_User_LoginC2S, DB_User_LoginC2S{})
	protos.SetMsg(ID_DB_User_LoginS2C, DB_User_LoginS2C{})

	protos.SetMsg(ID_DB_User_UpdateLastLoginTimeC2S, DB_User_UpdateLastLoginTimeC2S{})
}

//是否是有效的消息ID
func IsValidID(msgID uint16) bool {
	return msgID >= 11000 && msgID <= 14999
}

//是否是有效的异步DB消息
func IsValidAsyncID(msgID uint16) bool {
	return msgID >= 12000 && msgID <= 14999
}

//是否是有效的同步DB消息
func IsValidSyncID(msgID uint16) bool {
	return msgID >= 11000 && msgID <= 11999
}

//序列化
func MarshalProtoMsg(identification uint64, args proto.Message) []byte {
	msgID := protos.GetMsgID(args)

	msgBody, _ := proto.Marshal(args)

	result := make([]byte, 2+8+len(msgBody))
	binary.PutUint16LE(result[:2], msgID)
	binary.PutUint64LE(result[2:10], identification)
	copy(result[10:], msgBody)

	return result
}

//反序列化消息
func UnmarshalProtoMsg(msg []byte) protos.ProtoMsg {
	if len(msg) < 10 {
		return protos.NullProtoMsg
	}

	msgID := binary.GetUint16LE(msg[:2])
	if !IsValidID(msgID) {
		return protos.NullProtoMsg
	}

	identification := binary.GetUint64LE(msg[2:10])

	msgBody := protos.GetMsgObject(msgID)
	if msgBody == nil {
		return protos.NullProtoMsg
	}

	err := proto.Unmarshal(msg[10:], msgBody)
	if err != nil {
		return protos.NullProtoMsg
	}

	return protos.ProtoMsg{
		ID:             msgID,
		Body:           msgBody,
		Identification: identification,
	}
}
