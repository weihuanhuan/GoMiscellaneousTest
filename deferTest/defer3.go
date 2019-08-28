package deferTest

import (
    "fmt"
    "io"
)

type dummy struct{}

func Defer3() {
    fmt.Println()

    fmt.Println("defer3-enter")
    defer func() {
        fmt.Println("defer3-defer")
        switch p := recover(); p {
        case nil:
        case dummy{}:
            fmt.Println("dummy")
        default:
            fmt.Println("io.EOF")
            panic(io.EOF)
        }
    }()
    //defer 后面必须是一个函数调用，而不能是一个函数的定义，
    // 所以这里最后的()也不能省略,()代表了函数执行的动作
    //同时这里的函数调用是建立在函数值的基础上的，该函数值通过字面量的形式初始化的匿名函数

    fmt.Println("defer3-before-panic")
    panic(dummy{})
    //panic(io.EOF)
    fmt.Println("defer3-end")
}
