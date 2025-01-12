package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 商品状态类型
type ProductStatus string

const (
	ProductStatusOnSale  ProductStatus = "on_sale"  // 上架
	ProductStatusOffSale ProductStatus = "off_sale" // 下架
)

// 商品规格类型
type GoodSpec struct {
	global.GVA_MODEL
	GoodID uint   `json:"good_id" gorm:"comment:商品ID"` // 外键，关联商品
	Name   string `json:"name" gorm:"comment:规格名称"`    // 规格名称（如颜色、尺寸）
	Value  string `json:"value" gorm:"comment:规格值"`    // 规格值
}

// 商品分类
type GoodCategory struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:分类名称"` // 分类名称
}

// 商品品牌
type GoodBrand struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:品牌名称"` // 品牌名称
}

// 商品库存
type GoodInventory struct {
	global.GVA_MODEL
	GoodID uint `json:"good_id" gorm:"comment:商品ID"` // 外键，关联商品
	Stock  int  `json:"stock" gorm:"comment:库存数量"`   // 库存数量
}

// 商品价格
type GoodPrice struct {
	global.GVA_MODEL
	GoodID        uint    `json:"good_id" gorm:"comment:商品ID"`      // 外键，关联商品
	OriginalPrice float64 `json:"original_price" gorm:"comment:原价"` // 原价
	SalePrice     float64 `json:"sale_price" gorm:"comment:促销价"`    // 促销价
}

// 商品图片
type GoodImage struct {
	global.GVA_MODEL
	GoodID uint   `json:"good_id" gorm:"comment:商品ID"` // 外键，关联商品
	URL    string `json:"url" gorm:"comment:图片URL"`    // 图片URL
}

// 商品标签
type GoodTag struct {
	global.GVA_MODEL
	GoodID uint   `json:"good_id" gorm:"comment:商品ID"` // 外键，关联商品
	Tag    string `json:"tag" gorm:"comment:标签"`       // 标签
}

// 商品模型
type GoodModel struct {
	global.GVA_MODEL
	Name        string        `json:"name" gorm:"comment:商品名称"`              // 商品名称
	Description string        `json:"description" gorm:"comment:商品描述"`       // 商品描述
	Status      ProductStatus `json:"status" gorm:"comment:商品状态"`            // 商品状态
	CategoryID  uint          `json:"category_id" gorm:"comment:商品分类ID"`     // 外键，关联分类
	BrandID     uint          `json:"brand_id" gorm:"comment:商品品牌ID"`        // 外键，关联品牌
	Category    GoodCategory  `json:"category" gorm:"foreignKey:CategoryID"` // 商品分类
	Brand       GoodBrand     `json:"brand" gorm:"foreignKey:BrandID"`       // 商品品牌
	Price       GoodPrice     `json:"price" gorm:"foreignKey:GoodID"`        // 商品价格
	Inventory   GoodInventory `json:"inventory" gorm:"foreignKey:GoodID"`    // 商品库存
	Images      []GoodImage   `json:"images" gorm:"foreignKey:GoodID"`       // 商品图片列表
	Tags        []GoodTag     `json:"tags" gorm:"foreignKey:GoodID"`         // 商品标签
	Specs       []GoodSpec    `json:"specs" gorm:"foreignKey:GoodID"`        // 商品规格
}

// 自定义表名
func (GoodModel) TableName() string {
	return "good"
}

func (GoodCategory) TableName() string {
	return "good_category"
}

func (GoodBrand) TableName() string {
	return "good_brand"
}

func (GoodPrice) TableName() string {
	return "good_price"
}

func (GoodInventory) TableName() string {
	return "good_inventory"
}

func (GoodImage) TableName() string {
	return "good_image"
}

func (GoodTag) TableName() string {
	return "good_tag"
}

func (GoodSpec) TableName() string {
	return "good_spec"
}
