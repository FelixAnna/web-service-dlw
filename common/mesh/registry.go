package mesh

import (
	"os"

	"github.com/FelixAnna/web-service-dlw/common/aws"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	"go-micro.dev/v4/registry"
)

var (
	consulRegAddr string
	etcdRegAddr   string
	//kubernetesRegAddr string
)

func init() {
	consulRegAddr = aws.GetParameterByKey("mesh/consulRegAddr")
	etcdRegAddr = aws.GetParameterByKey("mesh/etcdRegAddr")
	//kubernetesRegAddr = aws.GetParameterByKey("mesh/kubernetesRegAddr")
}

func GetRegistry() registry.Registry {
	profile := os.Getenv("profile")
	if profile == "dev" {
		return GetConsulRegistry()
	}

	return GetKubernetesRegistry()
}

func GetConsulRegistry() registry.Registry {
	consulReg := consul.NewRegistry(registry.Addrs(consulRegAddr))

	return consulReg
}

func GetEtcdRegistry() registry.Registry {
	etcdReg := etcd.NewRegistry(registry.Addrs(etcdRegAddr))

	return etcdReg
}

func GetKubernetesRegistry() registry.Registry {
	k8sReg := kubernetes.NewRegistry()

	return k8sReg
}
