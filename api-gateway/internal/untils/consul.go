package untils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Consul struct {
	consulClient *api.Client
}

type ConsulRegister struct {
	Id      string
	Name    string
	Tags    []string
	Port    int
	Address string
}

func NewConsulClient(host string, port int) (*Consul, error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Consul{client}, nil
}
func (c *Consul) Register(register ConsulRegister) error {
	registration := api.AgentServiceRegistration{
		ID:      register.Id,
		Name:    register.Name,
		Tags:    register.Tags,
		Port:    register.Port,
		Address: register.Address,
	}

	return c.consulClient.Agent().ServiceRegister(&registration)
}
