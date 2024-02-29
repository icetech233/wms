package entity

import "time"

type Attr struct {
	AttrID    int64     `gorm:"primaryKey;column:attr_id;type:bigint;not null" json:"-"`
	AttrName  string    `gorm:"column:attr_name;type:varchar(128)" json:"attrName"` // 属性名称
	ShowType  int       `gorm:"column:show_type;type:int" json:"showType"`          // 1 单选, 2 布尔, 3 多选
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Attr) TableName() string {
	return "attr"
}

type SkuAttr struct {
	SkuAttrID   int64     `gorm:"primaryKey;column:sku_attr_id;type:bigint;not null" json:"-"`
	SkuID       int64     `gorm:"uniqueIndex:idx_s_a;column:sku_id;type:bigint;not null" json:"skuId"`
	SpuID       int64     `gorm:"uniqueIndex:idx_s_a;column:spu_id;type:bigint;not null" json:"spuId"`
	AttrID      int64     `gorm:"uniqueIndex:idx_s_a;column:attr_id;type:bigint;not null" json:"attrId"`
	AttrValueID int64     `gorm:"uniqueIndex:idx_s_a;column:attr_value_id;type:bigint;not null" json:"attrValueId"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *SkuAttr) TableName() string {
	return "sku_attr"
}
