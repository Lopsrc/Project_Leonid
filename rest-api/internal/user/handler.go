package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"
	"rest-api/m/rest-api/internal/apperror"
	"strconv"
	"time"

	"rest-api/m/rest-api/internal/auth"
	"rest-api/m/rest-api/internal/handlers"
	"github.com/jackc/pgtype"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Repository
type Repository interface {
	GetById(ctx context.Context, user *GetUser) (User, error)
	GetByEmail(ctx context.Context, email string) (auth.User, error)  // FIXME: Delete this method. Implement authentication using tokens.
	Update(ctx context.Context, user *UpdateUser) error
}

const (
	userGetURL    = "/user/get"
	userUpdateURL = "/user/up"
)

type handler struct {
	log        *slog.Logger
	repository Repository
}

func NewHandler(repository Repository, log *slog.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		log:        log,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, userGetURL, apperror.Middleware(h.GetUser))
	router.HandlerFunc(http.MethodPost, userUpdateURL, apperror.Middleware(h.UpdateUser))
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Info("GetUser.")
	// Get a parameters from the request.
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		w.WriteHeader(400)
	}
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }
	// Get the user from the database.
	user, err := h.repository.GetByEmail(context.TODO(), email)
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	// Compare the password.
	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err != nil {
		w.WriteHeader(400)
		return err
	}
	// Get the user data from the database.
	all, err := h.repository.GetById(context.TODO(), &GetUser{
		ID: 1,
	})
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Info("UpdateUser.")
	
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		w.WriteHeader(400)
	}
	name := r.FormValue("name")
	sex := r.FormValue("sex")
	birthdate := r.FormValue("birthdate")
	strage := r.FormValue("age")
	strweight := r.FormValue("weight")

	if name == "" || sex == "" || birthdate == "" || strage == "" || strweight == "" || strage == "0" || strweight == "0" {
		w.WriteHeader(400)
	}
	if  _, err := mail.ParseAddress(email); err!= nil {
        w.WriteHeader(400)
    }
	
	year, err := strconv.Atoi(birthdate[0:4])
	if err!= nil {
        w.WriteHeader(400)
        return err
    }
	month, err := strconv.Atoi(birthdate[5:7])
	if err!= nil {
        w.WriteHeader(400)
        return err
    }
	day, err := strconv.Atoi(birthdate[8:10])
	if err!= nil {
        w.WriteHeader(400)
        return err
    }
	fmt.Printf("year: %d, month: %d, day: %d \n", year, month, day)
	age, err := strconv.Atoi(strage)
	if err != nil {
		w.WriteHeader(500)
		return err
	}
	weight, err := strconv.Atoi(strweight)
	if err != nil {
		w.WriteHeader(500)
		return err
	}

	user, err := h.repository.GetByEmail(context.TODO(), email)
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	
	if err = bcrypt.CompareHashAndPassword(user.Passhash, []byte(password)); err != nil {
		w.WriteHeader(400)
		return err
	}

	var dt time.Time
	err = h.repository.Update(context.TODO(), &UpdateUser{
		Id: user.ID,
		Name: name,
		Sex:  sex,
		Birthdate: pgtype.Date{
			Time:            dt.AddDate(year-1, month-1, day-1),
			Status:           0,
			InfinityModifier: 0,
		},
		Age:    age,
		Weight: weight,
	})
	if err != nil {
		w.WriteHeader(400)
		h.log.Error("UpdateUser: %w", err)
		return err
	}

	allBytes, err := json.Marshal(true)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}
