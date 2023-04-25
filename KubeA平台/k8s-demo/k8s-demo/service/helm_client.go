package service

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"helm.sh/helm/v3/pkg/action"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"log"
	"os"
)

var HelmConfig helmConfig

type helmConfig struct{
	ActionConfigMap map[string]*action.Configuration
}

func(h *helmConfig) Init() {
	mp := make(map[string]*action.Configuration, 0)
	for cluster, kubeconfig := range K8s.KubeConfMap {
		client := K8s.ClientMap[cluster]
		namespaces, err := Namespace.GetNamespaces(client, "", 0, 0)
		if err != nil {
			panic(err)
		}
		for _, namespace := range namespaces.Items {
			actionConfig := new(action.Configuration)
			cf := genericclioptions.ConfigFlags{
				KubeConfig:       &kubeconfig,
				Namespace:        &namespace.Name,
			}
			if err := actionConfig.Init(&cf, namespace.Name, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
				logger.Error("helmConfig初始化失败，%+v", err)
				panic("helmConfig初始化失败, " + err.Error())
			}
			str := fmt.Sprintf("%s-%s", namespace.Name, cluster)
			mp[str] = actionConfig
			logger.Info(fmt.Sprintf("集群:%s,命名空间:%s,初始化actionConfig成功 ", cluster, namespace.Name))
		}
	}
	h.ActionConfigMap = mp
}
//获取action配置
func(*helmConfig) GetAc(cluster, namespace string) (*action.Configuration, error) {
	kubeconfig, ok := K8s.KubeConfMap[cluster]
	if !ok {
		logger.Error("actionConfig初始化失败, cluster不存在")
		return nil, errors.New("actionConfig初始化失败, cluster不存在")
	}
	actionConfig := new(action.Configuration)
	cf := &genericclioptions.ConfigFlags{
		KubeConfig:       &kubeconfig,
		Namespace:        &namespace,
	}
	if err := actionConfig.Init(cf, namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		logger.Error("actionConfig初始化失败, %+v", err)
		return nil, errors.New("actionConfig初始化失败, " + err.Error())
	}
	return actionConfig, nil
}