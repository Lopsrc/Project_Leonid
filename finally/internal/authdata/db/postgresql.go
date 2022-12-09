package authdb

import (
	"context"
	// "errors"
	"fmt"
	// "strconv"
	"strings"

	// "github.com/jackc/pgconn"
	"go.mod/internal/authdata"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

// FindAll implements authdata.Repository


func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, authdata *authdata.AuthData) error { //создание пользователя
	q := `INSERT INTO userauth (login, state, access_token, refresh_token) VALUES ($1, $2, $3, $4) RETURNING id` //запрос

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	//выполнение запроса, заполнение поля id
	// if err := r.client.QueryRow(ctx, q, authdata.Login, authdata.State, authdata.Access_token, authdata.Refresh_token).Scan(&authdata.Id); err != nil {
	// 	var pgErr *pgconn.PgError
	// 	if errors.As(err, &pgErr) {
	// 		pgErr = err.(*pgconn.PgError)
	// 		newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
	// 		r.logger.Error(newErr)
	// 		return newErr
	// 	}
	// 	return err
	// }
	_, err := r.client.Exec(ctx, q,authdata.Login, authdata.State, authdata.Access_token, authdata.Refresh_token)
	if err!=nil{return err}
	
	return nil
}

func (r *repository) FindOne(ctx context.Context, auth *authdata.AuthData) (bool) { //поиск пользователя
	q := `SELECT id FROM userauth WHERE login = $1` //запрос на поиск по login
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	fmt.Println(auth.Login)
	err := r.client.QueryRow(ctx, q, auth.Login).Scan(&auth.Id) //выполнение запроса и заполнение полей созданной модели
	if err!=nil {return false}
	if err==nil{return true}
	
	return true
}

func (r *repository) Update(ctx context.Context, user *authdata.AuthData) error {
	q := `
		UPDATE userauth
		SET  state = $2
		WHERE id = $1
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	_, err := r.client.Exec(ctx, q, user.Id, user.State)
	if err!=nil{return err}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM userauth WHERE id=$1
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	_, err := r.client.Exec(ctx, q, id)
	if err!=nil{return err}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) authdata.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
