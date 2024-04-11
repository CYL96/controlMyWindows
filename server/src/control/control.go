/**************************************************
*文件名：control.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package control

import (
	"context"
	"errors"
	"os/exec"

	"golang.org/x/sys/execabs"
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

	KeyTypeDefault     KeyType = 1  // 单键
	KeyTypeText        KeyType = 2  // 文本
	KeyTypeShortcutKey KeyType = 3  // 快捷键
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
		KeyType  KeyType    `json:"key_type" default:"0" example:"0"`  // 1 ：单键 2 ：文本 3 ：快捷键 99：延迟
		KeyPress PressType  `json:"key_press" default:"0" example:"0"` // 当KeyType == 1时 1：单击 2：双击 3：按下 4：抬起
		Input    string     `json:"input" default:"" example:""`       // 当KeyType == 2时 输入文本
		KeyList  []KeyListT `json:"key_list"`                          // 当KeyType == 3时
		Delay    int        `json:"delay" default:"0" example:"0"`     // 当KeyType == 99时 使用 ms
	}
	KeyListT struct {
		Id  int    `json:"id" default:"0" example:"0"` //
		Key string `json:"key" default:"" example:""`  //
	}
)

type ControlIntf interface {
	TouchNormal(ctx context.Context, key []KeyListT) (err error)
	TouchScript(ctx context.Context, key []ControlKeyList) (err error)
}

var control ControlIntf

func TouchKey(ctx context.Context, script ControlT) (err error) {
	if control == nil {
		err = errors.New("control not init")
		return
	}

	switch script.ControlType {
	case ControlTypeNormal:
		// 快捷键
		keyList := make([]KeyListT, 0, len(script.DetailKey))
		for i := range script.DetailKey {
			keyList = append(keyList, script.DetailKey[i].KeyListT)
		}
		return control.TouchNormal(ctx, keyList)
	case ControlTypeScript:
		// 脚本
		return control.TouchScript(ctx, script.DetailKey)
	case ControlTypeExplorer:
		// 打开文件夹目录
		OpenExplorer(ctx, script.Path)
	case ControlTypeWebsite:
		// 打开网页
		OpenUrl(ctx, script.Path)
	}
	return
}

func init() {
	control = new(robotGo)
}

func OpenExplorer(ctx context.Context, path string) {
	exec.Command(`cmd`, `/c`, `explorer`, path).CombinedOutput()

}

func OpenUrl(ctx context.Context, url string) {
	cmd := execabs.Command("rundll32", "url.dll,FileProtocolHandler", url)
	cmd.CombinedOutput()
}
