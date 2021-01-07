package main

import (
    "fmt"
)



func bubbleSort(input []int) {
    
    
    for i := 0; i < 10; i++ {
            
                for j := 0; j < 10-i-1; j++{
                if input[j]>input[j+1]{
                swap(j,input)
                                }
                }
            }
            fmt.Printf("%v", input)
            fmt.Println()
        }
func swap(j int, input []int){


    var t int=input[j]
    input[j]=input[j+1]
    input[j+1]=t
}
    
    
    

func main() {
    toBeSorted := make([]int, 10)
    fmt.Println("Enter array elements")
    for i := 0; i < 10; i++ {
		fmt.Scan(&toBeSorted[i])
	}
    bubbleSort(toBeSorted)
}
