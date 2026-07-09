package users

import (
	"context"
	// "fmt"

	"github.com/rs/zerolog"
	"github.com/I-Frostbyte/users/usersgrpc"
)

// Impl is the implementation of the UsersServiceServer interface.
type Impl struct {
	// repo repo.UsersRepo
	logger zerolog.Logger
	// config model.Config
}

func NewUsersService(logger zerolog.Logger) *Impl {
	return &Impl{
		logger: logger,
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

func (u *Impl) DeleteUser(ctx context.Context, req *usersgrpc.DeleteUserRequest) (*usersgrpc.DeleteUserResponse, error) {
	panic("DeleteUser is not implemented yet")
}