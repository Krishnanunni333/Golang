package main

import (
   "flag"
   "os"
   "fmt"
   )

func main(){

   csvFilename := flag.String("csv","problems.csv","a csv file in the format of 'question,answer'")
   flag.Parse()
   
   
   file,err:= os.Open(*csvFilename)
   
   if err != nil {
   
   exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
   }
   
   r := csv.NewReader(file)
   lines,err:=r.ReadAll()
   
   if err != nil {
   
   exit("Failed to parse the provided CSV file")
   }
   
   fmt.Println(lines)
}

func exit(msg string){
  
  fmt.Println(msg)
  os.Exit(1)

}
