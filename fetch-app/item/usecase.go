package item

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/golang-jwt/jwt"
)

type ItemUsecase interface{
	FetchData(c context.Context) (res []Item, err error)
	AggregateData(c context.Context) (res []ItemAgg, err error)
	UserData(c context.Context) (user User)
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

func(uc *itemUsecase) AggregateData(c context.Context) (res []ItemAgg, err error){
	items, err := uc.itemRepository.Fetch(c)
	lastIDR, err := uc.itemRepository.Currency(c)

	var item ItemAgg
	for _, itemData := range items{
		item.Department = itemData.Department
		item.Product = itemData.Product

		price, err := strconv.ParseFloat(itemData.Price, 32)
		if err != nil {
			return []ItemAgg{}, err
		}
		item.Price = float32(price) * lastIDR

		res = append(res, item)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Price < res[j].Price
	})
	
	return res, err
}

func (uc *itemUsecase) UserData(c context.Context) (user User){
	userInfo := c.Value("userInfo").(jwt.MapClaims)

	fmt.Println(userInfo["id"].(float64))
	user.ID = int(userInfo["id"].(float64))
	user.Nik = userInfo["nik"].(string)
	user.Role = userInfo["role"].(string)

	return user
}