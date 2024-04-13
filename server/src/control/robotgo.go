/**************************************************
*文件名：robotgo.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package control

import (
	"github.com/go-vgo/robotgo"
)

type (
	robotGo struct {
	}
)

func (r *robotGo) MouseScroll(num int, direction KMouseScrollDir) (err error) {
	dir := ""
	switch direction {
	case MouseScrollUp:
		dir = "up"
	case MouseScrollDown:
		dir = "down"
	case MouseScrollLeft:
		dir = "left"
	case MouseScrollRight:
		dir = "right"
	}
	robotgo.ScrollDir(num, dir)
	return
}

func (r *robotGo) TouchNormal(key KKey) (err error) {
	return robotgo.KeyTap(key.Str())
}

func (r *robotGo) TouchNormalDown(key KKey) (err error) {
	return robotgo.KeyDown(key.Str())

}

func (r *robotGo) TouchNormalUp(key KKey) (err error) {
	return robotgo.KeyUp(key.Str())
}

func (r *robotGo) TouchNormalDouble(key KKey) (err error) {
	err = robotgo.KeyTap(key.Str())
	if err != nil {
		return err
	}
	robotgo.Sleep(50)
	err = robotgo.KeyTap(key.Str())
	return

}

func (r *robotGo) Input(str string) (err error) {
	robotgo.TypeStr(str)
	return nil
}

func (r *robotGo) TouchCombinationKey(key []KKey) (err error) {
	if len(key) == 0 {
		return
	}
	var keyStr string
	var ketaddr []interface{}
	keyStr = key[len(key)-1].Str()

	for i := 0; i <= len(key)-2; i++ {
		ketaddr = append(ketaddr, key[i].Str())
	}
	return robotgo.KeyTap(keyStr, ketaddr...)
}

func (r *robotGo) Sleep(ms int) error {
	robotgo.MilliSleep(ms)
	return nil
}

func (r *robotGo) MouseClick(key KKey) (err error) {
	robotgo.Click(key.Str())
	return nil
}

func (r *robotGo) MouseMove(x int, y int) (err error) {
	robotgo.Move(0, 0)
	x_r, y_r := robotgo.Location()
	robotgo.Move(x-x_r, y-y_r)
	return
}

func (r *robotGo) MouseDoubleClick(key KKey) (err error) {
	robotgo.Click(key.Str(), true)
	return nil
}

func (r *robotGo) MouseUp(key KKey) (err error) {
	return robotgo.MouseUp(key.Str())
}

func (r *robotGo) MouseDown(key KKey) (err error) {
	return robotgo.MouseDown(key.Str())
}

func (r *robotGo) DriverName() KKey {
	return "robotgo"
}

func (r *robotGo) GetNowMousePosition() (x, y int) {
	return robotgo.Location()
}
