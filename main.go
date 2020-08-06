package main

import "fmt"

type CalcInterface interface {
	add() int
	sub() int
	mult(int) (int, int)
}

type CalcStruct struct {
	a int
	b int
}

func (c *CalcStruct) add() int {
	return c.a + c.b
}

func (c *CalcStruct) sub() int {
	return c.a - c.b
}

func (c *CalcStruct) mult(p int) (int, int) {
	return c.a * p, c.b * p
}

func main() {
	var calcObj CalcInterface
	calcObj = &CalcStruct{2, 3}
	sum := calcObj.add()
	diff := calcObj.sub()
	a, b := calcObj.mult(2)
	fmt.Println("Sum : ", sum, "Diff : ", diff, "Mult : ", a, b)
}
