package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	
	"rest-api/m/rest-api/internal/auth"
	"rest-api/m/rest-api/pkg/client/postgresql"

	"github.com/jackc/pgconn"
)
type repository struct {
	client postgresql.Client
	log *slog.Logger
}

func NewRepository(client postgresql.Client, log *slog.Logger) auth.Repository {
	return &repository{
		client: client,
		log: log,
	}
}

func (r *repository) Create(ctx context.Context, user *auth.User) error {
	query := "INSERT INTO auth (email, pass_hash) VALUES ($1, $2) RETURNING id"

	tx, err := r.client.Begin(ctx)
	if err != nil {
		r.log.Error("user.Create: %s",err)
		return err
	}
	defer tx.Rollback(ctx)

	if err = tx.QueryRow(ctx, query, user.Email, user.Passhash).Scan(&user.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.log.Error("auth.Create: %s",newErr)
			return newErr
		}
		return err
	}

	query = "INSERT INTO customers (customer_id, name, sex, birthdate, age, weight) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = tx.Exec(ctx, query,  user.ID, "000", "male", "1991-01-01", 1, 1)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.log.Error("user.Create: %s",newErr)
			return newErr
		}
		return err
	}

	if err  = tx.Commit(ctx); err != nil {
		r.log.Error("user.Create: %s",err)
        return err
    }

	return nil
}

func (r *repository) Update(ctx context.Context, user *auth.UpdateUser) error {
	query := "UPDATE auth SET pass_hash = $1 WHERE id = $2"

	_, err := r.client.Exec(ctx, query, user.Passhash, user.Id)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return newErr
        }
        return err
    }
	
	return nil
}

func (r *repository) GetByEmail(ctx context.Context, user *auth.User) (auth.User, error){
	query := "SELECT id, pass_hash, del FROM auth WHERE email = $1"
	var result auth.User

	if err := r.client.QueryRow(ctx, query, user.Email).Scan(&result.ID, &result.Passhash, &result.IsDeleted); err!= nil {
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

func (r *repository) Delete(ctx context.Context, user *auth.DeleteUser) error {
	query := "UPDATE auth SET del = $1 WHERE id = $2"

	_, err := r.client.Exec(ctx, query, true, user.Id)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return newErr
        }
        return err
    }
	
	return nil
}

func (r *repository) Recover(ctx context.Context, user *auth.RecoverUser) error {
	query := "UPDATE auth SET del = $1 WHERE id = $2"

	_, err := r.client.Exec(ctx, query, false, user.Id)
	if err!= nil {
		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            pgErr = err.(*pgconn.PgError)
            newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
            r.log.Error("auth.Update: %s",newErr)
            return newErr
        }
        return err
    }
	
	return nil
}
