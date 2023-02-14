package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FetchDataController(c *gin.Context)
	AggregateDataController(c *gin.Context)
}

type itemController struct{
	itemUsecase ItemUsecase
}

func NewItemController(itemUsecase ItemUsecase) ItemController{
	return &itemController{itemUsecase}
}

func (controller *itemController) FetchDataController(c *gin.Context){
	items, err := controller.itemUsecase.FetchData(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": items})
}


func (controller *itemController) AggregateDataController(c *gin.Context){
	items, err := controller.itemUsecase.AggregateData(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": items})	
}