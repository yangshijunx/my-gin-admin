package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type SysHandelFileApi struct {
}

func (s *SysHandelFileApi) AnalyzeExcel(c *gin.Context) {
	file, err := c.FormFile("file") // "excel" 是前端表单中字段的名称
	if err != nil {
		response.FailWithMessage("文件上传失败", c)
		return
	}
	// 校验文件类型是否为 Excel 文件
	if !utils.IsExcelFile(file) {
		response.FailWithMessage("上传文件类型错误，请上传 Excel 文件", c)
		return
	}
	//	解析excel
	result, err := SysHandelFileService.ParseExcelFile(file)
	if err != nil {
		response.FailWithMessage("解析文件失败", c)
		return
	}
	response.OkWithData(result, c)
}
