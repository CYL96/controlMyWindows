/**************************************************
*文件名：icon.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/15
**************************************************/

package module

import (
	"os"
	"path/filepath"
	"strings"

	"server/src/runCtx"
)

type (
	IconT struct {
		IconName string `json:"icon_name"`
		IconPath string `json:"icon_path"`
	}
)

func GetNowIcon(ctx *runCtx.RunCtx, path string) (list []IconT, err error) {
	var (
		dir []os.DirEntry
	)

	dir, err = os.ReadDir(path)
	if err != nil {
		ctx.Warn(err)
		err = nil
		return
	}
	for _, ent := range dir {
		if ent.IsDir() {
			data, err := GetNowIcon(ctx, path+"/"+ent.Name())
			if err != nil {
				ctx.Warn(err)
				err = nil
				continue
			}
			list = append(list, data...)
		} else {
			if isImageFile(ent.Name()) {
				list = append(list, IconT{
					IconName: ent.Name(),
					IconPath: path + "/" + ent.Name(),
				})
			}

		}
	}
	return
}
func isImageFile(filename string) bool {
	// 将文件名转换为小写，以便比较
	ext := strings.ToLower(filepath.Ext(filename))
	// 检查常见的图片文件扩展名
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg":
		return true
	}
	return false
}
