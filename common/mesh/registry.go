package mesh

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4/registry"
)

var consulRegAddr string

func init() {
	consulRegAddr = aws.GetParameterByKey("mesh/consulRegAddr")
}

func GetConsulRegistry() registry.Registry {
	consulReg := consul.NewRegistry(registry.Addrs(consulRegAddr))

	return consulReg
}
