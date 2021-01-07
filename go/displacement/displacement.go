package main

import (
    "fmt"
)


func GenDisplaceFn(a , v0, s0 float64) func (float64) float64{

fn:= func(t float64) float64{

return ((0.5*a*t*t) + (v0*t) + (s0))

}

return fn
}


func main(){

var a , v0, s0 , t float64

fmt.Println("Enter value for acceleration")
fmt.Scan(&a)

fmt.Println("Enter value for initial velocity")
fmt.Scan(&v0)

fmt.Println("Enter value for initial displacement")
fmt.Scan(&s0)

fn := GenDisplaceFn(a, v0, s0)

fmt.Println("Enter value for time")
fmt.Scan(&t)

fmt.Println(fn(t))


}
