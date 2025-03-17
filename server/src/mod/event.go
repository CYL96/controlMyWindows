/**************************************************
*文件名：event.go
*内容简述：*
*文件历史：
author CYL96 创建 2025/3/13
**************************************************/

package mod

import (
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
}

var HookCenter = HookCenterT{}

func InitHookCenter(ctx *runCtx.RunCtx) {
	HookCenter.ctx = ctx
}

func RestartHookCenter() {
	HookCenter.StopHook()
	go InitHookCenter(HookCenter.ctx)
}

func (t *HookCenterT) StopHook() {
	t.lk.Lock()
	defer t.lk.Unlock()
	t.stopHook()
}
func (t *HookCenterT) IsHook() bool {
	t.lk.Lock()
	defer t.lk.Unlock()
	return t.isHook
}

func (t *HookCenterT) stopHook() {
	t.ctx.Info("停止Hook")
	if !t.isHook {
		return
	}
	t.isHook = false
	hook.End()
	count := 100
	for {
		time.Sleep(100 * time.Millisecond)
		if !t.isHook {
			break
		}
		count--
		if count <= 0 {
			break
		}
	}
	t.ctx.Info("停止Hook结束")
}

func (t *HookCenterT) StartHook(id config.ControlListIdT) {
	t.lk.Lock()
	defer t.lk.Unlock()
	if t.isHook {
		if t.id.ControlId == id.ControlId {
			return
		} else {
			t.stopHook()
		}
	}
	t.id = id
	t.isHook = true
	go func() {
		t.hook()
	}()
}

func (t *HookCenterT) hook() {
	defer func() {
		t.lk.Lock()
		defer t.lk.Unlock()
		t.isHook = false
	}()
	t.ctx.Info("注册后台监控服务")
	result, err := config.GetControlInfo(t.id)
	if err != nil {
		t.ctx.Error(err)
		return
	}
	for i := range result.DetailList {
		ext := result.DetailList[i]
		if ext.ControlType != config.ControlTypeScript {
			continue
		}
		if len(ext.CombinationKey) > 0 {
			keys := make([]string, len(ext.CombinationKey))
			for i := range keys {
				kk := ext.CombinationKey[i].Key.Str()

				keys[i] = kk
			}
			match := true
			for _, key := range keys {
				if _, ok := hook.Keycode[key]; !ok {
					t.ctx.Error("注册 失败 未识别的key：", key)
					match = false
				}
			}
			if match {
				t.ctx.Info("注册：", keys, "  执行脚本:", ext.DetailName)
				// id := common.GetUniqueKey()
				hook.Register(hook.KeyDown, keys, func(e hook.Event) {
					ctx := runCtx.DefaultContext()
					ctx.Info("检测到：", keys, "  执行脚本:", ext.DetailName)
					err2 := ExecControlDetail(ctx, ExecControlDetailPara{
						ControlListIdT:     t.id,
						ControlDetailIdExt: ext.ControlDetailIdExt,
					})
					if err2 != nil {
						t.ctx.Error(err2)
					}
				})
			}
		}
	}
	start := hook.Start()
	<-hook.Process(start)
}
