package main

import "fmt"
import "strings"

func main(){

var s string
    fmt.Print("enter a string :- ")
    fmt.Scanf("%s", &s)
    s = strings.ToLower(s)
    if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "i"){
     fmt.Println("Found!")
     }else {
     fmt.Println("Not Found!")
     }
     
}
