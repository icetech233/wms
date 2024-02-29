package entity

import "time"

// AttrValue [...]
type AttrValue struct {
	AttrValueID   int64     `gorm:"primaryKey;column:attr_value_id;type:bigint;not null" json:"-"`
	AttrValueName string    `gorm:"column:attr_value_name;type:varchar(128)" json:"attrValueName"`
	AttrID        int64     `gorm:"index:fk_attr_id;column:attr_id;type:bigint;not null" json:"attrId"`
	Attr          Attr      `gorm:"joinForeignKey:attr_id;foreignKey:attr_id" json:"attrList"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *AttrValue) TableName() string {
	return "attr_value"
}
