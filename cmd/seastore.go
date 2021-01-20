package main

import (
	"log"
	"net/http"

	"github.com/alvinfebriando/seastore/pkg/user/delivery/rest"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	userController := rest.NewController()

	router.POST("/users/", userController.Register)
	router.GET("/users/:username", userController.FindUserByUsername)

	log.Fatal(http.ListenAndServe(":8080", router))
}
