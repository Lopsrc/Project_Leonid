package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"rest-api/m/rest-api/internal/user"
	"rest-api/m/rest-api/internal/auth"
	
	"rest-api/m/rest-api/pkg/client/postgresql"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	log *slog.Logger
}

func NewRepository(client postgresql.Client, log *slog.Logger) user.Repository {
	return &repository{
		client: client,
		log: log,
	}
}

func (r *repository) GetById(ctx context.Context, user *user.GetUser) (usr user.User, err error){
	query := "SELECT name, sex, birthdate, age, weight  FROM customers WHERE customer_id = $1"

	if err := r.client.QueryRow(ctx, query, user.ID).Scan(&usr.Name, &usr.Sex, &usr.Birthdate, &usr.Age, &usr.Weight); err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return usr, newErr
        }
        return usr, err
    }
	
	return usr, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (auth.User, error){
    query := "SELECT id, pass_hash, del FROM auth WHERE email = $1"
	var result auth.User

	if err := r.client.QueryRow(ctx, query, email).Scan(&result.ID, &result.Passhash, &result.IsDeleted); err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return auth.User{}, newErr
        }
        return auth.User{}, err
    }
	
	return result, nil
}

func (r *repository) Update(ctx context.Context, user *user.UpdateUser) error {
	query := "UPDATE customers SET name = $2, sex = $3, birthdate = $4, age = $5, weight = $6 WHERE customer_id = $1"
	r.log.Info("update")
	
	_, err := r.client.Exec(ctx, query, user.Id, user.Name, user.Sex, user.Birthdate.Time, user.Age, user.Weight)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return newErr
        }
		r.log.Error("auth.Update: %s",err)
        return err
    }
	return nil
}

