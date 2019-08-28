package funcTest

import (
    "fmt"
    "testing"
)

func TestFunction(t *testing.T) {
    var f1 = Func1()
    f1()

    fmt.Println()

    //命名函数值，且其没有初始化，其值为nil
    //var f2 func()
    //fmt.Println(f2)

    //函数声明
    //func f3()

    //函数定义,函数内部不能在声明或者定义函数了，这只是一个定义形式的展示，
    //但是函数内部可以通过函数值的形式来声明或者定义新的函数
    //func f3() {}

    //命名函数值，他通过函数字面量来初始化
    var f4 = func() string { return "func3 is func definition,it's a result as its called" }
    //直接输出函数值本身会返回一个地址值
    //fmt.Println(f4)
    fmt.Println(f4())

}
