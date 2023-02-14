package item

import "time"

type Item struct{
	ID int `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Price float32 `json:"price"`
	Department string `json:"department"`
	Product string `json:"product"`
	IDR float32 `json:"IDR"`
}

type ItemResponse struct{
	ID string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Price string `json:"price"`
	Department string `json:"department"`
	Product string `json:"product"`
}

type ItemAgg struct{
	Department string `json:"department"`
	Product string `json:"product"`
	Price float32 `json:"price"`
}
