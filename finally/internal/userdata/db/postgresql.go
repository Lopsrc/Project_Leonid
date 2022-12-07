package userdb

import (
	"context"
	"errors"
	"fmt"
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

// FindAll implements authdata.Repository
func (*repository) FindAll(ctx context.Context) (u []userdata.UserData, err error) {
	panic("unimplemented")
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, user *userdata.UserData, id int) error { //создание пользователя
	q := `INSERT INTO userdata (id, user_name, sex, birthdate, weight) VALUES ($1, $2, $3, $4, $5)` //запрос

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	//выполнение запроса, заполнение поля id
	if err := r.client.QueryRow(ctx, q, id, user.Name, user.Sex, user.Birthdate, user.Weight).Scan(&user.Id); err != nil {
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

func (r *repository) FindOne(ctx context.Context, id int) (userdata.UserData, error) { //поиск пользователя
	q := `SELECT id, user_name, sex, birthdate, weight FROM userdata WHERE id = $1` //запрос на поиск по login
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var ath userdata.UserData                                                                                            //модель для заполнения
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.Id, &ath.Name, &ath.Sex, &ath.Birthdate, &ath.Weight) //выполнение запроса и заполнение полей созданной модели
	if err != nil {
		return userdata.UserData{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, user userdata.UserData) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) userdata.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
