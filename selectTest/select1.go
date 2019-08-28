package selectTest

import (
    "fmt"
    "time"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

//go init func
func init() {
    fmt.Println("init start")
    go teller() // start the monitor goroutine
}

func Deposit(amount int) {
    fmt.Println("Deposit start")
    deposits <- amount
}
func Balance() int {
    fmt.Println("Balance start")
    return <-balances
}

func teller() {
    fmt.Println("teller start")
    var balance int // balance is confined to teller goroutine
    for {
        time.Sleep(20 * time.Millisecond)
        select {
        case amount := <-deposits:
            fmt.Println("case 1")
            balance += amount
        case balances <- balance:
            fmt.Println("case 2")
        default:
            fmt.Println("default")
        }
    }
}

