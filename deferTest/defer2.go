package deferTest

import (
    "fmt"
    "time"
)


func f2() func() {
    fmt.Println("f2")
    fmt.Println("before sleep###" + time.Now().String())
    time.Sleep(1 * time.Second)
    return func() {
        fmt.Println("f2-anonymous")
        fmt.Println("after defer###" + time.Now().String())
    }
    //返回一个函数定义供调用者调用该函数，这实际上是一个函数值，即字面量形式的匿名函数
}

func Defer2() {
    fmt.Println()

    //f2()返回的数据是一个函数，函数要有()才能表示对函数的调用，
    // 所以这里最后那对()是必要的，否则f2()将只执行f2函数本身，f2函数返回的函数值将不会被执行
    // 同时，defer将最后一层函数的执行推迟到return之后执行，其余的函数在执行defer语句时便执行了。
    defer f2()()
}
