package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type SysHandelFileApi struct {
}

// AnalyzeExcel 解析上传的 Excel 文件
// @Tags SysHandelFileApi
// @Summary 解析上传的 Excel 文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce application/json
// @Param file formData file true "上传的 Excel 文件"
// @Success 200 {object} response.Response "{"success":true,"data":{},"msg":"解析成功"}"
// @Failure 400 {object} response.Response "{"success":false,"data":{},"msg":"文件上传失败/上传文件类型错误/解析文件失败"}"
// @Router /sysHandelFile/analyzeExcel [post]
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
