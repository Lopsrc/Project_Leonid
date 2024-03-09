package mocks

import (
	context "context"
	"rest-api/m/rest-api/internal/auth"
	"runtime"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

const (
	passLenDefault = 10
)

func TestHandler_CreateUser(t *testing.T) {
	runtime.Gosched()
	// rep := Repository{}
	rep := NewRepository(t)
	user := &auth.User{
		Email: gofakeit.Email(),
        Passhash: []byte(randomPassword()),
    }
	rep.On("Create",  context.Background(),user).Return(nil).Once()
	// var err error
	err := rep.Create(context.Background(), user)
	assert.NoError(t, err)
}

func TestHandler_UpdateUser(t *testing.T) {
	runtime.Gosched()
	rep := NewRepository(t)
	user := &auth.UpdateUser{
		Id: 1,
        Passhash: []byte(randomPassword()),
    }
	rep.On("Update",  context.Background(),user).Return(nil).Once()
	err := rep.Update(context.Background(), user)
	assert.NoError(t, err)
}

func TestHandler_DeleteUser(t *testing.T) {
	runtime.Gosched()
	rep := NewRepository(t)
	user := &auth.DeleteUser{
		Id: 1,
    }
	rep.On("Delete",  context.Background(),user).Return(nil).Once()
	err := rep.Delete(context.Background(), user)
	assert.NoError(t, err)
}

func TestHandler_RecoveryUser(t *testing.T) {
	runtime.Gosched()
	rep := NewRepository(t)
	user := &auth.RecoverUser{
		Id: 1,
    }
	rep.On("Recover",  context.Background(),user).Return(nil).Once()
	err := rep.Recover(context.Background(), user)
	assert.NoError(t, err)
}


func TestHandler_GetUser(t *testing.T) {
	runtime.Gosched()
	rep := NewRepository(t)
	user := &auth.User{
		Email: gofakeit.Email(),
        Passhash: []byte(randomPassword()),
    }
	rep.On("GetByEmail",  context.Background(),user).Return(auth.User{}, nil).Once()
	_, err := rep.GetByEmail(context.Background(), user)
	assert.NoError(t, err)
}

func randomPassword() string {
	return gofakeit.Password(true, true, true, true, false, passLenDefault)
}