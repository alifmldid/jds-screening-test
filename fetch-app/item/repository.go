package item

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ItemRepository interface{
	Fetch(c context.Context) (items []ItemResponse, err error)
	Currency(c context.Context) (lastIDR float32, err error)
}

type itemRepository struct{}

func NewItemRepository() ItemRepository{
	return &itemRepository{}
}

func (repo *itemRepository) Fetch(c context.Context) (items []ItemResponse, err error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product", nil)
	if err != nil {
		return []ItemResponse{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return []ItemResponse{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&items)
	if err != nil {
		return []ItemResponse{}, err
	}

	return items, nil
}

func (repo *itemRepository) Currency(c context.Context) (lastIDR float32, err error){
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.freecurrencyapi.com/v1/latest?base_currency=USD&currencies=IDR", nil)
	request.Header.Set("apikey", "6b3vwey9yzl2Qlj1QzzA5Qyg91fmYJpuXugSKgjH")

	if err != nil {
		return 0, err
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	type idr struct{IDR float32 `json:"IDR"`}
	type currency struct{Data idr `json:"data"`}
	var lastCurrency currency

	fmt.Println(response.Body)
	err = json.NewDecoder(response.Body).Decode(&lastCurrency)
	if err != nil {
		return 0, err
	}

	lastIDR = lastCurrency.Data.IDR

	return lastIDR, nil
}