// Strategy pattern is a behavioral design pattern that
// enables an algorithm to be swapped for another at runtime.
package strategy

import (
	"fmt"
)

type Strategy interface {
	Eat(*foodCtx)
}

type foodCtx struct {
	name  string
	price int
}

type Lunch struct {
	foodCtx  *foodCtx
	strategy Strategy
}

func NewLunch(foodCtx *foodCtx, strategy Strategy) *Lunch {
	return &Lunch{
		foodCtx:  foodCtx,
		strategy: strategy,
	}
}

func (l *Lunch) Eat() {
	l.strategy.Eat(l.foodCtx)
}

type Pizza struct{}

func (p *Pizza) Eat(foodCtx *foodCtx) {
	fmt.Printf("%s eats a pizza with price %d\n", foodCtx.name, foodCtx.price)
}

type Steak struct{}

func (s *Steak) Eat(foodCtx *foodCtx) {
	fmt.Printf("%s eats a steak with price %d\n", foodCtx.name, foodCtx.price)
}
