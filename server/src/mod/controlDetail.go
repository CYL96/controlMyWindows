/**************************************************
*文件名：controlDetail.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package mod

import (
	"strconv"
	"sync"

	hook "github.com/robotn/gohook"
	"server/src/common"
	"server/src/config"
	"server/src/module"
	"server/src/runCtx"
)

func AddControlDetail(ctx *runCtx.RunCtx, para AddControlDetailPara) (err error) {
	for _, t := range para.Detail.CombinationKey {
		key := t.Key.Str()
		if _, ok := hook.Keycode[key]; !ok {
			err = ctx.ErrorNew("快捷键绑定失败，无法绑定：", key)
			return
		}
	}

	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}
	para.Detail.DetailId = common.GetUniqueKey()
	para.Detail.RunState = config.RunStateFree
	info.DetailList = append(info.DetailList, para.Detail)
	err = config.UpdateControl(info)
	if err != nil {
		ctx.Error(err)
		return
	}
	if para.Detail.ControlType == config.ControlTypeScript {
		go func() {
			HookCenter.StopHook()
			HookCenter.StartHook(para.ControlListIdT)
		}()
	}
	return
}

type (
	AddControlDetailPara struct {
		config.ControlListIdT
		Detail config.ControlDetailExt `json:"detail"`
	}
)

func UpdateControlDetail(ctx *runCtx.RunCtx, para UpdateControlDetailPara) (err error) {
	for _, t := range para.Detail.CombinationKey {
		key := t.Key.Str()
		if _, ok := hook.Keycode[key]; !ok {
			err = ctx.ErrorNew("快捷键绑定失败，无法绑定：", key)
			return
		}
	}

	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}
	find := false
	for i, _ := range info.DetailList {
		if info.DetailList[i].DetailId == para.Detail.DetailId {
			info.DetailList[i] = para.Detail
			find = true
			break
		}
	}
	if !find {
		return ctx.ErrorNew("detail not found")
	}
	err = config.UpdateControl(info)
	if err != nil {
		ctx.Error(err)
		return
	}
	if para.Detail.ControlType == config.ControlTypeScript {
		go func() {
			HookCenter.StopHook()
			HookCenter.StartHook(para.ControlListIdT)
		}()
	}
	return
}

type (
	UpdateControlDetailPara struct {
		config.ControlListIdT
		Detail config.ControlDetailExt `json:"detail"`
	}
)

func DeleteControlDetail(ctx *runCtx.RunCtx, para DeleteControlDetailPara) (err error) {
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}
	newList := make([]config.ControlDetailExt, 0, len(info.DetailList)-1)
	for i, _ := range info.DetailList {
		if info.DetailList[i].DetailId != para.DetailId {
			newList = append(newList, info.DetailList[i])
		}
	}
	info.DetailList = newList

	err = config.UpdateControl(info)
	if err != nil {
		ctx.Error(err)
		return
	}
	go func() {
		HookCenter.StopHook()
		HookCenter.StartHook(para.ControlListIdT)
	}()
	return
}

type (
	DeleteControlDetailPara struct {
		config.ControlListIdT
		config.ControlDetailIdExt
	}
)

func UpdateControlDetailOrder(ctx *runCtx.RunCtx, para UpdateControlDetailOrderPara) (err error) {
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}

	newList := make([]config.ControlDetailExt, 0, len(info.DetailList))
	data := make(map[int64]config.ControlDetailExt, len(info.DetailList))

	for _, ext := range info.DetailList {
		data[ext.DetailId] = ext
	}

	if len(para.Detail) != len(data) {
		return ctx.ErrorNew("para is err")
	}

	for _, order := range para.Detail {
		if value, ok := data[order.DetailId]; ok {
			newList = append(newList, value)
		} else {
			return ctx.ErrorNew("detail id not found")
		}
	}
	info.DetailList = newList

	err = config.UpdateControl(info)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

type (
	UpdateControlDetailOrderPara struct {
		config.ControlListIdT
		Detail []config.ControlDetailIdExt `json:"detail"`
	}
)

func GetControlDetailList(ctx *runCtx.RunCtx, para GetControlDetailListPara) (result GetControlDetailListResult, err error) {
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}

	result.Detail = make([]config.ControlDetailExt, len(info.DetailList))
	go func() {
		HookCenter.Activate()
	}()
	for i := range info.DetailList {
		if info.DetailList[i].CombinationKey == nil {
			info.DetailList[i].CombinationKey = []config.KeyListT{}
		}
	}
	copy(result.Detail, info.DetailList)
	return
}

type (
	GetControlDetailListPara struct {
		config.ControlListIdT
	}
	GetControlDetailListResult struct {
		Detail []config.ControlDetailExt `json:"detail"`
	}
)

var execLk sync.RWMutex

func ExecControlDetail(ctx *runCtx.RunCtx, para ExecControlDetailPara) (err error) {

	if para.ControlId <= 0 || para.DetailId <= 0 {
		return ctx.ErrorNew("id is zero")
	}
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return err
	}

	for _, detail := range info.DetailList {
		if detail.DetailId == para.DetailId {
			execLk.Lock()
			defer execLk.Unlock()
			if detail.RunState == config.RunStateRunning {
				err = ctx.ErrorNew("control is running")
				return
			}
			cctx, cancel := runCtx.WithCancel(ctx)
			config.SetControlDetailState(config.RunStateRunning, para.ControlId, para.DetailId)
			AddControlRunner(strconv.FormatInt(detail.DetailId, 10), cancel)
			go func() {
				var err error
				defer func() {
					config.SetControlDetailState(config.RunStateFree, para.ControlId, para.DetailId)
					DeleteControlRunner(strconv.FormatInt(detail.DetailId, 10))
				}()
				detail.MouseOffSet.PointX = info.MouseOffSetX
				detail.MouseOffSet.PointY = info.MouseOffSetY
				err = module.TouchKey(cctx, detail.ControlT)
				if err != nil {
					ctx.Error(err)
				}
			}()
			break
		}
	}
	return
}

type (
	ExecControlDetailPara struct {
		config.ControlListIdT
		config.ControlDetailIdExt
	}
)

func StopControlDetail(ctx *runCtx.RunCtx, para StopControlDetailPara) (err error) {

	if para.ControlId <= 0 || para.DetailId <= 0 {
		return ctx.ErrorNew("id is zero")
	}
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return err
	}
	for _, detail := range info.DetailList {
		if detail.DetailId == para.DetailId {
			if detail.RunState == config.RunStateFree {
				return
			}
			config.SetControlDetailState(config.RunStateFree, para.ControlId, para.DetailId)
			key := strconv.FormatInt(detail.DetailId, 10)
			cancel, ok := GetControlRunner(key)
			if ok {
				cancel()
				DeleteControlRunner(key)
			}
			break
		}
	}
	return
}

type (
	StopControlDetailPara struct {
		config.ControlListIdT
		config.ControlDetailIdExt
	}
)

func GetNowMousePosition(ctx *runCtx.RunCtx, para GetNowMousePositionPara) (result GetNowMousePositionResult, err error) {
	result.X, result.Y = module.GetNowMousePosition()
	return
}

type (
	GetNowMousePositionPara struct {
	}
	GetNowMousePositionResult struct {
		X int `json:"x" default:"0" example:"0"` //
		Y int `json:"y" default:"0"`
	}
)
