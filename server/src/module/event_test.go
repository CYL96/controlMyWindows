/**************************************************
*文件名：event_test.go
*内容简述：*
*文件历史：
author CYL96 创建 2025/3/13
**************************************************/

package module

import (
	"fmt"
	"testing"

	hook "github.com/robotn/gohook"
)

func TestHook(t *testing.T) {
	hook.Register(hook.KeyDown, []string{"q", "lctrl", "lshift"}, func(e hook.Event) {
		fmt.Println(111)
	})
	start := hook.Start()
	<-hook.Process(start)

}
