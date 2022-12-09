package authdata

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, auth *AuthData) error
	FindOne(ctx context.Context, auth *AuthData) (bool)
	Update(ctx context.Context, user *AuthData) error
	Delete(ctx context.Context, id int) error
}
