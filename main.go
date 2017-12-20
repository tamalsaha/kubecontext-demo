package main

import (
	"fmt"
	"path/filepath"

	"github.com/appscode/go/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube/config")

	config, err := BuildConfigFromContext(kubeconfigPath, "minikube")
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	kc := kubernetes.NewForConfigOrDie(config)
	nodes, err := kc.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, node := range nodes.Items {
		fmt.Println(node.Name)
	}
}

func BuildConfigFromContext(kubeconfigPath, contextName string) (*restclient.Config, error) {
	loader := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath}
	overrides := &clientcmd.ConfigOverrides{
		CurrentContext:  contextName,
		ClusterDefaults: clientcmd.ClusterDefaults,
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides).ClientConfig()
}
