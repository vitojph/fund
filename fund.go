package main

import "fmt"

type Fund struct {
    // balance is unexported (private), because it's lowercase
    balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
    // We can return a pointer to a new struct without worrying about
    // whether it's on the stack or heap: Go figures that out for us.
    return &Fund{
        balance: initialBalance,
    }
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
    return f.balance
}

func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}

func main() {
    fund := NewFund(1000)

    // Burn through them one at a time until they are all gone
    for i := 0; i < fund.Balance(); i++ {
        fund.Withdraw(1)
        fmt.Printf("%d\n", fund.Balance())
    }


    fmt.Printf("\n")
}
