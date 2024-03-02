package entity

import "time"

type Menu struct {
	MenuID    int64     `gorm:"primaryKey;column:menu_id;type:bigint;not null"`
	ParentID  int64     `gorm:"column:parent_id;type:bigint"`
	Key       string    `gorm:"column:key;type:varchar(16)"`
	Name      string    `gorm:"column:name;type:varchar(16)"`
	Icon      string    `gorm:"column:icon;type:varchar(16)"`
	Path      string    `gorm:"column:path;type:varchar(32)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *Menu) TableName() string {
	return "menu"
}
