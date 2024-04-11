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

	"server/src/control"
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
		RunState    RunState `json:"run_state" default:"0" example:"0"`  // 1:空闲 2：运行中
		DetailName  string   `json:"detail_name" default:"" example:""`  //
		DetailColor string   `json:"detail_color" default:"" example:""` //
		control.ControlT
	}
	ControlDetailIdExt struct {
		DetailId int64 `json:"detail_id" default:"0" example:"0"` //
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
		}
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
