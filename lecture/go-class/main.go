package main

import (
	"fmt"
	"lecture/go-class/controller"
	"lecture/go-class/model"	
)


func main() {
	ctl, err := controller.NewCTL()
	
	fmt.Println(ctl.CalcSum(2, 3))
	fmt.Println(ctl.CalcMul(2, 3)) 
	fmt.Println(ctl.CalcDiv(2, 3))
	fmt.Println(ctl.CalcSub(2, 3))

	mod, err := model.NewModel()

	fmt.Println(mod.Run("run"))
	fmt.Println(mod.Jump("jump"))
	fmt.Println(mod.Sleep("sleep"))
	fmt.Println(mod.Walk("walk"))
	fmt.Println(mod.Fly("fly"))
}