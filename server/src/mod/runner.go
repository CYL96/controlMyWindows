/**************************************************
*文件名：runner.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/7
**************************************************/

package mod

import "sync"

type (
	runnerControlExt struct {
		cancelMap sync.Map
	}
	runnerExt struct {
		Key    string `json:"key"` //
		Cancel func()
	}
)

var runnerControl runnerControlExt

func AddControlRunner(key string, cancel func()) {
	data := runnerExt{
		Key:    key,
		Cancel: cancel,
	}
	runnerControl.cancelMap.Store(key, data)
}

func GetControlRunner(key string) (cancel func(), ok bool) {
	var value any
	value, ok = runnerControl.cancelMap.Load(key)
	if !ok {
		return
	}
	data, _ := value.(runnerExt)
	cancel = data.Cancel
	return
}

func DeleteControlRunner(key string) {
	runnerControl.cancelMap.Delete(key)
}
