package main

import "fmt"
import "encoding/json"

func main(){
var name string
var address string
fmt.Print("enter a name :- ")
fmt.Scanf("%s", &name)
fmt.Print("enter an address :- ")
fmt.Scanf("%s", &address)

m := map[string]string{"name": name, "address": address}
j, _ := json.Marshal(m)
fmt.Println(string(j))

}
