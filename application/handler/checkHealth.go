package handler

import (
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckHealth() gin.HandlerFunc {
	return func(context *gin.Context) {
		response.JSON(context, "Evolve credit", 200, gin.H{
			"Home": "Evolve credit",
		}, nil)
	}
}
