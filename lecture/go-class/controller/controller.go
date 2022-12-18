package controller

import "fmt"

type Number struct {
	a int
	b int
}

func (n Number) NewCTL(a, b int) (Number, error){
	if n.a < 0 || n.b < 0 {
		return n, fmt.Errorf("invalid value %d, %d", n.a, n.b)
	}

	fmt.Println(n.a, n.b)
	return n, nil
}

func (n Number) CalcSum(a, b int) {
	fmt.Println(a + b)
}

func (n Number) CalcMul(a, b int) {
	fmt.Println(a * b)
}

func (n Number) CalcDiv(a, b int) {
	fmt.Println(a / b)
}

func (n Number) CalcSub(a, b int) {
	fmt.Println(a - b)
}