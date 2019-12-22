package decorators

import (
	"fmt"
	"github.com/lucasmbaia/punctorum/api/models/interfaces"
)

type Transaction struct {
	interfaces.Models
}

func NewTransaction(m interfaces.Models) interfaces.Models {
	return &Transaction{m}
}

func (t *Transaction) Get(data interface{}) (response interface{}, err error) {
	fmt.Println("Transaction")
	t.Models.Get(data)
	return
}
