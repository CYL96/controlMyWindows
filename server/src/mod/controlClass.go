/**************************************************
*文件名：controlClass.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package mod

import (
	"server/src/common"
	"server/src/config"
	"server/src/runCtx"
)

func AddControlClass(ctx *runCtx.RunCtx, para AddControlClassPara) (err error) {
	para.ControlId = common.GetUniqueKey()

	if para.KeyWidth == "" {
		para.KeyWidth = "80"
	}
	if para.KeyHeight == "" {
		para.KeyHeight = "80"
	}

	err = config.AddControl(config.ControlListExt{
		ControlListBase: para.ControlListBase,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

type (
	AddControlClassPara struct {
		config.ControlListBase
	}
)

func UpdateControlClass(ctx *runCtx.RunCtx, para UpdateControlClassPara) (err error) {
	if para.ControlId <= 0 {
		return ctx.ErrorNew("id is zero")
	}
	var info config.ControlListExt
	info, err = config.GetControlInfo(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}
	info.ControlName = para.ControlName
	info.KeyWidth = para.KeyWidth
	info.KeyHeight = para.KeyHeight
	info.MouseOffSetX = para.MouseOffSetX
	info.MouseOffSetY = para.MouseOffSetY
	err = config.UpdateControl(info)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

type (
	UpdateControlClassPara struct {
		config.ControlListBase
	}
)

func DeleteControlClass(ctx *runCtx.RunCtx, para DeleteControlClassPara) (err error) {
	err = config.DeleteControl(para.ControlListIdT)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

type (
	DeleteControlClassPara struct {
		config.ControlListIdT
	}
)

func UpdateControlClassOrder(ctx *runCtx.RunCtx, para UpdateControlClassOrderPara) (err error) {
	err = config.UpdateControlOrder(para.OrderList)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

type (
	UpdateControlClassOrderPara struct {
		OrderList []config.ControlListIdT `json:"order_list"`
	}
)

func GetControlClassList(ctx *runCtx.RunCtx, para GetControlClassListPara) (result GetControlClassListResult, err error) {
	result.List = config.GetControlList()
	return
}

type (
	GetControlClassListPara struct {
	}
	GetControlClassListResult struct {
		List []config.ControlListExt `json:"list"`
	}
)

func GetControlClassInfo(ctx *runCtx.RunCtx, para GetControlClassInfoPara) (result GetControlClassInfoResult, err error) {
	list := config.GetControlList()
	for _, ext := range list {
		if ext.ControlId == para.ControlId {
			result.ControlListBase = ext.ControlListBase
			return
		}
	}
	err = ctx.ErrorNew("control info not found")
	return
}

type (
	GetControlClassInfoPara struct {
		config.ControlListIdT
	}
	GetControlClassInfoResult struct {
		config.ControlListBase
	}
)
