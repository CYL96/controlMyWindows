package msg

type MsgApi string

func GetMsgApiDefine() map[MsgApi]MsgPushDefine {
	def := make(map[MsgApi]MsgPushDefine)
	def[MsgApiHeartBeat] = NewMsgApiHeartBeatDef()
	return def
}

// 心跳
const MsgApiHeartBeat MsgApi = "heartBeat"

type MsgApiHeartBeatData struct {
}

func NewMsgApiHeartBeatDef() MsgPushDefine {
	def := MsgPushDefine{}
	def.MsgApi = MsgApiHeartBeat
	def.MsgDesc = "心跳接口，时间间隔可配置"
	def.Definition = NewMsgData(MsgApiHeartBeat, MsgApiHeartBeatData{})
	return def
}
func NewMsgApiHeartBeatMsg() MsgPushData {
	msg := NewMsgData(MsgApiHeartBeat, MsgApiHeartBeatData{})
	return msg
}

// 心跳
const MsgApiScriptStateChange MsgApi = "scriptStateChange"

type MsgApiScriptStateChangeData struct {
	ControlId int64 `json:"control_id"`
	DetailId  int64 `json:"detail_id"`
	State     int   `json:"state"`
}

func NewMsgApiScriptStateChangeDef() MsgPushDefine {
	def := MsgPushDefine{}
	def.MsgApi = MsgApiScriptStateChange
	def.MsgDesc = "运行状态变化"
	def.Definition = NewMsgData(MsgApiScriptStateChange, MsgApiScriptStateChangeData{})
	return def
}
func NewMsgApiScriptStateChangeMsg(controlId int64, detailId int64, running int) MsgPushData {
	msg := NewMsgData(MsgApiScriptStateChange, MsgApiScriptStateChangeData{
		ControlId: controlId,
		DetailId:  detailId,
		State:     running,
	})
	return msg
}
