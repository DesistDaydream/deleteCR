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
	ObjectName string
	CRBaseInfo schema.GroupVersionResource
}

// NewDeleteTarget 实例化 DeleteTarget
func NewDeleteTarget() *DeleteTarget {
	return &DeleteTarget{
		Namespace:  "rabbitmq",
		ObjectName: "rabbitmq-bj-test",
		CRBaseInfo: schema.GroupVersionResource{
			Group:    "rabbitmq.com",
			Version:  "v1beta1",
			Resource: "rabbitmqclusters",
		},
	}
}

func (t *DeleteTarget) delete(clientset dynamic.Interface) error {
	return clientset.Resource(t.CRBaseInfo).Namespace(t.Namespace).Delete(context.TODO(), t.ObjectName, metav1.DeleteOptions{})
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
	flag.StringVar(&t.Namespace, "ns", t.Namespace, "指定名称空间")
	flag.StringVar(&t.ObjectName, "name", t.ObjectName, "指定 rabbitmqcluster 对象的名称")
	flag.StringVar(&t.CRBaseInfo.Group, "crgroup", t.CRBaseInfo.Group, "指定 CR 的 Group")
	flag.StringVar(&t.CRBaseInfo.Version, "crversion", t.CRBaseInfo.Version, "指定 CR 的 Version")
	flag.StringVar(&t.CRBaseInfo.Resource, "crname", t.CRBaseInfo.Resource, "指定 CR 的名称")
	flag.Parse()
}

func main() {
	t := NewDeleteTarget()
	t.ParseFlags()
	fmt.Printf("名称空间：%v\n对象名：%v\nCR组：%v\nCR版本：%v\nCR名：%v\n", t.Namespace, t.ObjectName, t.CRBaseInfo.Group, t.CRBaseInfo.Version, t.CRBaseInfo.Resource)
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	// fmt.Println(reflect.TypeOf(namespace))
	t.DeleteCR(config)
}
