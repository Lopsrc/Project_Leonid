package mocks

import (
	context "context"
	"testing"
    
	"rest-api/m/rest-api/internal/user"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetUser(t *testing.T) {
	rep := NewRepository(t)
	usr := &user.GetUser{
		ID: 1,
    }
	rep.On("GetById",  context.Background(),usr).Return(user.User{}, nil).Once()
	_, err := rep.GetById(context.Background(), usr)
    assert.NoError(t, err)
}

func TestHandler_UpdateUser(t *testing.T) {
	rep := NewRepository(t)
	usr := &user.UpdateUser{
		Id: 1,
        Name: gofakeit.Name(),
        Sex: "male",
        Birthdate: pgtype.Date{
            Time:            gofakeit.Date(),
            Status:           0,
            InfinityModifier: 0,
        },
        Age:    gofakeit.Number(0, 100),
        Weight: gofakeit.Number(0, 100),
    }
	rep.On("Update",  context.Background(),usr).Return(nil).Once()
	err := rep.Update(context.Background(), usr)
    assert.NoError(t, err)
}
