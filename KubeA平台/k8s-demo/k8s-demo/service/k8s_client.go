package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-demo/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientMap map[string]*kubernetes.Clientset
	KubeConfMap map[string]string
}


func(k *k8s) GetClient(cluster string) (*kubernetes.Clientset, error) {
	client, ok := k.ClientMap[cluster]
	if !ok {
		return nil, errors.New(fmt.Sprintf("集群:%s 不存在, 无法获取client", cluster))
	}
	return client, nil
}

func(k *k8s) Init() {
	mp := map[string]string{}
	k.ClientMap = map[string]*kubernetes.Clientset{}

	if err := json.Unmarshal([]byte(config.Kubeconfig), &mp); err != nil {
		panic(fmt.Sprintf("Kubeconfigs反序列化失败 %v\n", err))
	}
	k.KubeConfMap = mp
	for key, value := range mp {
		conf, err := clientcmd.BuildConfigFromFlags("", value)
		if err != nil {
			panic(fmt.Sprintf("集群%s: 创建K8s 配置失败 %v\n", key, err))
		}
		clientSet, err := kubernetes.NewForConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("集群%s: 创建K8sClient失败 %v\n", key, err))
		}

		k.ClientMap[key] = clientSet
		logger.Info(fmt.Sprintf("集群%s: 创建K8sClient成功 ", key))
	}
}