package api

import (
	"fmt"

	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/token"
	"github.com/alquraini/pawsitively/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	action     db.Action
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, action db.Action) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		action:     action,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// authentication
	router.POST("/users/register", server.registerUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// user
	authRoutes.PUT("/users/update", server.updateUser)
	authRoutes.DELETE("/users/delete", server.deleteUser)

	// pet
	authRoutes.POST("/users/pets/create", server.createNewPet)
	authRoutes.GET("/users/pets", server.getPets)
	authRoutes.GET("/users/pets/:id", server.getPetById)
	authRoutes.PUT("/users/pets/update", server.updatePet)
	authRoutes.DELETE("/users/pets/delete", server.deletePet)

	// post
	authRoutes.POST("/posts/create", server.createPost)
	router.GET("/posts/:id", server.getPost)
	router.GET("/posts", server.listPosts)
	router.GET("/users/:user_id/posts", server.listPostsByUserID)
	authRoutes.PUT("/posts/update", server.updatePost)
	authRoutes.DELETE("/posts/delete", server.deletePost)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func successResponse(message string) gin.H {
	return gin.H{"success": message}
}
