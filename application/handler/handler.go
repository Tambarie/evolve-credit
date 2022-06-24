package handler

import (
	"github.com/Tambarie/evolve-credit/domain/service"
)

type Handler struct {
	UserService service.UserService
}
