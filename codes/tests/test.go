package main

import "fmt"

func prnt() (string, string, string){
    var name = "sergey"
    var lastname = "pan"
    return name,"\n" ,lastname
}
func main(){
    fmt.Print(prnt())
}
