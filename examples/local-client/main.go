package main

import (
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig = flag.String("kubeconfig", "/Users/nbenaglia/.kube/config", "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)

	pod, err := clientset.CoreV1().Pods("my-namespace").Get("my-pod", metav1.GetOptions{})

	fmt.Println(pod.GetName())
}
