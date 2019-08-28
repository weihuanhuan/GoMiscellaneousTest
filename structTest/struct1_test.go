package structTest

import (
    "fmt"
    "testing"
)

func TestStruct(t *testing.T) {
    fmt.Println()

    var structNamedWithoutVal structNamed
    var structNamedWithVal = structNamed{1, 2}
    var structNamedWithNameVal = structNamed{a: 1, b: 2}

    fmt.Println(structNamedWithoutVal)
    fmt.Println(structNamedWithVal)
    fmt.Println(structNamedWithNameVal)

    //不带初始值的字面量形式
    var structAnonymousWithoutVal = struct {
        a, b int
    }{}

    //带初始值的字面量形式
    var structAnonymousWithVal = struct {
        a, b int
    }{3, 4}

    fmt.Println(structAnonymousWithoutVal)
    fmt.Println(structAnonymousWithVal)
}
