/**************************************************
*文件名：control.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package control

import (
	"errors"
	"os/exec"

	"golang.org/x/sys/execabs"
	"server/src/runCtx"
)

type (
	ControlType int
	KeyType     int
	PressType   int
)

const (
	ControlTypeNormal   ControlType = 1 // 1:快捷键
	ControlTypeScript   ControlType = 2 // 2：脚本
	ControlTypeExplorer ControlType = 3 // 3：打开文件夹目录
	ControlTypeWebsite  ControlType = 4 // 4:打开网页
	ControlTypeRunExe   ControlType = 5 // 4:打开exe
	ControlTypeRunCmd   ControlType = 6 // 4:打开cmd

	KeyTypeDefault     KeyType = 1  // 单键
	KeyTypeText        KeyType = 2  // 文本
	KeyTypeShortcutKey KeyType = 3  // 快捷键
	KeyTypeMouse       KeyType = 4  // 鼠标点击
	KeyTypeMouseMove   KeyType = 5  // 鼠标移动
	KeyTypeMouseScroll KeyType = 6  // 鼠标滚轮
	KeyTypeDelay       KeyType = 99 // 延迟

	PressTypeClick       PressType = 1 // 单击
	PressTypeDoubleClick PressType = 2 // 双击
	PressTypePressDown   PressType = 3 // 按下
	PressTypePressUp     PressType = 4 // 抬起
)

type (
	ControlT struct {
		ControlType ControlType      `json:"control_type" default:"0" example:"0"` // 1:快捷键，2：脚本 3：打开文件夹目录  4:打开网页
		Path        string           `json:"path" default:"" example:""`           // 目录或网页
		DetailKey   []ControlKeyList `json:"detail_key"`
	}
	ControlKeyList struct {
		KeyListT
		KeyType  KeyType    `json:"key_type" default:"0" example:"0"`  // 1 ：单键 2 ：文本 3 ：快捷键 4 :鼠标点击 5 :鼠标移动 6:鼠标滚轮 99：延迟
		KeyPress PressType  `json:"key_press" default:"0" example:"0"` // 当KeyType == 1时 1：单击 2：双击 3：按下 4：抬起
		Input    string     `json:"input" default:"" example:""`       // 当KeyType == 2时 输入文本
		KeyList  []KeyListT `json:"key_list"`                          // 当KeyType == 3时
		ControlKeyMouse
		Delay int `json:"delay" default:"0" example:"0"` // 当KeyType == 99时 使用 ms
	}
	ControlKeyMouse struct {
		PointX    int             `json:"point_x" default:"0" example:"0"` // 当KeyType == 5时 使用 X
		PointY    int             `json:"point_y" default:"0" example:"0"` // 当KeyType == 5时 使用 Y
		Scroll    int             `json:"scroll" default:"0" example:"0"`  //
		ScrollDir KMouseScrollDir `json:"scroll_dir" default:"0" example:"0"`
	}
	KeyListT struct {
		Id  int  `json:"id" default:"0" example:"0"` //
		Key KKey `json:"key" default:"" example:""`  //
	}
)

func KeyListToKKeyList(data []KeyListT) (list []KKey) {
	list = make([]KKey, 0, len(data))
	for _, key := range data {
		list = append(list, key.Key)
	}
	return
}

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
		return ExecScript(ctx, script.DetailKey)
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

func init() {
	control = new(robotGo)
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
