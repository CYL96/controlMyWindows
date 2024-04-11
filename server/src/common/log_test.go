/**************************************************
*文件名：log_test.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package common

import (
	"fmt"
	"runtime"
	"testing"
)

func TestInfo(t *testing.T) {
	pc, codePath, codeLine, ok := runtime.Caller(0)
	fmt.Println(pc, codePath, codeLine, ok)
	fdsafd3124321safdsa()
}

func fdsafd3124321safdsa() {
	fdsafdsafdsa()
}

func fdsafdsafdsa() {
	GetCaller()
}

func GetCaller() {
	for i := 1; true; i++ {
		_, path, line, ok := runtime.Caller(i)
		if !ok {
			return
		}
		fmt.Println(i, path, line, ok)
	}
	return
}
