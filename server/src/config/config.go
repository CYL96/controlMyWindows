/**************************************************
*文件名：config.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/27
**************************************************/

package config

import "server/src/runCtx"

func InitConfig(ctx *runCtx.RunCtx) (err error) {

	ctx.Info("读取系统配置")
	err = readSystemConfig(ctx)
	if err != nil {
		ctx.Error("读取系统配置 err:", err)
		return
	}
	ctx.Info("读取控制配置")
	err = ReadControlConfig(ctx)
	if err != nil {
		ctx.Error("读取系统配置 err:", err)
		return
	}
	return

}
