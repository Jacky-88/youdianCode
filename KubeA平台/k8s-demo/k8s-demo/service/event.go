package service

import (
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-demo/dao"
	"k8s-demo/model"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"time"
)

var Event event

type event struct{}

//获取列表
func(*event) GetList(name, cluster string, page, limit int) (*dao.Events, error) {
	data, err := dao.Event.GetList(name, cluster, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}


//informer 监听event
func (*event) WatchEventTask(cluster string) {
	informerFactory := informers.NewSharedInformerFactory(K8s.ClientMap[cluster], time.Minute)
	informer := informerFactory.Core().V1().Events()
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    func(obj interface{}){
				onAdd(obj, cluster)
			},
			//UpdateFunc: onUpdate,
			//DeleteFunc: onDelete,
		},
	)
	stopCh := make(chan struct{})
	defer close(stopCh)
	informerFactory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced) {
		logger.Error("Timed out waiting for caches to sync")
		return
	}
	<-stopCh

	return
}
//新增时落库
func onAdd(obj interface{}, cluster string) {
	event := obj.(*v1.Event)
	_, has, err := dao.Event.HasEvent(event.InvolvedObject.Name,
		event.InvolvedObject.Kind,
		event.InvolvedObject.Namespace,
		event.Reason,
		event.CreationTimestamp.Time,
		cluster,
	)
	if err != nil {
		return
	}
	if has {
		logger.Error(fmt.Sprintf("Event数据已存在, %s %s %s %s %v %s\n",
			event.InvolvedObject.Name,
			event.InvolvedObject.Kind,
			event.InvolvedObject.Namespace,
			event.Reason,
			event.CreationTimestamp.Time,
			cluster),
		)
		return
	}
	data := &model.Event{
		Name: event.InvolvedObject.Name,
		Kind: event.InvolvedObject.Kind,
		Namespace: event.InvolvedObject.Namespace,
		Rtype: event.Type,
		Reason: event.Reason,
		Message: event.Message,
		EventTime: &event.CreationTimestamp.Time,
		Cluster: cluster,
	}
	if err := dao.Event.Add(data); err != nil {
		return
	}
}