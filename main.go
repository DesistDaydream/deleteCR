package main

import (
	"context"
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

var crBaseInfo = schema.GroupVersionResource{
	Group:    "rabbitmq.com",
	Version:  "v1beta1",
	Resource: "rabbitmqclusters",
}

func delete(clientset dynamic.Interface, namespace string, name string) error {
	return clientset.Resource(crBaseInfo).Namespace(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

// DeleteCR 删除一个 CR 对象
func DeleteCR(config *rest.Config, ns string, name string) {
	clientset, _ := dynamic.NewForConfig(config)
	if err := delete(clientset, ns, name); err != nil {
		fmt.Printf("namespace:%v\nerror:%v\n", ns, err)
	}
}

func main() {
	namespace := flag.String("ns", "default", "指定名称空间")
	objectName := flag.String("name", "rabbitmq", "指定 rabbitmqcluster 对象的名称")
	flag.Parse()
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	// fmt.Println(reflect.TypeOf(namespace))
	DeleteCR(config, *namespace, *objectName)
}
