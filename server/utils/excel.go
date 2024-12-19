package utils

import (
	"mime/multipart"
	"path/filepath"
)

func IsExcelFile(file *multipart.FileHeader) bool {
	// 判断文件扩展名
	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		return false
	}

	// 如果你需要更严格的 MIME 类型检查，可以通过文件头信息来判断
	// 这里的 MIME 类型只是一个简单的示例，具体可以根据需求扩展
	return file.Header.Get("Content-Type") == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
}
