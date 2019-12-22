package controllers

import (
	"fmt"
	"github.com/lucasmbaia/punctorum/api/models/interfaces"
)

type Resources struct {
	GetModel func() interfaces.Models
	GetFields func() interface{}
}

func (r *Resources) Get(c interface{}) {
	fmt.Println("PORRA")
	fmt.Println(r.GetModel())
	r.GetModel().Get(c)
	return
}

func (r *Resources) Post(c interface{}) {
	return
}
