package main

import "time"

type StructA struct {
	DateTime string
	Age      int
}

type StructB struct {
	// ab    time.Duration
	DateTime time.Time
	Age      int
}
type Spu struct {
	SpuID     int64     `gorm:"primaryKey;column:spu_id;type:bigint;not null" json:"-"`
	SpuCode   string    `gorm:"column:spu_code;type:varchar(16)" json:"spuCode"`
	SpuDesc   string    `gorm:"column:spu_desc;type:varchar(32)" json:"spuDesc"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Spu) TableName() string {
	return "spu"
}

type Sku struct {
	SkuID     int64     `gorm:"primaryKey;column:sku_id;type:bigint;not null" json:"-"`
	SpuID     int64     `gorm:"column:spu_id;type:bigint" json:"spuId"`
	SkuCode   string    `gorm:"column:sku_code;type:varchar(16)" json:"skuCode"`
	SkuDesc   string    `gorm:"column:sku_desc;type:varchar(32)" json:"skuDesc"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Sku) TableName() string {
	return "sku"
}
