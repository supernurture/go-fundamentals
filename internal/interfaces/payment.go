package interfaces

import (
	"fmt"

	zg "github.com/rs/zerolog/log"
)

/*

1. DEFINITION & FUNCTION
   - Definition: A contract defining a set of method signatures without implementation.
   - Key Concept: Uses implicit implementation (no "implements" keyword).
   - Benefits: Polymorphism, decoupling (easier unit testing/mocking), and abstraction.

2. DRAWBACKS & PERFORMANCE IMPACT
   - Performance: Minor overhead (nanoseconds) due to dynamic dispatch (itable lookup)
     and potential heap allocation (escape analysis).
   - Drawbacks: Over-engineering risk, slightly harder code navigation.

3. POINTER VS. VALUE RECEIVERS
   - Pointer Receiver (*T): Must pass pointer (&T) to interface. Use when mutating data
     or working with large structs.
   - Value Receiver (T): Accepts both values (T) and pointers (&T). Use for small,
     immutable structs.

4. NAMING CONVENTIONS
   - Single-method interfaces use the method name + "-er" suffix (e.g., Read -> Reader).
   - For a single Pay() method, the idiomatic interface name is Payer.

*/

type Payer interface {
	Pay(amount float64) error
}

type Xendit struct {
	Balance float64
}

func (x *Xendit) Pay(amount float64) error {
	if x.Balance < amount {
		return fmt.Errorf("an error occurred: %s", "Insufficient Balance")
	}
	x.Balance -= amount
	zg.Info().Msg("Payment successful via Xendit")
	return nil
}

type Midtrans struct {
	Balance float64
}

func (m *Midtrans) Pay(amount float64) error {
	if m.Balance < amount {
		return fmt.Errorf("an error occurred: %s", "Insufficient Balance")
	}
	m.Balance -= amount
	zg.Info().Msg("Payment successful via Midtrans")
	return nil
}

func Checkout(payer Payer, amount float64) error {
	return payer.Pay(amount)
}
