package main

import "fmt"

func main() {
    fmt.Println("Start of main")
    
    defer fmt.Println("This will run last")  // Deferred function
    
    fmt.Println("Middle of main")
    
    // Another defer example
    defer fmt.Println("This runs before the last one")

	
    defer fmt.Println("This runs ?")
    
    fmt.Println("End of main")
}
