package request

type AddSpuRequest struct {
	SpuCode string `json:"spuCode"`
	SpuDesc string `json:"spuDesc"`
}

type AddAttrRequest struct {
	AttrName string `json:"attrName"` // 属性名称
	ShowType int    `json:"showType"` // 1 单选, 2 布尔, 3 多选
}
