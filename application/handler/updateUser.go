package handler

import (
	"github.com/Tambarie/evolve-credit/domain/helpers"
	"github.com/Tambarie/evolve-credit/domain/models"
	"github.com/Tambarie/evolve-credit/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		updateInfo := &models.UpdateUserInfo{
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		if err := helpers.Decode(context, &updateInfo); err != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"unable to decode json"})
			return
		}

		hsd, err := h.UserService.FindUserByID(updateInfo.ID)
		log.Println(err, hsd, "check")
		if err == nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"user id not valid"})
			return
		}

		_, err = h.UserService.UpdateUser(updateInfo.ID, updateInfo)
		if err != nil {
			response.JSON(context, "", http.StatusBadRequest, nil, []string{"unable to update user"})
			return
		}

		response.JSON(context, "", http.StatusCreated, nil, []string{})
	}

}
