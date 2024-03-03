package model

type Attr struct {
	AttrID    int64        `json:"attrID"`
	AttrName  string       `json:"attrName"`
	ShowType  int          `json:"showType"` // 1 单选, 2 布尔, 3 多选
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
	Val       []*AttrValue `json:"val"`
}

type AttrValue struct {
	AttrValueID   int64  `json:"attrValueID"`
	AttrValueName string `json:"attrValueName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
