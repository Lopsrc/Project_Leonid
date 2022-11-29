package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.mod/internal/postgres/model"
)

const (
	conn = "user=serpc password=12345 dbname=serpc sslmode=disable"
	table_auth = "UserAuth"
	table_data = "UserData"
)

func ConnectDB() (db *sql.DB) {				//подключение к базе данных
	//connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", conn)	//подключаемся к базе данных, если не получается, то метод вернет ошибку в err
	if err != nil {panic(err)}				//если err не нулевой указатель, то паникуем
	fmt.Println("Connect db")
	return db								//возвращаем указаетль на "бд"
}

func GetTable(db *sql.DB) {			//получаем всю таблицу
	query_buff := "SELECT * from " + table_auth		//Запрос
	ch_data := model.AuthData{}						//Структура аутентификации
	result, err := db.Query(query_buff)				//Отправляем запрос
	if err != nil {panic(err)}						

	for result.Next() { //с помощью .Next() возвращаем логическое значение о наличии строки в таблице
		err := result.Scan(&ch_data.Id, &ch_data.Login, &ch_data.State, &ch_data.Access_token, &ch_data.Refresh_token) //запись в структуру аутентификации данных строки
		if err != nil {panic(err)}
		fmt.Println(ch_data)	//вывод в консоль строки
	}
}

func FindData(db *sql.DB, tmp *model.AuthData) bool {	//Поиск данных пользователя в таблице. Принимает указатели на объект БД и структуру аутентификации, введенную пользователем. Возвращает логическое значение 
	query_buff := "SELECT * from " + table_auth		//Запрос
	
	result, err := db.Query(query_buff)				//Выполняем запрос
	if err != nil {panic(err)}
	//считытвать данные из таблицы по login, и сравнивать с данными, переданными пользователем, в случае ошибки
	//вернуть сообщение об ошибки (если совпадения найдены),  хотя зачем если у нас OAuth, который отдает
	//клиенту токены и логины, которые в случае не совпадения добавляются в базу данных, тк наша задача хранить
	// лишь логин, токены и роли, поэтому мы должны ища логин сравнивать токены. Но как именно гугл предоставляет
	//API для работы с токенами доступа и токенами обновления нужно выяснить. Если мы храним лишь рефреш токен, а аксесс токен выдается
	//гуглом, а мы лишь меняем аксесс токен с помощью рефреш токена при смерти аксесс токена, то впринципе понятно как происходит взаимодествие
	// с сервером, но если структура иная... короче нужно лучше изучить OAuth2.0 лучше
	for result.Next() {		//считываем данные по аналогии с GetTable()
		ch_data := model.AuthData{}		//структура аутентификации
		err := result.Scan(&ch_data.Id, &ch_data.Login, &ch_data.State, &ch_data.Access_token, &ch_data.Refresh_token) 
		if err != nil {panic(err)}

		if ch_data.Login == tmp.Login {		//если значения полей сопадают, то возвращаем true и заполняем нужные данные для формирования ответа клиенту
			if ch_data.Access_token == tmp.Access_token {
				tmp.Id = ch_data.Id
				tmp.State = ch_data.State
				return true
			}
		}
	}
	return false		//если нет пользователя , то возвращаем false
}

func AddData(db *sql.DB, tmp *model.AuthData) {		//Добавляем значение в таблицу
	query_buff := "insert into" + table_auth + " (id, login, token) values (26, '" + tmp.Login + "', '" + tmp.Access_token + "')"
	result, err := db.Exec(query_buff)
	if err != nil {panic(err)}
	fmt.Println(result.RowsAffected())

}

func DeleteAuthData(db *sql.DB, tmp *model.AuthData) {	//удаляем данные из таблицы
	query_buff := "delete from "+table_auth+" where login = " + tmp.Login

	result, err := db.Exec(query_buff, 2)
	if err != nil {panic(err)}

	fmt.Println(result.RowsAffected()) // количество удаленных строк
}