package userdb

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"github.com/jackc/pgconn"
	"go.mod/internal/userdata"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, user *userdata.UserData) error { //создание пользователя
	q := `INSERT INTO userdata (id, user_name, sex, birthdate, weight) VALUES ($1, $2, $3, $4, $5)` //запрос

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	//выполнение запроса, заполнение поля id
	if err := r.client.QueryRow(ctx, q, strconv.Itoa(user.Id) , user.Name, user.Sex, user.Birthdate,strconv.Itoa(user.Weight) ).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}


func (r *repository) FindOne(ctx context.Context, user *userdata.UserData) (bool, error) { //поиск пользователя
	q := `SELECT id, user_name, sex, birthdate, weight FROM userdata WHERE id = $1` //запрос на поиск по login
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	_, err := r.client.Query(ctx, q, strconv.Itoa(user.Id)) //выполнение запроса и заполнение полей созданной модели
	if err!=nil {
		return false, err
	}
	return true, nil
}

func (r *repository) Update(ctx context.Context, user *userdata.UserData) error {
	q := `
		UPDATE userdata
		SET user_name = $2, sex = $3, birthdate = $4, weight = $5
		WHERE id = $1
		RETURNING id
	`
	
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	err := r.client.QueryRow(ctx, q, user.Id, user.Name, user.Sex, user.Birthdate, user.Weight).Scan(&user.Id)
	if err!=nil{return err}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM userdata RETURNING id=$1
	`
	
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	err := r.client.QueryRow(ctx, q, strconv.Itoa(id)).Scan(id)
	if err!=nil{return err}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) userdata.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
