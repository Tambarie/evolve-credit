package handler

import (
	"github.com/Tambarie/evolve-credit/domain/helpers"
	"github.com/Tambarie/evolve-credit/domain/models"
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

// CreateWallet Handler to create models for user
func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {

		user := &models.User{}
		user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		user.ID = uuid.New().String()

		if errs := helpers.Decode(context, &user); errs != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"unable to decode json"})
			return
		}
		HashedPassword, err := helpers.GenerateHashPassword(user.Password)
		user.HashedPassword = string(HashedPassword)
		if err != nil {
			log.Fatalf("hash password err: %v\n", err)
			return
		}
		_, err = h.UserService.FindUserByEmail(user.Email)
		if err != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"user already exists"})
			return
		}

		_, err = h.UserService.CreateUser(user)
		if err != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"unable to create user"})
			return
		}
		response.JSON(context, "user created successfully", http.StatusCreated, nil, []string{})
	}
}
