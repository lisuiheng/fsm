package dao

import (
	"context"
	"github.com/lisuiheng/fsm"
)

// ServersStore stores connection information for a `User`
type UsersStore interface {
	// All returns all Users in the store
	All(context.Context) ([]fsm.User, error)
	Add(ctx context.Context, src *fsm.User) (err error)
	Get(ctx context.Context, id string) (*fsm.User, error)
	Update(ctx context.Context, src *fsm.User) error
	Delete(context.Context, *fsm.User) error
}
