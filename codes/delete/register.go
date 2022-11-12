package main

import _ "fmt"

type Authdata struct {
    passwd  string
    login   string
    name    string
    age     uint8
}
func check_correct(data *Authdata) bool{
    //проверка на правильность использования символов
    return true
}
func check_DB(data *Authdata) bool{
    //проверка на соответствие данным в БД
    return true
}
func write_data(data *Authdata) bool{
    return true
}
func  sugnin(data *Authdata) bool {
    // data.passwd = 
    // data.login =
    if check_correct(data) == false {
        return  false
    }
    if check_DB(data) == false {
        return false
    }
    return true
}
func register(data *Authdata) bool {

    //получение пароля и логина и других возможных данных от клиента

    if check_correct(data) == false {
        return  false
    }
    if check_DB(data) == false {
        return false
    }
    return true
}


func main() {
    data := new(Authdata)
    if register(data) == false{
        write_data(data)
    }
    //signin()
}
