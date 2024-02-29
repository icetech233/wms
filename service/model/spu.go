package model

type Spu struct {
	SpuID     int64  `json:"spuID"`
	SpuCode   string `json:"spuCode"`
	SpuDesc   string `json:"spuDesc"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Attr struct {
	AttrID    int64  `json:"attrID"`
	AttrName  string `json:"attrName"` // 属性名称
	ShowType  int    `json:"showType"` // 1 单选, 2 布尔, 3 多选
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
