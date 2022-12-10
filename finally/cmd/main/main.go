package main

import (
	"context"
	"fmt"
	"time"

	// "go.mod/internal/authdata"
	"github.com/jackc/pgtype"
	"go.mod/internal/authdata"
	"go.mod/internal/authdata/db"
	"go.mod/internal/config"
	"go.mod/internal/userdata"
	"go.mod/internal/userdata/db"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
)

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

func GetData(years int, month int, day int) (date time.Time){
	return date.AddDate(years-1,month-1,day-1)
} 

func main(){
	
	action := UserState{
		action_db: DELETEA,
	}
	auth := authdata.AuthData{
		Login: "sofia01@mail.ru",
		State: true,
		Access_token: "access_token_1",
		Refresh_token: "refresh_token_1",
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
	// tm := time.Time{}

	dt := pgtype.Date{
		Time: GetData(2005,7,14),
	}
	
	data := userdata.UserData{
		Id: auth.Id,
		Name: "Sofia",
		Sex: "woman",
		Birthdate: dt,
		Weight: 55,
	}

	fmt.Println(data.Id)
	switch action.action_db {
	case ENTER: //проверено
		if !(action.find_state) {
			fmt.Println("Пользователь не найден, вход невозможен")
			panic("Sosi huy")
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
			data.Id =1

			find_state_user, err := repositoryUserData.FindOne(context.TODO(), &data)
			if !(find_state_user) {
				fmt.Println("Пользователь не найден")
				logger.Fatalf("%v", err)
			}
			fmt.Println(auth.Id)
			err = repositoryUserData.Delete(context.TODO(), data.Id)
			if err!=nil{logger.Fatalf("%v", err)}
		} 
	}
	fmt.Println("fuser")	
}