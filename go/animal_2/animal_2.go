package main

import "fmt"
import "strings"
import "bufio"
import "os"

type Animal interface{
    
    Eat()
    Move()
    Speak()

}

type Cow struct{

    food string
    locomotion string
    noise string

}

func (c Cow) Eat(){

    fmt.Println(c.food)

}

func (c Cow) Move(){

    fmt.Println(c.locomotion)

}

func (c Cow) Speak(){

    fmt.Println(c.noise)

}

type Bird struct{

    food string
    locomotion string
    noise string

}

func (b Bird) Eat(){

    fmt.Println(b.food)

}

func (b Bird) Move(){

    fmt.Println(b.locomotion)

}

func (b Bird) Speak(){

    fmt.Println(b.noise)

}

type Snake struct{

    food string
    locomotion string
    noise string

}

func (s Snake) Eat(){

    fmt.Println(s.food)

}

func (s Snake) Move(){

    fmt.Println(s.locomotion)

}

func (s Snake) Speak(){

    fmt.Println(s.noise)

}



func main(){
    
    m := make(map[string]Animal)
    in := bufio.NewReader(os.Stdin)
    
    for ;true; {
        
        fmt.Printf(">")
        s, err := in.ReadString('\n')
        
        
        if err!=nil{
            
            fmt.Println(err)
            return
        
        }
        
	res := strings.Fields(s)
	
		
	if len(res) <= 0{
	
	    return
	    
	}
	
	if res[0] == "newanimal"{
	
	    switch res[2] {
		    
	        case "cow":
			    
	            m[res[1]] = Cow{"grass","walk","moo"}
	            fmt.Println("Created it!")
			    
	        case "bird":
			    
	            m[res[1]] = Bird{"worms","fly","peep"}
	            fmt.Println("Created it!")
			    
		case "snake":
			    
		    m[res[1]] = Snake{"mice","slither","hss"}
		    fmt.Println("Created it!")
			    
		default:
			    
		    fmt.Println("Animal not listed")
		    
		    } 
		    
        } else if res[0] == "query"{
        
              value, ok := m[res[1]]
              
              if ok {               
		    
              switch res[2] {
		    
	        case "eat":
			    
	            value.Eat()
			    
	        case "move":
			    
	            value.Move()
			    
		case "speak":
			    
		    value.Speak()
			    
		default:
			    
		    fmt.Println("No function")
		    
		    } 
		    
		    }else{
		    
		    fmt.Println("Key not found")
		    
		    }
			    
	} else{
       
            fmt.Println("Bad Command")
            
            }
		    
            

      }

}




