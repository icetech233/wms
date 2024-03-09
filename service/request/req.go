package request

type AddSpuRequest struct {
	SpuCode string `json:"spuCode"`
	SpuDesc string `json:"spuDesc"`
}

type AddAttrRequest struct {
	AttrName string `json:"attrName"` // 属性名称
	ShowType int    `json:"showType"` // 1 单选, 2 布尔, 3 多选
}

type EditAttrRequest struct {
	AttrID   int64  `json:"attrID"`   // 属性名称
	AttrName string `json:"attrName"` // 属性名称
	ShowType int    `json:"showType"` // 1 单选, 2 布尔, 3 多选
}

type AddMenuRequest struct {
	// MenuID   int64  `json:"menuID"`
	ParentID int64  `json:"parentID"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Path     string `json:"path"`
}
