package broker

import (
	"github.com/lucasmbaia/punctorum/api/repository/filter"
)

type Brokers interface {
	Create(interface{}) error
	Read([]filter.Filters, interface{}, ...interface{}) error
	Delete(interface{}) (bool, error)
}
