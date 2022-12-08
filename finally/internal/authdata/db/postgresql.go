package authdb

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
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
	if err := r.client.QueryRow(ctx, q, authdata.Login, authdata.State, authdata.Access_token, authdata.Refresh_token).Scan(&authdata.Id); err != nil {
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

// Оставлю может нужна будет

// func (r *repository) FindAll(ctx context.Context) (u []authdata.AuthData, err error) {			//Вряд ли нам нужен данный метод
// q := `SELECT id, login FROM AuthData;`							//если нужно вывести все поля, то лучше это сделать через консоль на сервере
// 	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

// 	rows, err := r.client.Query(ctx, q)
// 	if err != nil {
// 		return nil, err
// 	}

// 	authors := make([]authdata.AuthData, 0)

// 	for rows.Next() {
// 		var ath author.Author

// 		err = rows.Scan(&ath.ID, &ath.Name)
// 		if err != nil {
// 			return nil, err
// 		}

// 		authors = append(authors, ath)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return authors, nil
// }

func (r *repository) FindOne(ctx context.Context, auth *authdata.AuthData) (bool, error) { //поиск пользователя
	q := `SELECT id FROM userauth WHERE login = $1` //запрос на поиск по login
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	res, err := r.client.Query(ctx, q, auth.Login) //выполнение запроса и заполнение полей созданной модели
	if err != nil {
		panic(err)
		//return false, err
	}
	err = res.Scan(&auth.Id) //error number of field descriptions must equal number of values, got 5 and 0
	fmt.Println(auth.Id)
	if err!=nil {
		panic(err)
		// return false, err
	}
	return true, nil
}

func (r *repository) Update(ctx context.Context, user *authdata.AuthData) error {
	q := `
		UPDATE userauth
		SET  state = $2
		WHERE id = $1
		RETURNING id
	`
	
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	err := r.client.QueryRow(ctx, q, user.Id, user.State).Scan(&user.Id)
	if err!=nil{return err}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM userauth RETURNING id=$1
	`
	delete_me :=""
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	res, err := r.client.Query(ctx, q, strconv.Itoa(id))
	res.Scan(&delete_me)
	if err!=nil{return err}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) authdata.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
