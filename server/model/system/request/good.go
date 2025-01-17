package request

// CreateCategory 注册
type CreateCategory struct {
	Name string `json:"name" binding:"required"` // 分类名称，必填
}

// CreateGoods 注册
type CreateGoods struct {
	Name string `json:"name" binding:"required"`
}
