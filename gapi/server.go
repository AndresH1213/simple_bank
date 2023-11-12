package gapi

import (
	"fmt"

	db "github.com/AndresH1213/simple_bank/db/sqlc"
	"github.com/AndresH1213/simple_bank/pb"
	"github.com/AndresH1213/simple_bank/token"
	"github.com/AndresH1213/simple_bank/util"
)

// Server serves gRPC requests for our banking services
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
