package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// DeleteTarget 想要删除的 CR 的信息
type DeleteTarget struct {
	Namespace  string
	CRName     string
	CRBaseInfo schema.GroupVersionResource
}

func (t *DeleteTarget) delete(clientset dynamic.Interface) error {
	return clientset.Resource(t.CRBaseInfo).Namespace(t.Namespace).Delete(context.TODO(), t.CRName, metav1.DeleteOptions{})
}

// DeleteCR 删除一个 CR 对象
func (t *DeleteTarget) DeleteCR(config *rest.Config) {
	clientset, _ := dynamic.NewForConfig(config)
	if err := t.delete(clientset); err != nil {
		fmt.Printf("namespace:%v\nerror:%v\n", t.Namespace, err)
	}
}

// ParseFlags 解析命令行标志
func (t *DeleteTarget) ParseFlags() {
	t.Namespace = *flag.String("ns", "default", "指定名称空间")
	t.CRName = *flag.String("name", "rabbitmq", "指定 rabbitmqcluster 对象的名称")
	t.CRBaseInfo.Group = *flag.String("crgroup", "rabbitmq.com", "指定 CR 的 Group")
	t.CRBaseInfo.Version = *flag.String("crversion", "v1beta1", "指定 CR 的 Version")
	t.CRBaseInfo.Resource = *flag.String("crname", "rabbitmqclusters", "指定 CR 的名称")
	flag.Parse()
}

func main() {
	t := new(DeleteTarget)
	t.ParseFlags()
	// fmt.Println(t.CRBaseInfo.Group)
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	// fmt.Println(reflect.TypeOf(namespace))
	t.DeleteCR(config)
}
