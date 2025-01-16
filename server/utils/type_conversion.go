package utils

import (
	"strconv"
)

// StrToUint 将字符串转换为 uint
func StrToUint(s string) (uint, error) {
	// 将字符串转换为 uint64
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	// 将 uint64 转换为 uint
	return uint(u), nil
}
