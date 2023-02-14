package item

import "time"

type Item struct{
	ID int `json:"id" example:"1"`
	CreatedAt time.Time `json:"createdAt"  example:"2021-06-09T09:37:05.527Z"`
	Price float32 `json:"price" example:"218"`
	Department string `json:"department" example:"Outdoors"`
	Product string `json:"product" example:"Salad"`
	IDR float32 `json:"IDR" example:"3317529"`
}

type ItemResponse struct{
	ID string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Price string `json:"price"`
	Department string `json:"department"`
	Product string `json:"product"`
}

type ItemAgg struct{
	Department string `json:"department" example:"Games"`
	Product string `json:"product" example:"Computer"`
	Price float32 `json:"price" example:"60872.094"`
}

type User struct{
	ID int `json:"id" example:"2"`
	Nik string `json:"nik" example:"3574526883729838"`
	Role string `json:"role" example:"admin"`
}

type FetchDataResponse struct{
	Message string `json:"message" example:"success"`
	Data []Item `json:"data"`
}

type AggregateDataResponse struct{
	Message string `json:"message" example:"success"`
	Data []ItemAgg `json:"data"`
}

type VerifyTokenResponse struct{
	Message string `json:"message" example:"success"`
	Data []User `json:"data"`
}