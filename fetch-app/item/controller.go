package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FetchDataController(c *gin.Context)
	AggregateDataController(c *gin.Context)
	VerifyController(c *gin.Context)
}

type itemController struct{
	itemUsecase ItemUsecase
}

func NewItemController(itemUsecase ItemUsecase) ItemController{
	return &itemController{itemUsecase}
}

// Fetch Data godoc
// @Summary fetch data endpoint
// @Schemes
// @Description Fetch data from https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product
// @Tags Fetch Data
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} item.FetchDataResponse
// @Router /fetch [get]
func (controller *itemController) FetchDataController(c *gin.Context){
	items, err := controller.itemUsecase.FetchData(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": items})
}

// Aggregate Data godoc
// @Summary fetch data endpoint
// @Schemes
// @Description Aggregate Data from https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product  ) aggregation by department, product, price IDR and order ascending by price
// @Tags Aggregate Data
// @Accept json
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} item.AggregateDataResponse
// @Router /aggregate [get]
func (controller *itemController) AggregateDataController(c *gin.Context){
	items, err := controller.itemUsecase.AggregateData(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": ""})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": items})	
}


// Token Verify godoc
// @Summary token verify endpoint
// @Schemes
// @Description Display token private claims data
// @Tags Verify Token
// @Accept json
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} item.VerifyTokenResponse
// @Router /verify [get]
func (controller *itemController) VerifyController(c *gin.Context){
	user := controller.itemUsecase.UserData(c)

	c.JSON(200, gin.H{
		"message": "success",
		"data": user,
	})
}