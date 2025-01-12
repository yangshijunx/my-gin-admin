package goods

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	GoodRouter
}

var (
	goodApi = api.ApiGroupApp.GoodsApiGroup
)
