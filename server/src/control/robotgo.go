/**************************************************
*文件名：robotgo.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/28
**************************************************/

package control

import (
	"context"

	"github.com/go-vgo/robotgo"
)

type (
	robotGo struct {
	}
)

func (r *robotGo) DriverName() string {
	return "robotgo"
}

func (r *robotGo) TouchNormal(ctx context.Context, key []KeyListT) (err error) {
	if len(key) == 0 {
		return
	}
	var keyStr string
	var ketaddr []interface{}
	keyStr = key[len(key)-1].Key

	for i := 0; i <= len(key)-2; i++ {
		ketaddr = append(ketaddr, key[i].Key)
	}
	return robotgo.KeyTap(keyStr, ketaddr...)
}

func (r *robotGo) TouchScript(ctx context.Context, key []ControlKeyList) (err error) {
	if len(key) == 0 {
		return
	}
	for i := range key {
		select {
		case <-ctx.Done():
			return
		default:
			switch key[i].KeyType {
			case KeyTypeDefault:
				// 单键
				switch key[i].KeyPress {
				case PressTypeClick:
					// 单击
					err = robotgo.KeyTap(key[i].Key)
				case PressTypeDoubleClick:
					// 双击
					err = robotgo.KeyDown(key[i].Key)
					err = robotgo.KeyUp(key[i].Key)
					robotgo.MilliSleep(50)
					err = robotgo.KeyDown(key[i].Key)
					err = robotgo.KeyUp(key[i].Key)
				case PressTypePressDown:
					// 按下
					err = robotgo.KeyDown(key[i].Key)
				case PressTypePressUp:
					// 抬起
					err = robotgo.KeyUp(key[i].Key)
				}
			case KeyTypeText:
				// 文本
				robotgo.TypeStr(key[i].Input)
			case KeyTypeShortcutKey:
				// 快捷键
				err = r.TouchNormal(ctx, key[i].KeyList)
			case KeyTypeDelay:
				// 延迟
				delay := key[i].Delay
				for {
					select {
					case <-ctx.Done():
						return
					default:
						break
					}
					var interval = 100
					if delay <= interval {
						robotgo.MilliSleep(delay)
						break
					}
					robotgo.MilliSleep(interval)
					delay -= interval
				}
			}
		}

	}
	return
}
