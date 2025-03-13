/**************************************************
*文件名：event.go
*内容简述：*
*文件历史：
author CYL96 创建 2025/3/13
**************************************************/

package mod

import (
	"fmt"
	"sync"
	"time"

	hook "github.com/robotn/gohook"
	"server/src/config"
	"server/src/runCtx"
)

type HookCenterT struct {
	ctx    *runCtx.RunCtx
	id     config.ControlListIdT
	isHook bool
	lk     sync.RWMutex
	now    int64
}

var HookCenter = HookCenterT{}

func InitHookCenter(ctx *runCtx.RunCtx) {
	HookCenter.ctx = ctx
	go HookCenter.monitor()
}

func (t *HookCenterT) StopHook() {
	t.lk.Lock()
	defer t.lk.Unlock()
	if !t.isHook {
		return
	}
	hook.End()
}

func (t *HookCenterT) Activate() {
	t.lk.Lock()
	defer t.lk.Unlock()
	if !t.isHook {
		return
	}
	t.now = time.Now().Unix()
}

func (t *HookCenterT) StartHook(id config.ControlListIdT) {
	t.lk.Lock()
	defer t.lk.Unlock()
	if t.isHook {
		if t.id.ControlId == id.ControlId {
			return
		} else {
			hook.End()
		}
	}
	t.id = id
	time.Sleep(3 * time.Second)
	t.isHook = true
	go func() {
		t.hook()
	}()
}

func (t *HookCenterT) monitor() {
	for {
		time.Sleep(10 * time.Second)
		now := time.Now().Unix()
		if !t.isHook {
			continue
		}
		if now-t.now > 30 {
			t.ctx.Warn("长时间未注册，停止")
			hook.End()
		}
	}
}

func (t *HookCenterT) hook() {
	defer func() {
		t.lk.Lock()
		defer t.lk.Unlock()
		t.isHook = false
	}()
	fmt.Println("注册后台监控服务")
	result, err := config.GetControlInfo(t.id)
	if err != nil {
		t.ctx.Error(err)
		return
	}
	for _, ext := range result.DetailList {
		if ext.ControlType != config.ControlTypeScript {
			continue
		}
		if len(ext.CombinationKey) > 0 {
			keys := make([]string, len(ext.CombinationKey))
			for i := range keys {
				keys[i] = ext.CombinationKey[i].Key.Str()
			}
			t.ctx.Info("注册：", keys, "  执行脚本:", ext.DetailName)
			hook.Register(hook.KeyDown, keys, func(e hook.Event) {
				t.ctx.Info("检测到：", keys, "  执行脚本:", ext.DetailName)
				err2 := ExecControlDetail(t.ctx, ExecControlDetailPara{
					ControlListIdT:     t.id,
					ControlDetailIdExt: ext.ControlDetailIdExt,
				})
				if err2 != nil {
					t.ctx.Error(err2)
				}
			})
		}
	}
	start := hook.Start()
	<-hook.Process(start)
}
