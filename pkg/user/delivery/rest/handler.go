package rest

import (
	"encoding/json"
	"net/http"

	"github.com/alvinfebriando/seastore/pkg/domain"
	"github.com/alvinfebriando/seastore/pkg/user/repository"
	"github.com/alvinfebriando/seastore/pkg/user/service"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	s *service.Service
}

func NewController() *Controller {
	r := repository.NewSliceRepository()
	s := service.NewService(r)
	return &Controller{s: s}
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var body registerDTO
	json.NewDecoder(r.Body).Decode(&body)

	user := domain.User{Name: body.Name, Email: body.Email, Username: body.Username}
	c.s.Register(user)
	createdUser, err := c.s.FindByUsername(body.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (c *Controller) FindUserByUsername(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	username := p.ByName("username")

	user, err := c.s.FindByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		newError := domain.ErrorMessage{Message: err.Error()}
		json.NewEncoder(w).Encode(newError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
