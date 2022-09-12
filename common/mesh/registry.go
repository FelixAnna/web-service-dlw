package mesh

import (
	"os"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	"go-micro.dev/v4/registry"
)

type Registry struct {
	consulRegAddr string
}

func ProvideRegistry(awsService *aws.AWSService) *Registry {
	consulRegAddr := awsService.GetParameterByKey("mesh/consulRegAddr")

	return &Registry{consulRegAddr: consulRegAddr}
}

func (service *Registry) GetRegistry() registry.Registry {
	profile := os.Getenv("profile")
	if profile == "dev" || profile == "prod" {
		return service.getConsulRegistry()
	}

	return service.getKubernetesRegistry()
}

func (service *Registry) getConsulRegistry() registry.Registry {
	consulReg := consul.NewRegistry(registry.Addrs(service.consulRegAddr))

	return consulReg
}

/*func GetEtcdRegistry() registry.Registry {
	etcdReg := etcd.NewRegistry(registry.Addrs(etcdRegAddr))

	return etcdReg
}*/

func (service *Registry) getKubernetesRegistry() registry.Registry {
	k8sReg := kubernetes.NewRegistry()

	return k8sReg
}
