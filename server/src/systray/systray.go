/***************************************************文件名：tray.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/2
**************************************************/

package systray

import (
	"encoding/base64"
	"log"

	"github.com/getlantern/systray"
)

func RunSystray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	ico, err := base64.StdEncoding.DecodeString(IconBase64)
	if err == nil {
		systray.SetIcon(ico)
	} else {
		log.Println(err)
	}
	systray.SetTitle("Control My Windows")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	// clean up here
}
