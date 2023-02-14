package item

import (
	"context"
	"strconv"
)

type ItemUsecase interface{
	FetchData(c context.Context) (res []Item, err error)
}

type itemUsecase struct{
	itemRepository ItemRepository
}

func NewItemUsecase(itemRepository ItemRepository) ItemUsecase{
	return &itemUsecase{itemRepository}
}

func(uc *itemUsecase) FetchData(c context.Context) (res []Item, err error){
	items, err := uc.itemRepository.Fetch(c)
	lastIDR, err := uc.itemRepository.Currency(c)

	var item Item
	for _, itemData := range items{
		id, err := strconv.Atoi(itemData.ID)
		if err != nil {
			return []Item{}, err
		}
		item.ID = id
		item.CreatedAt = itemData.CreatedAt
		price, err := strconv.ParseFloat(itemData.Price, 32)
		if err != nil {
			return []Item{}, err
		}
		item.Price = float32(price)
		item.Department = itemData.Department
		item.Product = itemData.Product
		item.IDR = float32(price) * lastIDR
		res = append(res, item)
	}
	
	return res, err
}

