package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodApi struct{}

func (g *GoodApi) CreateGoods(c *gin.Context) {
	goodService.PrintHello()
}

// CreateCategory
// @Tags      Good
// @Summary   创建商品分类
// @Security   ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /goods/createCategory [post]
func (g *GoodApi) CreateCategory(c *gin.Context) {
	var r systemReq.CreateCategory
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.CreateCategoryVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category := &goods.GoodCategory{Name: r.Name}
	createReturn, err := goodService.CreateCategory(*category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("创建商品分类失败!", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("创建商品分类成功!", zap.String("name", r.Name))
	response.OkWithDetailed(createReturn, "创建成功", c)
}

// GetAllCategory
// @Tags      Good
// @Summary   获取所有商品分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]goods.GoodCategory,msg=string}  "获取所有商品分类成功"
// @Router    /goods/getAllCategory [get]
func (g *GoodApi) GetAllCategory(c *gin.Context) {
	// 从路由获取query参数 id
	id := c.Query("id")
	if id == "" {
		// 如果为空，则获取所有商品分类
		categories, err := goodService.GetAllCategory()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			global.GVA_LOG.Error("获取所有商品分类失败!", zap.Error(err))
			return
		}
		global.GVA_LOG.Info("获取所有商品分类成功!")
		response.OkWithDetailed(categories, "获取成功", c)
		return
	} else {
		uid, err := utils.StrToUint(id)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			global.GVA_LOG.Error("获取商品分类失败!", zap.Error(err))
			return
		}
		category, err := goodService.GetCategoryById(uid)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			global.GVA_LOG.Error("获取商品分类失败!", zap.Error(err))
			return
		}
		global.GVA_LOG.Info("获取商品分类成功!")
		response.OkWithDetailed(category, "获取成功", c)
	}
}
