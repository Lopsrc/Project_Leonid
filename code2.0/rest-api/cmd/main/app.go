package main

import (
	//"database/sql"

	"fmt"

	_ "github.com/lib/pq"
	"go.mod/internal/postgres"
	"go.mod/internal/postgres/model"
)
//			Code actions
//	3604 - Регистрация
//	3605 - Вход
//  3606 - Удалить аккаунт
//	3704 - Получить данные о себе
//	3705 - Изменить данные о себе
//	3706 - Удалить данные о себе
type ParamsConn struct{
	find_state bool
	action_state int
}
const(
	REGISTER_ = 3604
	ENTER_ = 3605
	DELETE_ = 3606
	GET__ = 3704
	ADD_ = 3705
	CHANGE__ = 3706
	DELETE__ = 3707
)
func manage(entr_auth *model.AuthData, entr_data *model.UserData  ,pconn *ParamsConn) string {

	
	db, err := postgres.ConnectDB()			//подключаемся к базе данных
	if err!=nil{panic(err)}
	
	defer db.Close()					//разрываем подключение в конце функции
	
	fnd_state, err := postgres.FindData(db, entr_auth)
	if err!=nil{panic(err)}

	if fnd_state{	//ищим совпадение в таблице, реализация аутентификации.
		fmt.Println("Совпадение найдено")
		pconn.find_state = true
	} else {
		fmt.Println("Совпадение не найдено")
		pconn.find_state = false
	}


	switch pconn.action_state {
	case REGISTER_: //поиск пользователя	добавление в таблицы!	возврат данных
		if pconn.find_state {return "Пользователь найден, регистрация невозможна"}

		_, err := postgres.AddData(db,entr_auth)
		if err!=nil {panic(err)}
		//if !(postgres.FindData(db,entr_auth)){return "Данные не найдены на этапе проверки после добавления в auth table"}
		return "Пользователь добавлен в базу"
		
	case ENTER_:	//поиск пользователя	возврат данных
		if pconn.find_state {return "Пользователь найден, регистрация невозможна"}

		return "Возврат данных" //придумать либо возврат кода по которому будет создаваться отчет, либо сразу составлять отчет и возвращать в main

	case DELETE_:	//поиск пользователя	удаление данных
		if !(pconn.find_state) {return "Пользователь не найден, удаление невозможно"}

		err := postgres.DeleteData(db, entr_auth)
		if err!=nil{panic(err)}

	case GET__:		//поиск пользователя	возврат данных
		if pconn.find_state {return "Пользователь найден, регистрация невозможна"}
	case CHANGE__:	//поиск пользователя	изменение данных
		if !(pconn.find_state) {return "Пользователь не найден, изменения невозможны"}

		err := postgres.AddUserData(db, entr_data, entr_auth.Id)
		if err!=nil{panic(err)}
		return "Данные записаны"

	case ADD_:
		if pconn.find_state {return "Пользователь найден, регистрация невозможна"}

		err := postgres.AddUserData(db, entr_data, entr_auth.Id)
		if err != nil {panic(err)}
	case DELETE__:	//поиск пользователя	удаление данных
		if pconn.find_state {return "Пользователь найден, регистрация невозможна"}


	default:
		fmt.Println("Некорректный код действия action_state")
	}
	
	//postgres.GetTable(db)				//получаем таблицу
	
	
	return ""
}

func main(){

	pconn := ParamsConn{
		find_state: false,		//нашли ли данные 
		action_state: CHANGE__,	//код действия пользователя 
	}
	entr_user := model.UserData{
		Name: "Sergey",
		Sex: "man",
		Bithdate: "2002-04-04",
		Weight: 70,
	}
	entr_auth := model.AuthData{
		Id:            0,
		Login:         "serpan2002@mail.ru",
		State:         false,
		Access_token:  "Access_token_01",
		Refresh_token: "refresh_token_01",
	}
	fmt.Println(manage(&entr_auth, &entr_user ,&pconn))
}