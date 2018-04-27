package consulsd

import (
	"net"

	"github.com/hashicorp/consul/api"
)

type serviceDiscovery struct {
	consul *api.Client
}

func NewServiceDiscovery() (ServiceDiscovery, error) {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return &serviceDiscovery{
		consul: c,
	}, nil
}

func (sd *serviceDiscovery) Get(id string) (Service, error) {
	ss, _, err := sd.consul.Catalog().Service(id, "", nil)
	if err != nil {
		return nil, err
	}

	if len(ss) == 0 {
		return nil, NoServiceError
	}

	s := ss[0]

	address := s.ServiceAddress
	port := s.ServicePort

	if address == "" {
		address = s.Address
	}

	return &service{
		address: address,
		port:    port,
	}, nil
}

type service struct {
	address string
	port    int
}

func (s *service) Addr() (net.Addr, error) {
	ip, err := net.ResolveIPAddr("ip", s.address)
	if err != nil {
		return nil, err
	}

	return &net.TCPAddr{
		IP:   ip.IP,
		Port: s.port,
	}, nil
}

func (s *service) MustAddr() net.Addr {
	addr, err := s.Addr()
	if err != nil {
		panic(err)
	}

	return addr
}

func (s *service) Address() string {
	return s.address
}

func (s *service) Port() int {
	return s.port
}
