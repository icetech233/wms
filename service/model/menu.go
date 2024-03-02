package model

type Menu struct {
	MenuID    int64  `json:"menuID"`
	ParentID  int64  `json:"parentID"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
