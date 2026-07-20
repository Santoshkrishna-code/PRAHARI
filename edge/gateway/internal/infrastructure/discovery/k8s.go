package discovery

import (
	"fmt"
)

// K8sResolver resolves service names dynamically to Kubernetes DNS records.
type K8sResolver struct {
	namespace string
}

// NewK8sResolver constructs a K8sResolver.
func NewK8sResolver(namespace string) *K8sResolver {
	if namespace == "" {
		namespace = "default"
	}
	return &K8sResolver{namespace: namespace}
}

// Resolve DNS names matching cluster services.
func (r *K8sResolver) Resolve(serviceName string) ([]string, error) {
	// E.g. incident-service -> []string{"incident-service.default.svc.cluster.local:8082"}
	dnsName := fmt.Sprintf("%s.%s.svc.cluster.local", serviceName, r.namespace)
	return []string{dnsName}, nil
}
