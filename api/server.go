package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/w-h-a/simplebank/db/sqlc"
	"github.com/w-h-a/simplebank/token"
	"github.com/w-h-a/simplebank/util"
)

type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST(
		"/users",
		s.createUser,
	)

	router.POST(
		"/users/login",
		s.loginUser,
	)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.POST(
		"/accounts",
		s.createAccount,
	)

	authRoutes.GET(
		"/accounts/:id",
		s.getAccount,
	)

	authRoutes.GET(
		"/accounts",
		s.listAccounts,
	)

	authRoutes.POST(
		"/transfers",
		s.createTransfer,
	)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
