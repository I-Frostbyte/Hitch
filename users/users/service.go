package users

import (
	"context"
	// "fmt"

	"github.com/I-Frostbyte/Hitch/protobufs/usersgrpc"
	"github.com/I-Frostbyte/Hitch/users/db/repo"
	"github.com/I-Frostbyte/Hitch/users/public/model"
	"github.com/rs/zerolog"
)

// Impl is the implementation of the UsersServiceServer interface.
type Impl struct {
	repo repo.UsersRepo
	logger zerolog.Logger
	config model.Config

	usersgrpc.UnsafeUserServiceServer
}

// NewUsersService returns a new instance of the UsersServiceServer implementation.
func NewUsersService(
	repo repo.UsersRepo,
	logger zerolog.Logger,
	config model.Config,
	) *Impl {
	return &Impl{
		repo: repo,
		logger: logger,
		config: config,
	}
}

func (u *Impl) GetUserByID(ctx context.Context, req *usersgrpc.GetUserByIDRequest) (*usersgrpc.GetUserByIDResponse, error) {
	panic("GetUserByID is not implemented yet")
}

func (u *Impl) CreateUser(ctx context.Context, req *usersgrpc.CreateUserRequest) (*usersgrpc.CreateUserResponse, error) {
	panic("CreateUser is not implemented yet")
}

func (u *Impl) UpdateUser(ctx context.Context, req *usersgrpc.UpdateUserRequest) (*usersgrpc.UpdateUserResponse, error) {
	panic("UpdateUser is not implemented yet")
}