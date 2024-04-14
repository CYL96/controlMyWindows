/**************************************************
*文件名：set.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/12
**************************************************/

package mod

import (
	"server/src/config"
	"server/src/control"
	"server/src/runCtx"
)

func EditSystemConfig(ctx *runCtx.RunCtx, para EditSystemConfigPara) (err error) {
	err = config.UpdateSystemConfig(ctx, para.UpdateSystemConfigPara)
	if err != nil {
		ctx.Error(err)
		return
	}
	control.ReUpdateRobotGo(para.IsScale)
	return

}

type (
	EditSystemConfigPara struct {
		config.UpdateSystemConfigPara
	}
)

func GetNowSystemConfig(ctx *runCtx.RunCtx) (resp GetNowSystemConfigResp, err error) {
	resp.GetSystemConfigResult = config.GetSystemConfig(ctx)
	return
}

type (
	GetNowSystemConfigResp struct {
		config.GetSystemConfigResult
	}
)
