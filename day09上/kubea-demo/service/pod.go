package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubea-demo/config"
)

var Pod pod

type pod struct{}
//定义列表的返回类型
type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}

//获取pod列表
//client用于选择哪个集群
func(p *pod) GetPods(client *kubernetes.Clientset, filterName, namespace string, limit, page int)(podsResp *PodsResp, err error) {
	//context.TODO()用于声明一个空的context上下文，用于List方法内设置这个请求的超时（源码），这里的常
	//用用法
	//metav1.ListOptions{}用于过滤List数据，如使用label，field等
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("获取Pod列表失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取Pod列表失败, %v\n", err))
	}
	//实例化dataSelector对象
	selectableData := &dataSelector{
		GenericDataList:   p.toCells(podList.Items),
		dataSelectorQuery: &DataSelectorQuery{
			FilterQuery:   &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	//先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	//在排序和分页
	data := filtered.Sort().Paginate()
	//将[]DataCell类型的pod列表转为v1.pod列表
	pods := p.fromCells(data.GenericDataList)

	return &PodsResp{
		Items: pods,
		Total: total,
	}, nil
}

//获取pod详情
func(p *pod) GetPodDetail(client *kubernetes.Clientset, podName, namespace string) (pod *corev1.Pod, err error) {
	pod, err = client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("获取Pod详情失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取Pod详情失败, %v\n", err))
	}
	return pod, nil
}

//删除pod
func(p *pod) DeletePod(client *kubernetes.Clientset, podName, namespace string) (err error) {
	err = client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("删除Pod失败, %v\n", err))
		return errors.New(fmt.Sprintf("删除Pod失败, %v\n", err))
	}
	return nil
}

//更新pod
//content就是pod的整个json体
func(p *pod) UpdatePod(client *kubernetes.Clientset, namespace, content string) (err error) {
	//content转成pod结构体
	var pod = &corev1.Pod{}
	//反序列化成pod对象
	err = json.Unmarshal([]byte(content), &pod)
	if err != nil {
		logger.Error(fmt.Sprintf("反序列化失败, %v\n", err))
		return errors.New(fmt.Sprintf("反序列化失败, %v\n", err))
	}
	//更新pod
	_, err = client.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("更新Pod失败, %v\n", err))
		return errors.New(fmt.Sprintf("更新Pod失败, %v\n", err))
	}
	return nil
}

//获取pod中的容器名
func(p *pod) GetPodContainer(client *kubernetes.Clientset, podName, namespace string) (containers []string, err error) {
	//获取pod详情
	pod, err := p.GetPodDetail(client, podName, namespace)
	if err != nil {
		return nil, err
	}
	//从pod对象中拿到容器名
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

//获取pod中的容器日志
func(p *pod) GetPodLog(client *kubernetes.Clientset, containerName, podName, namespace string) (log string, err error) {
	//设置日志的配置，容器名以及tail的行数
	lineLimit := int64(config.PodLogTailLine)
	option := &corev1.PodLogOptions{
		Container:                    containerName,
		TailLines:                    &lineLimit,
	}
	//获取request实例
	req := client.CoreV1().Pods(namespace).GetLogs(podName, option)
	//发起request请求，返回一个ioReadCloser类型（等同于response.body）
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logger.Error(fmt.Sprintf("获取PodLog失败, %v\n", err))
		return "", errors.New(fmt.Sprintf("获取PodLog失败, %v\n", err))
	}
	defer podLogs.Close()
	//将response body写入缓冲区，目的是为了转成string返回
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logger.Error(fmt.Sprintf("复制PodLog失败, %v\n", err))
		return "", errors.New(fmt.Sprintf("复制PodLog失败, %v\n", err))
	}
	return buf.String(), nil
}


//定义DataCell到Pod类型转换的方法
func(p *pod) toCells(std []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = podCell(std[i])
	}
	return cells
}

func(p *pod) fromCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}