/**************************************************
*文件名：ip.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/6/21
**************************************************/

package common

import "net"

func GetAllIp() []string {
	var ipList []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipList
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipList = append(ipList, ipnet.IP.String())
			}
		}
	}
	return ipList
}
