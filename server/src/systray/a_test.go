/**************************************************
*文件名：a_test.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/2
**************************************************/

package systray

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func Test_hahah(t *testing.T) {
	file, err := os.ReadFile("./1.ico")
	if err != nil {
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(file))
}
