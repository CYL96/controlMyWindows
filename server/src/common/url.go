/**************************************************
*文件名：url.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/6/21
**************************************************/

package common

import "os/exec"

func OpenUrl(url string) {
	if url == "" {
		return
	}
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}
