package handler

import (
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h *Handler) GetAllUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := context.Query("limit")
		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			log.Println(err)
		}

		users, err := h.UserService.GetAllUsers(intLimit)
		if err != nil {
			response.JSON(context, "", 500, nil, []string{err.Error()})
			return
		}

		response.JSON(context, "retrieved users successfully", 200, gin.H{
			"users": users,
		}, nil)
	}
}
