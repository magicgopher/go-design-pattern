package goods

// Goods 表示生产的产品
type Goods struct {
	ID    int    // 产品ID
	Value string // 产品值
}

// NewGoods 创建一个新产品
func NewGoods(id int, value string) Goods {
	return Goods{ID: id, Value: value}
}
