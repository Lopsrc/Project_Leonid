package userdata

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *UserData) error
	FindOne(ctx context.Context, user *UserData) (bool, error)
	Update(ctx context.Context, user *UserData) error
	Delete(ctx context.Context, id int) error
}
