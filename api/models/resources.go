package models

import (
	"reflect"
	"fmt"

	"github.com/lucasmbaia/punctorum/api/models/interfaces"
	"github.com/lucasmbaia/punctorum/api/models/decorators"
)

type Resources struct {
	Model	reflect.Value
}

func NewResources(m interface{}) interfaces.Models {
	var model = reflect.ValueOf(m)

	return decorators.NewTransaction(&Resources{model})
}

func (r *Resources) Get(data interface{}) (response interface{}, err error) {
	var args []reflect.Value

	args = append(args, reflect.ValueOf(data))
	r.Model.MethodByName("Get").Call(args)

	fmt.Println("TA AQUI")
	return
}

func (r *Resources) Post(data interface{}) (async bool, err error) {
	return
}
