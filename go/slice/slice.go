package main

import "fmt"
import "strconv"
import "sort"
func main(){

sli := make([]int,3)
var k string
for ;true;{
fmt.Print("enter an integer number :- ")
fmt.Scanf("%s", &k)
if k=="X"{
break
}
i,err := strconv.Atoi(k)
if err != nil {
        // handle error
        fmt.Println(err)
    }
sli = append(sli,i)
sort.Ints(sli)
fmt.Println(sli)
}
}
