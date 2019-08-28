package selectTest

import (
    "testing"
    "time"
)

func TestSelect1(t *testing.T) {

    Deposit(100)
    time.Sleep(100 * time.Millisecond)
    Balance()

}
