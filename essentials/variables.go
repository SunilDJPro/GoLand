package main

import "fmt"

func main() {
    // Explicit type declaration
    var age int = 25
    
    // Type inference
    var name = "Gopher"
    
    // Short variable declaration (most common inside functions)
    isLearning := true
    
    fmt.Printf("Name: %s, Age: %d, Learning: %v\n", name, age, isLearning)
}
