package register

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"log"
)
func init (){
	consulserver := NewRegister()
	consulserver.Register()
}
type ConsulServer struct {
	Client consul.Client
}
func NewRegister() *ConsulServer {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	apiserver , err := api.NewClient(config)
	if err != nil {
		log.Println(err)
		return nil
	}
	client := consul.NewClient(apiserver)
	return &ConsulServer{Client: client}
}

func (c ConsulServer) Register () {
	err :=  c.Client.Register(&api.AgentServiceRegistration{
		Port: 8085,
		ID: "test",
		Address: "127.0.0.1",
		Name: "testname",
		Check: &api.AgentServiceCheck{GRPC: "127.0.0.1:8500/health",
			Interval: "5s",
			Timeout: "5s",
			},
	})
	if err != nil {
		log.Println(err)
	}
}