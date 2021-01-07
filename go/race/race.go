package main 
  
import ( 
    "fmt"
    "time"
) 
  
func execute(some string) { 
  
    
    for i := 1; true; i++ { 
  
        fmt.Println(some, i) 
  
        time.Sleep(time.Millisecond * 100) 
  
    } 
} 
  
func main() { 
  
    execute("First") 
      
    execute("Second") 
      
    
    fmt.Println("program ends successfully") 
      
} 

/*Race condition is defined as the condition of an electronics, software, or other systems where the system’s substantive behavior is dependent on the sequence or timing of other uncontrollable events.
In the above code execute is a simple function that prints some text passed to it and an integer number. The loop inside this function is an infinite loop. Inside the loop the text along with the integer is printed on to the screen and then the code remains idle for some seconds. The first execute function call will never get finished and this makes the second execute function to wait for an infinite amount of time. It won't get executed until the first execute function call finishes. All lines below the first execute function call will never get executed.*/
