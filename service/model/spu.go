package model

type Spu struct {
	SpuID     int64  `json:"spuID"`
	SpuCode   string `json:"spuCode"`
	SpuDesc   string `json:"spuDesc"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
