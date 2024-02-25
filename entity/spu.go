package entity

import "time"

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
