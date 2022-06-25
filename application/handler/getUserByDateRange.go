package handler

import (
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func (h *Handler) GetUserByDateRange() gin.HandlerFunc {
	return func(context *gin.Context) {
		firstDate := context.Param("first")
		secondDate := context.Param("second")

		start, err := time.Parse("2006-01-02", firstDate)

		end, err := time.Parse("2006-01-02", secondDate)
		if err != nil {
			log.Fatal(err)
		}

		users, err := h.UserService.GetUserByDate(start, end)
		if err != nil {
			response.JSON(context, "", 500, nil, []string{err.Error()})
			return
		}
		response.JSON(context, "users retrieved successfully", 200, gin.H{
			"users": users,
		}, nil)
	}
}
