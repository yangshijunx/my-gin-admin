package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
)

type GoodService struct{}

var GoodsServiceApp = new(GoodService)

func (g *GoodService) PrintHello() {
	println("hello")
}

func (g *GoodService) CreateCategory(category goods.GoodCategory) (goods.GoodCategory, error) {
	println("create category")
	err := global.GVA_DB.Create(&category).Error
	return category, err
}
