/**************************************************
*文件名：oneKey.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/19
**************************************************/

package common

import (
	"sync"
	"time"
)

type (
	uniqueKeyExt struct {
		data int64
		lk   sync.RWMutex
	}
)

var uniqueKey uniqueKeyExt

// GetUniqueKey
// Description: 获取唯一key 时间戳格式
// Author: CYL96
// Date: 2023-08-16 14:23:57
// Return int :
func GetUniqueKey() int64 {
	uniqueKey.lk.Lock()
	defer uniqueKey.lk.Unlock()
	now := time.Now().Unix()
	if now > uniqueKey.data {
		uniqueKey.data = now
	} else {
		uniqueKey.data++
	}
	return uniqueKey.data
}

// GetUniqueKeyStr
// Description: 获取唯一key string 20060102150405
// Author: CYL96
// Date: 2023-08-16 14:25:16
// Return int :
func GetUniqueKeyStr() string {
	uniqueKey.lk.Lock()
	defer uniqueKey.lk.Unlock()
	now := time.Now().Unix()
	if now > uniqueKey.data {
		uniqueKey.data = now
	} else {
		uniqueKey.data++
	}
	return time.Unix(uniqueKey.data, 0).Format("20060102150405")
}
