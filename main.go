package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
    howMuch("Alice", 100)
}
func howMuch(name string, money int) {
    fmt.Printf("%s has %d yen", name, money)
}
