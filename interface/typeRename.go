package main

import "fmt"

type ServiceType string

const (
	ServiceTypeCluster ServiceType = "ClusterIP"

	ServiceTypeNodePort ServiceType = "NodePort"

	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"

	ServiceTypeExternalName ServiceType = "ExternalName"
)

func main() {
	// ClusterIP NodePort LoadBalancer ExternalName
	fmt.Println(
		ServiceTypeCluster,
		ServiceTypeNodePort,
		ServiceTypeLoadBalancer,
		ServiceTypeExternalName,
	)
}