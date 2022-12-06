package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"go.mod/internal/postgres/model"
)

const (
	conn = "user=serpc password=12345 dbname=serpc sslmode=disable"
	table_auth = "UserAuth"
	table_data = "UserData"
)

// func Auth(db *sql.DB, adat *model.AuthData) bool{
	
// 	return true
// }

func ConnectDB() (db *sql.DB, err error){				//подключение к базе данных
	db, err = sql.Open("postgres", conn)	//подключаемся к базе данных, если не получается, то метод вернет ошибку в err
	//if err != nil {panic(err)}				//если err не нулевой указатель, то паникуем
	fmt.Println("Connect db")
	return db, err								//возвращаем указаетль на "бд"
}

func FindData(db *sql.DB, tmp *model.AuthData) (result sql.Result, err error)  {	//Поиск данных пользователя в таблице. Принимает указатели на объект БД и структуру аутентификации, введенную пользователем. Возвращает логическое значение 
	query_buff := "SELECT * from " + table_auth		//Запрос
	
	result, err = db.QueryContext(ctx, query_buff)				//Выполняем запрос
	if err != nil {
		panic(err)
	}
	for result.Next() {		//считываем данные по аналогии с GetTable()
		ch_data := model.AuthData{}		//структура аутентификации
		err := result.Scan(&ch_data.Id, &ch_data.Login, &ch_data.State, &ch_data.Access_token, &ch_data.Refresh_token) 
		//if err != nil {panic(err)}

		if ch_data.Login == tmp.Login {		//если значения полей сопадают, то возвращаем true и заполняем нужные данные для формирования ответа клиенту
			if ch_data.Access_token == tmp.Access_token {
				tmp.Id = ch_data.Id
				tmp.State = ch_data.State
				return true, err 
			}
		}
	}
	return false, err		//если нет пользователя , то возвращаем false
}

func AddData(db *sql.DB, tmp *model.AuthData) (result sql.Result, err error) {		//Добавляем значение в таблицу
	query_buff := "insert into " + table_auth + " (login, state, access_token, refresh_token) values ('" + tmp.Login + "',"+strconv.FormatBool(tmp.State)+", '" + tmp.Access_token + "', '"+ tmp.Refresh_token +"')"
	result, err = db.Exec(query_buff)
	return result, err

}

func AddUserData(db *sql.DB, tmp *model.UserData, id int) (result sql.Result, err error) {		//Добавляем значение в таблицу
	query_buff := "insert into " + table_data + " (id, user_name, sex, birthdate, weight) values ("+strconv.Itoa(id)+", '" + tmp.Name + "', '" + tmp.Sex + "', '"+ tmp.Bithdate +"', "+strconv.Itoa(tmp.Weight)+")"

	result, err = db.Exec(query_buff)
	return result, err
	
}

func DeleteAuthData(db *sql.DB, tmp *model.AuthData) (result sql.Result, err error){	//удаляем данные из таблицы
	query_buff := "delete from "+table_auth+" where login = " + tmp.Login
	
	result, err = db.Exec(query_buff, 2)
	
	return result, err
}
func DeleteUserData(db *sql.DB, tmp *model.AuthData) (result sql.Result, err error){	//удаляем данные из таблицы
	query_buff_2 := "delete from "+table_data+" where login = " + tmp.Login
	
	result, err = db.Exec(query_buff_2, 2)
	return result, err
}