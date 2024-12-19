package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysHandelFileRouter struct{}

func (s *SysHandelFileRouter) InitSysHandelFileRouter(Router *gin.RouterGroup) {
	sysHandelFileRouter := Router.Group("sysHandelFile").Use(middleware.OperationRecord())
	{
		// 解析excel
		sysHandelFileRouter.POST("analysisExcel", sysHandelFileApi.AnalyzeExcel)
	}
}
