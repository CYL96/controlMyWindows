/***************************************************文件名：tray.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/2
**************************************************/

package systray

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/getlantern/systray"
	"server/src/common"
	"server/src/config"
	"server/src/runCtx"
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

	cfg := config.GetSystemConfig(runCtx.DefaultContext())
	var urlOpen = func(url string) {
		ch := systray.AddMenuItem(url, url)
		go func() {
			<-ch.ClickedCh
			common.OpenUrl(url)
		}()
	}

	if cfg.RunIp == "0.0.0.0" {
		ips := common.GetAllIp()
		for _, ip := range ips {
			url := fmt.Sprintf("http://%s:%d", ip, cfg.RunPort)
			urlOpen(url)
		}
	} else {
		url := fmt.Sprintf("http://%s:%d", cfg.RunIp, cfg.RunPort)
		urlOpen(url)
	}

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

}

func onExit() {
	// clean up here
}
