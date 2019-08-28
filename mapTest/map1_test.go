package mapTest

import (
    "fmt"
    "testing"
)

func TestMap(t *testing.T) {
    fmt.Println()

    var map1 map[string]string
    //map1["key1"]="val1"
    //panic: assignment to entry in nil map
    //注意Println是nil安全的，即使是nil的类型也不会出现错误，
    // 比如这里值为nil的map，但是不确定是否所有类似的nil都可以处理
    fmt.Println(map1)

    var map2 map[string]string
    {
    }
    //这里是另一个语法块，其实和map02的声明已经没有关系了，所以map2和map1其实是一样的。都是值为nil的map
    // 编译器在声明语句的末尾已经判定为另一个语句了，即末尾已经添加了;(分号)。
    fmt.Println(map2)

    var map3 = map[string]string{}
    map3["key3"] = "val3"
    //这里是声明并且定义了一个map，其值是空，并不是nil，=(复制)操作为他分配了一个值为空的map。
    // 和上面的单纯声明不一样的是，这里map的值不是系统默认赋予他的nil
    fmt.Println(map3)

    map4 := map[string]string{"key4": "val4"}
    fmt.Println(map4)

    var map5 = make(map[string]string, 10)
    map5["key5"] = "val5"
    fmt.Println(map5)

}
