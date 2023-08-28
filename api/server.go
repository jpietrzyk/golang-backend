package api

import (
	"fmt"
	db "jck/golangbackend/db/sqlc"
	"jck/golangbackend/token"
	"jck/golangbackend/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/healthcheck", server.healthcheck)

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/refresh", server.renewAccessToken)

	// TODO: Use this group for routes which needs to be authenticated
	// TODO: ex authRoutes.POST("/places, srerver.places)
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
