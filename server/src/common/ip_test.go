/**************************************************
*文件名：ip_test.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/6/21
**************************************************/

package common

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestGetAllIp(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return
	}
	for _, inte := range interfaces {
		fmt.Println("Name:", inte.Name)
		fmt.Println("HardwareAddr:", inte.HardwareAddr)
		fmt.Println("flag:", inte.Flags)
		addrs, err := inte.Addrs()
		fmt.Println("ip:", addrs, err)
		fmt.Println("-----------------------")
	}
}
