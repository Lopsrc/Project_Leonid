package authdata

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, auth *AuthData) error
	FindAll(ctx context.Context) (u []AuthData, err error)
	FindOne(ctx context.Context, id string) (AuthData, error)
	Update(ctx context.Context, user AuthData) error
	Delete(ctx context.Context, id string) error
}
