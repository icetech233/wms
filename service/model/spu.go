package model

type Spu struct {
	//SpuID     int64  `json:"-"`
	SpuCode   string `json:"spuCode"`
	SpuDesc   string `json:"spuDesc"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
