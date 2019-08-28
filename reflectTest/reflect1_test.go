package reflectTest

import (
    "fmt"
    "reflect"
    "testing"
)

func TestFlect1(t *testing.T){

    var testVar = Var{"fitch", 186}

    //index starts from zero
    name := reflect.ValueOf(testVar).Type().Field(1).Name
    fmt.Println(name)

    typ := reflect.ValueOf(testVar).Type().Field(1).Type
    fmt.Println(typ)

    val := reflect.ValueOf(testVar).Field(1)
    fmt.Println(val)

    kind := reflect.ValueOf(testVar).Type().Kind()
    fmt.Println(kind)

    //key := reflect.ValueOf(testVar).Type().Key()
    //fmt.Println(key)

    //elem := reflect.ValueOf(testVar).Type().Elem()
    //fmt.Println(elem)

    var testVar1 = Var{name: "fitch", tel: 186}
    of := reflect.TypeOf(testVar1)
    fmt.Println(of)

    //reflect.New()
    //reflect.Zero()

}
