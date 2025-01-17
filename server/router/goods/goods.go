package goods

import "github.com/gin-gonic/gin"

type GoodRouter struct{}

func (g *GoodRouter) InitGoodsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("goods")
	{
		baseRouter.POST("createGoods", goodApi.CreateGoods)
		//	新增商品类目
		baseRouter.POST("createCategory", goodApi.CreateCategory)
		//	获取所有的商品类目
		baseRouter.GET("getAllCategory", goodApi.GetAllCategory)
		// 创建品牌
		baseRouter.POST("createBrand", goodApi.CreateBrand)
	}
	return baseRouter
}
