/**************************************************
*文件名：control.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package module

import (
	"errors"
	"os/exec"

	. "server/src/config"

	"golang.org/x/sys/execabs"
	"server/src/runCtx"
)

// left
// right
// wheelLeft

type ControlIntf interface {
	TouchNormal(key KKey) (err error)                           // 单键
	TouchNormalDown(key KKey) (err error)                       // 单键按下
	TouchNormalUp(key KKey) (err error)                         // 单键抬起
	TouchNormalDouble(key KKey) (err error)                     // 单键双击
	TouchCombinationKey(key []KKey) (err error)                 // 组合键
	Input(str string) (err error)                               // 文本输入
	Sleep(ms int) error                                         // 睡眠
	MouseClick(key KKey) (err error)                            // 鼠标点击
	MouseDoubleClick(key KKey) (err error)                      // 鼠标双击
	MouseUp(key KKey) (err error)                               // 鼠标抬起
	MouseDown(key KKey) (err error)                             // 鼠标按下
	MouseMove(x int, y int) (err error)                         // 鼠标移动
	MouseScroll(num int, direction KMouseScrollDir) (err error) // direction 1：向上 2：向下 3:向左 4：向右
	GetNowMousePosition() (x int, y int)
}

var control ControlIntf

func TouchKey(ctx *runCtx.RunCtx, script ControlT) (err error) {
	if control == nil {
		err = errors.New("control not init")
		return
	}

	switch script.ControlType {
	case ControlTypeNormal:
		// 快捷键
		keyList := make([]KKey, 0, len(script.DetailKey))
		for _, key := range script.DetailKey {
			keyList = append(keyList, key.Key)
		}
		return control.TouchCombinationKey(keyList)
	case ControlTypeScript:
		// 脚本
		return ExecScript(ctx, script.MouseBackToOrigin, script.DetailKey)
	case ControlTypeExplorer:
		// 打开文件夹目录
		OpenExplorer(ctx, script.Path)
	case ControlTypeWebsite:
		// 打开网页
		OpenUrl(ctx, script.Path)
	case ControlTypeRunExe:
		// 打开程序
		RunExe(ctx, script.Path)
	case ControlTypeRunCmd:
		// 执行命令
		RunCmd(ctx, script.Path)

	}
	return
}

func InitControl(ctx *runCtx.RunCtx) {
	control = NewRobotGo(GetSystemConfig(ctx).IsScale)
}

func OpenExplorer(ctx *runCtx.RunCtx, path string) {
	exec.Command(`cmd`, `/c`, `explorer`, path).CombinedOutput()

}

func OpenUrl(ctx *runCtx.RunCtx, url string) {
	cmd := execabs.Command("rundll32", "url.dll,FileProtocolHandler", url)
	cmd.CombinedOutput()
}

func RunExe(ctx *runCtx.RunCtx, path string) {
	exec.Command(`cmd`, `/c`, path).CombinedOutput()

}

func RunCmd(ctx *runCtx.RunCtx, path string) {
	exec.Command(`cmd`, `/c`, path).CombinedOutput()

}

func GetNowMousePosition() (x, y int) {
	return control.GetNowMousePosition()
}
