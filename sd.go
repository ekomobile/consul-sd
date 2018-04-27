package consulsd

import (
	"errors"
	"net"
)

var NoServiceError = errors.New("no service")

type ServiceDiscovery interface {
	Get(id string) (Service, error)
}

type Service interface {
	Address() string
	Port() int
	Addr() (net.Addr, error)
	MustAddr() net.Addr
}
