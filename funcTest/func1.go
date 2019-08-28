package funcTest

import (
    "bytes"
    "fmt"
    "io"
)

func Func1() func() {
    fmt.Println()

    fmt.Println("Func1")
    return func() {
        fmt.Println("Func1-return-anonymous-func")
    }
    //返回一个函数定义供调用者调用该函数，这实际上是一个函数值，即字面量形式的匿名函数

    var writer = io.Writer()
    var buf *bytes.Buffer
}
