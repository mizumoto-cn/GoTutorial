// interface is a set of methods
// for example
// go back to Student and Salaryman in sector 4
// Student : Poopoo() Sayhello() GotoSchool()
// Salaryman : Poopoo() Sayhello() GotoWork()
// So the interfaces are Poopoo() and Sayhello()
package main

import "fmt"

type Pen struct {
	brand    string
	haveCap  bool
	inkColor string
}

type InkPen struct {
	Pen
	inkVolume float32
}

type BallPen struct {
	Pen
	refillVolume float32
}

func (p *Pen) Write(sentence string) {
	fmt.Println("Write ", sentence, "to notebook.")
}

func (p *Pen) Brand() string {
	return p.brand
}

func (p *Pen) HaveCap() bool {
	return p.haveCap
}

func (p *Pen) InkState() {
	fmt.Println("Ink is ", p.inkColor)
}

func (ip *InkPen) Write(sentence string) {
	if ip.inkVolume >= 10 {
		ip.inkVolume -= 10
		fmt.Println("Write ", sentence, "to notebook with ", ip.inkColor, " ink pen.")
	} else {
		ip.inkVolume = 0
		fmt.Println("Wrote something to notebook with ", ip.inkColor, " ink pen. but INK is not enough.")
	}
}

func (ip *InkPen) InkState() {

	fmt.Println("Ink is ", ip.inkColor, " at ", ip.inkVolume)

}

func (ip *InkPen) FillInk(color string) {
	ip.inkVolume = 100
	ip.inkColor = color
	fmt.Println("Ink filled to 100 with color of", color)
}

func (bp *BallPen) Write(sentence string) {
	if bp.refillVolume >= 10 {
		bp.refillVolume -= 10
		fmt.Println("Write ", sentence, "to notebook with ", bp.inkColor, " ball pen.")
	} else {
		bp.refillVolume = 0
		fmt.Println("Wrote something to notebook with ", bp.inkColor, " ball pen. but refill is used up.")
	}
}

func (bp *BallPen) InkState() {

	fmt.Println("Refill Ink is color of", bp.inkColor, bp.refillVolume, " left.")

}

func (bp *BallPen) ChangeRefill(color string) {
	bp.refillVolume = 100
	bp.inkColor = color
	fmt.Println("Changed to new ", color, "refill.")
}

type IPen interface {
	Write(sentence string)
	Brand()
	InkState()
	HaveCap()
}

type IInkPen interface {
	Write(sentence string)
	InkState()
	FillInk(color string)
}

type IBallPen interface {
	Write(sentence string)
	InkState()
	ChangeRefill(color string)
}

func interfacedemo() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	a = s
	fmt.Println(a)
}
