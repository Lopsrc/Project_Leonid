package userdata

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *UserData, id int) error
	FindOne(ctx context.Context, id int) (UserData, error)
	Update(ctx context.Context, user *UserData) error
	Delete(ctx context.Context, id string) error
}
