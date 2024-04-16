/**************************************************
*文件名：robotgo.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package module

import (
	"github.com/go-vgo/robotgo"
	"server/src/config"
)

type (
	RobotGo struct {
	}
)

func NewRobotGo(isScale bool) *RobotGo {
	robotgo.Scale = isScale
	CheckZero()
	return &RobotGo{}
}
func ReUpdateRobotGo(isScale bool) {
	robotgo.Scale = isScale
	CheckZero()
	return
}

var ZeroX, ZeroY = 0, 0

func CheckZero() {
	nowX, NowY := robotgo.Location()
	robotgo.Move(0, 0)
	ZeroX, ZeroY = robotgo.Location()

	robotgo.Move(nowX-ZeroX, NowY-ZeroY)
}

func (r *RobotGo) MouseScroll(num int, direction config.KMouseScrollDir) (err error) {
	dir := ""
	switch direction {
	case config.MouseScrollUp:
		dir = "up"
	case config.MouseScrollDown:
		dir = "down"
	case config.MouseScrollLeft:
		dir = "left"
	case config.MouseScrollRight:
		dir = "right"
	}
	robotgo.ScrollDir(num, dir)
	return
}

func (r *RobotGo) TouchNormal(key config.KKey) (err error) {
	return robotgo.KeyTap(key.Str())
}

func (r *RobotGo) TouchNormalDown(key config.KKey) (err error) {
	return robotgo.KeyDown(key.Str())

}

func (r *RobotGo) TouchNormalUp(key config.KKey) (err error) {
	return robotgo.KeyUp(key.Str())
}

func (r *RobotGo) TouchNormalDouble(key config.KKey) (err error) {
	err = robotgo.KeyTap(key.Str())
	if err != nil {
		return err
	}
	robotgo.Sleep(50)
	err = robotgo.KeyTap(key.Str())
	return

}

func (r *RobotGo) Input(str string) (err error) {
	robotgo.TypeStr(str)
	return nil
}

func (r *RobotGo) TouchCombinationKey(key []config.KKey) (err error) {
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

func (r *RobotGo) Sleep(ms int) error {
	robotgo.MilliSleep(ms)
	return nil
}

func (r *RobotGo) MouseClick(key config.KKey) (err error) {
	robotgo.Click(key.Str())
	return nil
}

func (r *RobotGo) MouseMove(x int, y int) (err error) {
	robotgo.MoveSmooth(x-ZeroX, y-ZeroY, 0.01, 0.01)
	return
}

func (r *RobotGo) MouseDoubleClick(key config.KKey) (err error) {
	robotgo.Click(key.Str(), true)
	return nil
}

func (r *RobotGo) MouseUp(key config.KKey) (err error) {
	return robotgo.MouseUp(key.Str())
}

func (r *RobotGo) MouseDown(key config.KKey) (err error) {
	return robotgo.MouseDown(key.Str())
}

func (r *RobotGo) DriverName() config.KKey {
	return "robotgo"
}

func (r *RobotGo) GetNowMousePosition() (x, y int) {
	return robotgo.Location()
}
