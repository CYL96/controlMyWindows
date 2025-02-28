/**************************************************
*文件名：script.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/11
**************************************************/

package module

import (
	. "server/src/config"
	"server/src/runCtx"
)

func ExecScript(ctx *runCtx.RunCtx, script ControlT) (err error) {
	if len(script.DetailKey) == 0 {
		return
	}
	var x, y = control.GetNowMousePosition()

	for _, key := range script.DetailKey {
		select {
		case <-ctx.Done():
			return
		default:

			if key.KeyType == KeyTypeMouseMove ||
				key.KeyType == KeyTypeMouseMoveStartingPoint ||
				key.KeyType == KeyTypeMouseMoveSmooth ||
				key.KeyType == KeyTypeMouseMoveSmoothStartingPoint {

				if key.KeyType == KeyTypeMouseMoveStartingPoint || key.KeyType == KeyTypeMouseMoveSmoothStartingPoint {
					key.PointX = x
					key.PointY = y
				}

				key.PointX += script.MouseOffSet.PointX
				key.PointY += script.MouseOffSet.PointY
			}

			ExecKey(ctx, key)
		}

	}
	return
}
