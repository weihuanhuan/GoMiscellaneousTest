package interfaceTest

import (
    "fmt"
    "testing"
)

func TestInterface(t *testing.T) {

    var bc ByteCounter

    bc.Write([]byte("1"))
    fmt.Println(bc)
    bc.Write([]byte("3"))
    fmt.Println(bc)
    bc = 0
    fmt.Println(bc)

    fmt.Println()
    fmt.Fprintf(&bc,"%s","5")
    fmt.Println(bc)
    fmt.Fprintf(&bc,"%s","7")
    fmt.Println(bc)
    bc=0
    fmt.Println(bc)


    fmt.Println()
    fmt.Fprintln(&bc,[]byte("8"))
    fmt.Println(bc)
    fmt.Fprintln(&bc,"9")
    fmt.Println(bc)
    bc=0
    fmt.Println(bc)


}
