package deferTest

import "fmt"

func f1() {
    fmt.Println("f1")
}

func Defer1() {
    fmt.Println()

    defer f1()
}