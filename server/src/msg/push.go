package msg

import (
	"sync"
	"time"

	// "QA/service/hardware"
	"github.com/gorilla/websocket"
	"server/src/common"
	"server/src/runCtx"
)

var (
	MsgCenter *msgCenter
)

type (
	deviceStatePushWsExt struct {
		ctx      *runCtx.RunCtx
		ClientID WsClientID
		ok       bool
		lk       sync.RWMutex
		conn     *websocket.Conn
	}
)

func (d *deviceStatePushWsExt) PushData(data interface{}) {
	defer func() {
		// 防止崩溃
		if r := recover(); r != nil {
			d.ctx.Error("recover:", r)
		}
	}()
	d.lk.Lock()
	defer d.lk.Unlock()
	if !d.ok {
		// 说明这个推送不可用了
		return
	}
	err := d.conn.WriteJSON(data)
	if err != nil {
		// 说明无法推送了,设置状态为不可用
		d.ok = false
		d.ctx.Error(err)
		return
	}
}

func InitDeviceStatePushService(ctx *runCtx.RunCtx) {
	MsgCenter = new(msgCenter)
	// 给够10个缓存，防止消息过多导致阻塞
	MsgCenter.ctx = ctx
	MsgCenter.pushChan = make(chan MsgPushData, 10)
	MsgCenter.targets = map[WsClientID]*deviceStatePushWsExt{}

	go MsgCenter.pushDataService()
}

func NewClientID() WsClientID {
	return WsClientID(common.GetUniqueKey())
}

func (c *msgCenter) pushDataService() {
	defer func() {
		if r := recover(); r != nil {
			c.ctx.Error("recover:", r)
		}
		time.Sleep(1 * time.Second)
	}()
	tk := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		// 当有外部数据进入的时候需要先处理
		case data := <-c.pushChan:
			// 推送数据
			c.PushData(data)
		case <-tk.C:
			// 定时去给前端同步一下数据
			if len(c.targets) == 0 {
				continue
			}
			// 推送数据
			c.PushHeartData()

		}
	}
}

func (c *msgCenter) PushHeartData() {
	data := NewMsgData(MsgApiHeartBeat, MsgApiHeartBeatData{})
	c.PushData(data)
}
func (c *msgCenter) PushData(data MsgPushData) {
	conns := c.GetAllDevicePusher()
	for _, conn := range conns {
		go conn.PushData(data)
	}
}

func (c *msgCenter) PushApiData(api MsgApi, data any) {
	pushData := NewMsgData(api, data)
	c.PushData(pushData)
}

func (c *msgCenter) AddDevicePusher(conn *websocket.Conn) {
	c.lk.Lock()
	defer c.lk.Unlock()

	id := NewClientID()
	target := &deviceStatePushWsExt{
		ClientID: id,
		ctx:      c.ctx,
		ok:       true,
		lk:       sync.RWMutex{},
		conn:     conn,
	}
	conn.SetCloseHandler(c.closeConnFunc(id))

	// 立即推送一下数据
	pushResult := NewMsgData(MsgApiHeartBeat, MsgApiHeartBeatData{})
	target.PushData(pushResult)

	c.targets[id] = target

}

func (c *msgCenter) DelAndCloseDevicePusher(clientId WsClientID) {
	c.lk.Lock()
	defer c.lk.Unlock()
	c.delAndCloseDevicePusher(clientId)
}

func (c *msgCenter) delAndCloseDevicePusher(clientId WsClientID) {
	defer func() {
		if r := recover(); r != nil {
			c.ctx.Error("recover:", r)
		}
	}()
	conn, ok := c.targets[clientId]
	if !ok {
		return
	}
	delete(c.targets, conn.ClientID)
	conn.lk.Lock()
	defer conn.lk.Unlock()
	conn.conn.Close()
}

func (c *msgCenter) GetAllDevicePusher() (conns []*deviceStatePushWsExt) {
	c.lk.Lock()
	defer c.lk.Unlock()

	conns = make([]*deviceStatePushWsExt, 0, len(c.targets))
	for _, conn := range c.targets {
		if !conn.ok {
			c.delAndCloseDevicePusher(conn.ClientID)
		} else {
			conns = append(conns, conn)
		}
	}
	return
}

func (c *msgCenter) closeConnFunc(clientId WsClientID) func(code int, text string) error {
	return func(code int, text string) error {
		c.ctx.Warn("ws close : ", code, text)
		c.DelAndCloseDevicePusher(clientId)
		return nil
	}
}
