package main

import "fmt"

func main() {
    // If/Else
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }

    // For loop
    fmt.Print("Counting: ")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // Switch
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the work week!")
    case "Friday":
        fmt.Println("Almost weekend!")
    default:
        fmt.Println("Just another day.")
    }
}
