package msg

import (
	"sync"
	"time"

	"server/src/common"
	"server/src/runCtx"
)

type (
	MsgPushData struct {
		MsgID    string `json:"id"`   // 消息id
		MsgApi   MsgApi `json:"api"`  // 消息api
		TimeUnix int64  `json:"unix"` // 推送时间戳
		Time     string `json:"time"` // 时间
		Data     any    // 数据内容
	}
	MsgPushDefine struct {
		MsgApi     MsgApi      `json:"api"`        // api
		MsgDesc    string      `json:"desc"`       // api 描述
		Definition MsgPushData `json:"definition"` // 定义
	}
)

func NewMsgData(api MsgApi, data any) MsgPushData {
	now := time.Now()
	return MsgPushData{
		MsgApi:   api,
		Data:     data,
		MsgID:    common.GetUniqueKeyStr(),
		TimeUnix: now.Unix(),
		Time:     now.Format("2006-01-02 15:04:05"),
	}
}

type msgCenter struct {
	ctx      *runCtx.RunCtx
	pushChan chan MsgPushData // 推送数据的通道
	lk       sync.RWMutex
	targets  map[WsClientID]*deviceStatePushWsExt
}

type WsClientID int64
