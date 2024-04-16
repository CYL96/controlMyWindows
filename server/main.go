/**************************************************
*文件名：main.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package main

import (
	"os"

	"server/src/config"
	"server/src/hd"
	"server/src/module"
	"server/src/runCtx"
	"server/src/systray"
)

func main() {
	ctx := runCtx.DefaultContext()
	ctx.Info("程序开始启动...")

	ctx.Info("读取配置...")
	err := config.InitConfig(ctx)

	ctx.Info("初始化控制器...")
	module.InitControl(ctx)

	if err != nil {
		return
	}
	ctx.Info("启动服务器配置...")
	go func() {
		err = hd.Serve(ctx)
		if err != nil {
			os.Exit(-1)
			return
		}
	}()

	systray.RunSystray()

}
