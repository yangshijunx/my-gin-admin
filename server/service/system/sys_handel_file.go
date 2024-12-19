package system

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

type SysHandelFileService struct{}

// ParseExcelFile 解析 Excel 文件并返回结果
func (s *SysHandelFileService) ParseExcelFile(file *multipart.FileHeader) ([]map[string]string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer f.Close()

	// 使用 excelize 读取 Excel 文件
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return nil, fmt.Errorf("无法读取 Excel 文件: %v", err)
	}

	// 获取第一个工作表的名称
	sheetName := excel.GetSheetName(excel.GetActiveSheetIndex())

	// 获取所有行数据
	rows, err := excel.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("无法获取行数据: %v", err)
	}

	// 解析表头
	var headers []string
	for _, cell := range rows[0] {
		headers = append(headers, cell)
	}

	// 解析每一行数据
	var result []map[string]string
	for i, row := range rows {
		if i == 0 { // 跳过表头
			continue
		}

		rowData := make(map[string]string)
		for j, cell := range row {
			if j < len(headers) {
				rowData[headers[j]] = cell
			}
		}
		result = append(result, rowData)
	}

	return result, nil
}
