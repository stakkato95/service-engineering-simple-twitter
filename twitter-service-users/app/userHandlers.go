package app

import (
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/handlers"
	"github.com/stakkato95/twitter-service-users/service"
)

type userHandlers struct {
	service service.UserService
}

func (h *userHandlers) hello(w http.ResponseWriter, r *http.Request) {
	handlers.WriteResponse(w, http.StatusOK, "hello")
}
