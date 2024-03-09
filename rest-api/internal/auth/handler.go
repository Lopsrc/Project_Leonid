package auth

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/mail"
	"rest-api/m/rest-api/internal/apperror"
	"rest-api/m/rest-api/internal/handlers"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Repository
type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, user *User) (User, error)
	Update(ctx context.Context, user *UpdateUser) error
	Delete(ctx context.Context, user *DeleteUser) error
	Recover(ctx context.Context, user *RecoverUser) error
}

const (
	authCreateURL  = "/auth/reg"
	authGetURL     = "/auth/get"
	authUpdateURL  = "/auth/up"
	authDeleteURL  = "/auth/del"
	authRecoverURL = "/auth/rec"
)
type Handler struct {
	log     *slog.Logger
	repository Repository
}

func NewHandler(repository Repository, log *slog.Logger) handlers.Handler {
	return &Handler{
		repository: repository,
		log:     log,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost,   authCreateURL,  apperror.Middleware(h.RegisterUser))
	router.HandlerFunc(http.MethodGet, 	  authGetURL,     apperror.Middleware(h.GetUser))
	router.HandlerFunc(http.MethodPost,   authUpdateURL,  apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPost,   authDeleteURL,  apperror.Middleware(h.DeleteUser))
	router.HandlerFunc(http.MethodPost,   authRecoverURL, apperror.Middleware(h.RecoverUser))
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) error {

	h.log.Info("RegisterUser.")

	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		w.WriteHeader(400)
	}
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }

	passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(500)
	}

	err = h.repository.Create(context.TODO(), &User{
		Email: email,
        Passhash: passhash,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) error {

	h.log.Info("GetUser.")

	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
        w.WriteHeader(400)
    }
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }

	user, err := h.repository.GetByEmail(context.TODO(), &User{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err!= nil {
		w.WriteHeader(400)
        return err
    }

	allBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Info("UpdateUser.")

	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordnew := r.FormValue("passwordnew")
	if email == "" || password == "" || passwordnew == "" {
        w.WriteHeader(400)
    }
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }

	user, err := h.repository.GetByEmail(context.TODO(), &User{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	
	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err!= nil {
        w.WriteHeader(400)
        return err
    }
	
	passhash, err := bcrypt.GenerateFromPassword([]byte(passwordnew), bcrypt.DefaultCost)
	if err!= nil {
        w.WriteHeader(500)
    }
	
	err = h.repository.Update(context.TODO(), &UpdateUser{
		Id: user.ID,
		Passhash: passhash,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Info("DeleteUser.")

	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
        w.WriteHeader(400)
    }
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }
	
	user, err := h.repository.GetByEmail(context.TODO(), &User{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err!= nil {
        w.WriteHeader(400)
        return err
    }
	
	err = h.repository.Delete(context.TODO(), &DeleteUser{
		Id: user.ID,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handler) RecoverUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Info("RecoverUser.")
	
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
        w.WriteHeader(400)
    }
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }
	
	user, err := h.repository.GetByEmail(context.TODO(), &User{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	
	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err!= nil {
        w.WriteHeader(400)
        return err
    }
	
	err = h.repository.Recover(context.TODO(), &RecoverUser{
		Id: user.ID,
    })
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}