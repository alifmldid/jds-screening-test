package server

import (
	"fetch-app/item"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func registerItemRoutes(r *gin.Engine, controller item.ItemController){
	item := r.Group("/")
	item.Use(authRequired)
	item.GET("/fetch", controller.FetchDataController)
}

func authRequired(c *gin.Context){
	authorizationHeader := c.Request.Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte("pAnGAqeB4KCJCK6E5ziXOyAPKnzouixINdEWEQa2w2reP3cJjpx1kyk1ohMrxkC1"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	c.Next()	
}