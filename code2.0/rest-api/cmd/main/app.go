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
const(
	REGISTER_ = 3604
	ENTER_ = 3605
	DELETE_ = 3606
	GET__ = 3704
	CHANGE__ = 3705
	DELETE__ = 3706
)
func manage(entr *model.AuthData, action int){
	
	
	db := postgres.ConnectDB()			//подключаемся к базе данных
	defer db.Close()					//разрываем подключение в конце функции
	
	if postgres.FindData(db, entr){	//ищим совпадение в таблице
		fmt.Println("Совпадение найдено")
	} else {
		fmt.Println("Совпадение не найдено")
	}
	switch action {
	case REGISTER_: //поиск пользователя	добавление в таблицы!	возврат данных
		
	case ENTER_:	//поиск пользователя	возврат данных
	case DELETE_:	//поиск пользователя	удаление данных
	case GET__:		//поиск пользователя	возврат данных
	case CHANGE__:	//поиск пользователя	изменение данных
	case DELETE__:	//поиск пользователя	удаление данных
	}
	
	postgres.GetTable(db)				//получаем таблицу
	
	
	
}

func main(){
	
	entr := model.AuthData{		//вводные пользовательские данные
		Login: "serpan",
		Access_token: "access_token_01",
	}
	manage(&entr, GET__)
}