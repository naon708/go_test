package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
    howMuch("Alice", 100)

    fmt.Println(bravo())
}
func howMuch(name string, money int) {
    fmt.Printf("%s has %d yen\n", name, money)
}

func bravo() string {
    return "Bravo!"
}
