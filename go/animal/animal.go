package main

import "fmt"
import "strings"
import "bufio"
import "os"

type Animal struct{
    
    food string
    locomotion string
    noise string

}

func (an *Animal) Eat(){

    fmt.Println(an.food)
    
}

func (an *Animal) Move(){

    fmt.Println(an.locomotion)
    
}

func (an *Animal) Speak(){

    fmt.Println(an.noise)
    
}

func main() {

    cow := Animal{"grass","walk","moo"}
    
    bird := Animal{"worms","fly","peep"}
    
    snake := Animal{"mice","slither","hss"}
    
    
    in := bufio.NewReader(os.Stdin)
    
    for ;true; {
        
        fmt.Printf(">")
        
        s, err := in.ReadString('\n')
        
        
        if err!=nil{
        
            break
        
        }
        
	res := strings.Fields(s)
	
	if len(res)<=0{
	
	    break
	    
	}
	 
        if res[0] == "cow"{
		
            switch res[1] {
		    
	        case "eat":
			    
	            cow.Eat()
			    
	        case "move":
			    
	             cow.Move()
			    
		case "speak":
			    
		      cow.Speak()
			    
		default:
			    
		      fmt.Println("Option not listed")
		    
		    }
		
		
		} else if res[0] == "bird"{
		
		    switch res[1] {
		    
			case "eat":
			    
			   bird.Eat()
			    
			case "move":
			    
			   bird.Move()
			    
			case "speak":
			    
			    bird.Speak()
			    
			default:
			    
			    fmt.Println("Option not listed")
		    
		    }
		
		
		} else if res[0] == "snake"{
		
		    switch res[1] {
		    
			case "eat":
			    
			    snake.Eat()
			    
			case "move":
			    
			    snake.Move()
			    
			case "speak":
			    
			    snake.Speak()
			    
			default:
			    
			    fmt.Println("Option not listed")
		    
		    }
		
		
		} else{
       
            fmt.Println("Bad query structure")
            
            }
		
	    }}

