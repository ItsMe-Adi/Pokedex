package models

import "github.com/gin-gonic/gin"

type Response struct {
	Status  bool   `json:"status"`
	Data    gin.H  `json:"data"`
	Message string `json:"message"`
}
