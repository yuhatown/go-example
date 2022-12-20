package controller

import "go-example/lecture/go-loger/model"

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md : rep}
	return r, nil
}