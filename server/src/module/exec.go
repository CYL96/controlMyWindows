/**************************************************
*文件名：exec.go
*内容简述：*
*文件历史：
author CYL96 创建 2025/2/26
**************************************************/

package module

import (
	. "server/src/config"
	"server/src/runCtx"
)

func ExecKey(ctx *runCtx.RunCtx, key ControlKeyList) (err error) {
	log(ctx, key)
	switch key.KeyType {
	case KeyTypeDefault:
		// 单键
		switch key.KeyPress {
		case PressTypeClick:
			// 单击
			err = control.TouchNormal(key.Key)
		case PressTypeDoubleClick:
			// 双击
			err = control.TouchNormalDouble(key.Key)
		case PressTypePressDown:
			// 按下
			err = control.TouchNormalDown(key.Key)
		case PressTypePressUp:
			// 抬起
			err = control.TouchNormalUp(key.Key)
		}
	case KeyTypeText:
		// 文本
		err = control.Input(key.Input)
	case KeyTypeShortcutKey:
		// 快捷键
		keys := KeyListToKKeyList(key.KeyList)
		err = control.TouchCombinationKey(keys)
	case KeyTypeMouse:
		// 鼠标点击
		switch key.KeyPress {
		case PressTypeClick:
			// 单击
			err = control.MouseClick(key.Key)
		case PressTypeDoubleClick:
			// 双击
			err = control.MouseDoubleClick(key.Key)
		case PressTypePressDown:
			// 按下
			err = control.MouseDown(key.Key)
		case PressTypePressUp:
			// 抬起
			err = control.MouseUp(key.Key)
		}
	case KeyTypeMouseMove, KeyTypeMouseMoveStartingPoint:
		// 鼠标移动
		err = control.MouseMove(key.PointX, key.PointY)
	case KeyTypeMouseMoveSmooth, KeyTypeMouseMoveSmoothStartingPoint:
		// 鼠标移动
		err = control.MouseMoveSmooth(key.PointX, key.PointY)
	case KeyTypeMouseScroll:
		// 鼠标移动
		err = control.MouseScroll(key.Scroll, key.ScrollDir)

	case KeyTypeDelay:
		// 延迟
		delay := key.Delay
		interval := 10
		if delay > 100 {
			interval = 100
		}
		for delay > 0 {
			select {
			case <-ctx.Done():
				return
			default:
			}
			sleepTime := interval
			if delay <= interval {
				sleepTime = delay
			}
			err = control.Sleep(sleepTime)
			delay -= interval
		}
	}
	if err != nil {
		ctx.Error(err)
	}
	return
}

func log(ctx *runCtx.RunCtx, key ControlKeyList) {
	info := func(msg ...any) {
		if key.Remark != "" {
			msg = append(msg, "     ---------", key.Remark)
		}
		ctx.Info(msg...)
	}
	switch key.KeyType {
	case KeyTypeDefault:
		// 单键
		switch key.KeyPress {
		case PressTypeClick:
			// 单击
			info("单击:", key.Key)
		case PressTypeDoubleClick:
			// 双击
			info("双击:", key.Key)
		case PressTypePressDown:
			// 按下
			info("按下:", key.Key)
		case PressTypePressUp:
			// 抬起
			info("抬起:", key.Key)
		}
	case KeyTypeText:
		// 文本
		info("文本:", key.Input)
	case KeyTypeShortcutKey:
		// 快捷键
		keys := KeyListToKKeyList(key.KeyList)
		info("快捷键:", keys)
	case KeyTypeMouse:
		// 鼠标点击
		switch key.KeyPress {
		case PressTypeClick:
			// 单击
			info("鼠标单击:", key.Key)
		case PressTypeDoubleClick:
			// 双击
			info("鼠标双击:", key.Key)
		case PressTypePressDown:
			// 按下
			info("鼠标按下:", key.Key)
		case PressTypePressUp:
			// 抬起
			info("鼠标抬起:", key.Key)
		}
	case KeyTypeMouseMove, KeyTypeMouseMoveStartingPoint:
		// 鼠标移动
		info("鼠标定位:", key.PointX, key.PointY)
	case KeyTypeMouseMoveSmooth, KeyTypeMouseMoveSmoothStartingPoint:
		// 鼠标移动
		info("鼠标移动:", key.PointX, key.PointY)
	case KeyTypeMouseScroll:
		// 鼠标移动
		info("鼠标滚轮:", key.ScrollDir, key.Scroll)
	case KeyTypeDelay:
		// 延迟
		info("等待:", key.Delay, "ms")
	}
}
