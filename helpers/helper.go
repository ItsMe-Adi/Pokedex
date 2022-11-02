package helpers

import (
	"pokedex/models"

	"github.com/gin-gonic/gin"
)

func StandardResponse(status bool, data gin.H, message string) models.Response {
	response := models.Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
	return response
}

func GetOffset(page int, limit int) int {
	return (page - 1) * limit
}
