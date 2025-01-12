package goods

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	GoodApi
}

var (
	goodService = service.ServiceGroupApp.GoodsServiceGroup
)
