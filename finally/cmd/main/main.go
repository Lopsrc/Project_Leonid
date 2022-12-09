package main

import (
	"context"
	"fmt"

	// "go.mod/internal/authdata"
	"go.mod/internal/authdata/db"
	"go.mod/internal/authdata"
	"go.mod/internal/config"
	"go.mod/internal/userdata"
	"go.mod/internal/userdata/db"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
)

//
//проверить findone() -> user:
// error: can't scan into dest[3]: unable to assign to *string. Но данные считывает корректно, кроме даты 

	


const (
	ENTER = 1	//enter
	REGISTER =2	//create auth data 
	CREATE = 6 //create user data
	UPDATEA = 3	//change auth data
	UPDATEU = 7 //change user data
	FINDU = 8	//find user data
	DELETEA = 5 //delete auth data
	DELETEU = 9 //delete userdata	
)

type UserState struct{
	action_db int
	find_state bool
}
// func Register(repo authdata.Repository ,auth *authdata.AuthData) (bool, error){
// 	res, err := repo.FindOne(context.TODO(), auth)
// 	if err != nil{return false, err}
// 	return true, nil
// }
func main(){
	
	action := UserState{
		action_db: FINDU,
	}
	auth := authdata.AuthData{
		Login: "serpan2002@mail.ru",
		State: true,
		Access_token: "access_token_8",
		Refresh_token: "refresh_token_8",
	}	
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	
	client , err :=postgresql.NewClient(context.TODO(),3, cfg.Storage)
	if err!=nil{logger.Fatal("%v", err)}

	repository := authdb.NewRepository(client, logger)
	repositoryUserData := userdb.NewRepository(client, logger)

	action.find_state = repository.FindOne(context.TODO(), &auth)
	fmt.Println(action.find_state)
	// panic("1")

	data := userdata.UserData{
		Id: auth.Id,
		Name: "Sofia",
		Sex: "woman",
		Birthdate: "2001-04-04",
		Weight: 80,
	}

	fmt.Println(data.Id)
	switch action.action_db {
	case ENTER: //проверено
		if action.find_state {
			fmt.Println("Пользователь не найден, вход невозможен")
		}
		fmt.Println("Вход разрешен")
	case REGISTER://проверено
		if action.find_state {
			fmt.Println("Пользователь найден, регистрация невозможна")
		}else {
			err = repository.Create(context.TODO(), &auth)
			if err!=nil{logger.Fatalf("%v", err)}
		}
	case CREATE://проверено
		if !(action.find_state) {
			fmt.Println("Зарегистрированный пользователь не найден, создание пользователя невозможно")
		}else {
			s, err := repositoryUserData.FindOne(context.TODO(), &data)
			if s {
				fmt.Println("Создание пользователя невозможно, тк такой пользователь существует")
				logger.Fatalf("%v", err)
				
			}
			
			err = repositoryUserData.Create(context.TODO(), &data)
			if err!=nil{logger.Fatalf("%v", err)}
		}
	case UPDATEA://провернео
		if !(action.find_state) {
			fmt.Println("Пользователь не найден, изменения невозможны")
		}else {
			err = repository.Update(context.TODO(), &auth)
			if err!=nil{logger.Fatalf("%v", err)}
		}
	case UPDATEU://проверено
		if !(action.find_state) {
			fmt.Println("Пользователь не найден, изменения невозможны")
		}else {

			err = repositoryUserData.Update(context.TODO(), &data)
			if err!=nil{logger.Fatalf("%v", err)}
		}
	case FINDU: //проверено
		find_state_user, err := repositoryUserData.FindOne(context.TODO(), &data)
		if !(find_state_user) {
			fmt.Println("Пользователь не найден")
			logger.Fatalf("%v", err)
		}
		fmt.Println("Пользователь найден")
		fmt.Println(data)


	case DELETEA: //проверено
		if !(action.find_state) {
			fmt.Println("Пользователь не найден, удаление невозможно")
		}else {
			fmt.Println(auth.Id)
			err = repository.Delete(context.TODO(), auth.Id)
			if err!=nil{logger.Fatalf("%v", err)}
			fmt.Println(auth.Id)
			err = repositoryUserData.Delete(context.TODO(), data.Id)
			if err!=nil{logger.Fatalf("%v", err)}
		}
	case DELETEU: //проверено
		if !(action.find_state) {
			fmt.Println("Пользователь не найден, удаление невозможно")
		}else {
			fmt.Println(auth.Id)
			err = repositoryUserData.Delete(context.TODO(), data.Id)
			if err!=nil{logger.Fatalf("%v", err)}
		} 
	}
	fmt.Println("fuser")	
}