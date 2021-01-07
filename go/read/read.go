package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strings"

type name struct {
    fname string
    lname string
}

func main(){
names := make([]name, 0)
var file_name string
fmt.Print("enter filename :- ")
fmt.Scanf("%s", &file_name)
file,err:= os.Open(file_name)
if err != nil {
	log.Fatalf("failed opening file: %s", err)
	}
 
scanner := bufio.NewScanner(file)
scanner.Split(bufio.ScanLines)
 
for scanner.Scan() {
r := strings.Split(scanner.Text(), " ")
names = append(names,name{r[0], r[1]})
	}
file.Close()
	
for i,j := range names {
fmt.Print("First name:- "+j.fname+" ")
fmt.Print("Last name:- "+j.lname)
fmt.Println()
i=i
	}
	

}
