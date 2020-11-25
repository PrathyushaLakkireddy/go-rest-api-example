package server

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/modules/user"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	count := 0

	for _, route := range user.Routes {
		count++
		log.Println(count, route.Name, route.Method, route.Path)
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
