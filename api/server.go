package api

import (
	db "github.com/christy712/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// creates new gttp server and router
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)

	server.router = router
	return server

}

// runs the http server
func (server *Server) Start(address string) error {

	return server.router.Run(address)
}

func errorResponse(err error) *gin.H {
	return &gin.H{"error": err.Error()}

}
