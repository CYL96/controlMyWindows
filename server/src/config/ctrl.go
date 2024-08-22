/**************************************************
*文件名：ctrl.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package config

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"server/src/runCtx"
)

type (
	RunState int
)

const (
	RunStateFree    RunState = 1 // 空闲
	RunStateRunning RunState = 2 // 运行中
)

type (
	ShowType int
)

const (
	ShowTypeColor ShowType = 1
	ShowTypePic   ShowType = 2
)

type (
	ControlListExt struct {
		ControlListBase
		DetailList []ControlDetailExt `json:"detail_list"`
	}
	ControlListIdT struct {
		ControlId int64 `json:"control_id" default:"0" example:"0"` //
	}
	ControlListBase struct {
		ControlListIdT
		ControlName  string `json:"control_name" default:"" example:""`    //
		ControlOrder int    `json:"control_order" default:"0" example:"0"` //

		KeyWidth  string `json:"key_width"`  //
		KeyHeight string `json:"key_height"` //
	}
	ControlDetailExt struct {
		ControlDetailIdExt
		RunState       RunState `json:"run_state" default:"0" example:"0"`        // 1:空闲 2：运行中
		DetailName     string   `json:"detail_name" default:"" example:""`        //
		DetailShowName bool     `json:"detail_show_name"`                         // 是否显示名称
		DetailShowType ShowType `json:"detail_show_type" default:"0" example:"0"` // 1:color 2:图片
		DetailColor    string   `json:"detail_color" default:"" example:""`       //
		DetailPic      string   `json:"detail_pic" default:"" example:""`         // 图片
		ControlT
	}
	ControlDetailIdExt struct {
		DetailId int64 `json:"detail_id" default:"0" example:"0"` //
	}
	ControlT struct {
		ControlType       ControlType      `json:"control_type" default:"0" example:"0"` // 1:快捷键，2：脚本 3：打开文件夹目录  4:打开网页
		MouseBackToOrigin bool             `json:"mouse_back_to_origin"`                 // 在脚本中，如果鼠标移动，是否返回原点
		Path              string           `json:"path" default:"" example:""`           // 目录或网页
		DetailKey         []ControlKeyList `json:"detail_key"`
	}
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

	KeyTypeDefault KeyType = 1 // 单键
	KeyTypeText    KeyType = 2 // 文本

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

var controlList []ControlListExt
var controlLk sync.RWMutex

func ReadControlConfig(ctx *runCtx.RunCtx) (err error) {
	file, err := os.ReadFile("./config/control.json")
	if err != nil {
		ctx.Warn(err)
		err = nil
		return
	}
	err = json.Unmarshal(file, &controlList)
	if err != nil {
		return
	}
	for i := range controlList {
		for i2 := range controlList[i].DetailList {
			controlList[i].DetailList[i2].RunState = RunStateFree

			switch controlList[i].DetailList[i2].DetailShowType {
			default:
				controlList[i].DetailList[i2].DetailShowType = ShowTypeColor
			case ShowTypePic:

			}
		}
	}
	return
}

func KeyListToKKeyList(data []KeyListT) (list []KKey) {
	list = make([]KKey, 0, len(data))
	for _, key := range data {
		list = append(list, key.Key)
	}
	return
}

func SaveControlConfig222() (err error) {
	file, err := json.Marshal(controlList)
	if err != nil {
		return
	}
	err = os.WriteFile("./config/control.json", file, 0666)
	if err != nil {
		return
	}
	return
}

func SaveControlConfig(config []ControlListExt) (err error) {
	file, err := json.Marshal(config)
	if err != nil {
		return
	}
	err = os.WriteFile("./config/control.json", file, 0666)
	if err != nil {
		return
	}
	return
}

func AddControl(para ControlListExt) (err error) {
	controlLk.Lock()
	defer controlLk.Unlock()
	newList := make([]ControlListExt, len(controlList)+1)
	copy(newList, controlList)

	para.ControlOrder = len(newList)

	newList[len(newList)-1] = para

	err = SaveControlConfig(newList)
	if err != nil {
		return err
	}
	controlList = newList
	return
}
func UpdateControl(para ControlListExt) (err error) {
	controlLk.Lock()
	defer controlLk.Unlock()
	newList := make([]ControlListExt, len(controlList))
	copy(newList, controlList)
	for i := range newList {
		if newList[i].ControlId == para.ControlId {
			newList[i] = para
			break
		}
	}
	err = SaveControlConfig(newList)
	if err != nil {
		return err
	}
	controlList = newList
	return
}
func DeleteControl(id ControlListIdT) (err error) {
	controlLk.Lock()
	defer controlLk.Unlock()
	newList := make([]ControlListExt, 0, len(controlList)-1)
	for i := range controlList {
		if controlList[i].ControlId != id.ControlId {
			newList = append(newList, controlList[i])
			newList[len(newList)-1].ControlOrder = len(newList)
		}
	}
	err = SaveControlConfig(newList)
	if err != nil {
		return err
	}
	controlList = newList
	return
}
func GetControlInfo(id ControlListIdT) (result ControlListExt, err error) {
	controlLk.Lock()
	defer controlLk.Unlock()
	for _, info := range controlList {
		if info.ControlId == id.ControlId {
			return info, nil
		}
	}
	err = errors.New("control id not found")
	return
}

func SetControlDetailState(state RunState, id int64, detailId int64) {
	controlLk.Lock()
	defer controlLk.Unlock()
	for idx, _ := range controlList {
		if controlList[idx].ControlId == id {
			for i, _ := range controlList[idx].DetailList {
				if controlList[idx].DetailList[i].DetailId == detailId {
					controlList[idx].DetailList[i].RunState = state
					break
				}
			}
			break
		}

	}
}

func GetControlList() (list []ControlListExt) {
	controlLk.Lock()
	defer controlLk.Unlock()
	list = make([]ControlListExt, len(controlList))
	for i := range list {
		list[i].ControlListBase = controlList[i].ControlListBase
	}
	return
}
func UpdateControlOrder(para []ControlListIdT) (err error) {
	controlLk.Lock()
	defer controlLk.Unlock()
	data := make(map[int64]ControlListExt, len(controlList))
	for _, ext := range controlList {
		data[ext.ControlId] = ext
	}
	newList := make([]ControlListExt, 0, len(controlList)+1)

	for _, t := range para {
		if value, ok := data[t.ControlId]; ok {
			newList = append(newList, value)
			newList[len(newList)-1].ControlOrder = len(newList)
		} else {
			err = errors.New("control id not found")
		}
	}

	err = SaveControlConfig(newList)
	if err != nil {
		return err
	}
	controlList = newList
	return
}
