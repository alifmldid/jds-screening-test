package server

import (
	"fetch-app/item"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	itemRepo := item.NewItemRepository()
	itemUsecase := item.NewItemUsecase(itemRepo)
	itemController := item.NewItemController(itemUsecase)

	registerItemRoutes(r, itemController)
}