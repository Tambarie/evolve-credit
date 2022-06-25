package handler

import (
	"github.com/Tambarie/evolve-credit/domain/helpers"
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUerByEmail() gin.HandlerFunc {
	return func(context *gin.Context) {
		email := struct {
			Email string `json:"email"`
		}{}

		if err := helpers.Decode(context, &email); err != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"could not decode user"})
			return
		}

		user, err := h.UserService.GetUserDetailsByEmail(email.Email)
		if err != nil {
			response.JSON(context, "", http.StatusInternalServerError, nil, []string{"could not get user by email"})
		}

		response.JSON(context, "user retrieved successfully", http.StatusCreated, gin.H{
			"user": user,
		}, nil)
	}
}
