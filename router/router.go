package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/public-service/types"
	"github.com/sysu-go-online/static-service/controller"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
	})
	r := mux.NewRouter()

	// user collection
	r.Handle("/{username}/{filepath:.*}", types.ErrorHandler(controller.ImageFileHandler)).Methods("GET")

	// Use classic server and return it
	s := negroni.Classic()
	s.Use(c)
	s.UseHandler(r)
	return s
}
