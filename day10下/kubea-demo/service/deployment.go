package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"time"
)

var Deployment deployment

type deployment struct{}

type DeploymentResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

//定义Deployment创建的结构体
type DeployCreate struct {
	Name         string   `json:"name"`
	Namespace    string   `json:"namespace"`
	Replicas     int32    `json:"replicas"`
	Image        string   `json:"image"`
	Label        map[string]string `json:"label"`
	Cpu          string   `json:"cpu"`
	Memory       string   `json:"memory"`
	ContainerPort int32   `json:"container_port"`
	HealthCheck  bool     `json:"health_check"`
	HealthPath   string   `json:"health_path"`
	Cluster      string   `json:"cluster"`
}

//获取deployment列表
func(d *deployment) GetDeployments(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (deploymentResp *DeploymentResp, err error) {
	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("获取Deployment列表失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取Deployment列表失败, %v\n", err))
	}
	//实例化dataSelector对象
	selectableData := &dataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
		dataSelectorQuery: &DataSelectorQuery{
			FilterQuery:   &FilterQuery{Name:filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	//先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	//再排序和分页
	data := filtered.Sort().Paginate()

	deployments := d.fromCells(data.GenericDataList)

	return &DeploymentResp{
		Items: deployments,
		Total: total,
	}, nil
}
//获取deployment详情
func(d *deployment) GetDeploymentDetail(client *kubernetes.Clientset, deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deployment, err = client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("获取Deployment详情失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取Deployment详情失败, %v\n", err))
	}

	return deployment, nil
}
//更新deployment
func(d *deployment) UpdateDeployment(client *kubernetes.Clientset, namespace, content string) (err error) {
	var deploy = &appsv1.Deployment{}

	err = json.Unmarshal([]byte(content), deploy)
	if err != nil {
		logger.Error(fmt.Sprintf("反序列化失败, %v\n", err))
		return errors.New(fmt.Sprintf("反序列化失败, %v\n", err))
	}

	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("更新Deployment失败, %v\n", err))
		return errors.New(fmt.Sprintf("更新Deployment失败, %v\n", err))
	}
	return nil
}
//删除deployment
func(d *deployment) DeleteDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	err = client.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("删除Deployment失败, %v\n", err))
		return errors.New(fmt.Sprintf("删除Deployment失败, %v\n", err))
	}

	return nil
}

//修改Deployment副本数
func(d *deployment) ScaleDeployment(client *kubernetes.Clientset, deploymentName, namespace string, scaleNum int) (replica int32, err error) {
	//获取aotuscalingv1.Scale类型的对象，能点出当前的副本数
	scale, err := client.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("获取Deployment副本信息失败, %v\n", err))
		return 0, errors.New(fmt.Sprintf("获取Deployment副本信息失败, %v\n", err))
	}
	//修改副本数
	scale.Spec.Replicas = int32(scaleNum)
	//副本副本数，传入scale对象
	newScale, err := client.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("更新Deployment副本信息失败, %v\n", err))
		return 0, errors.New(fmt.Sprintf("更新Deployment副本信息失败, %v\n", err))
	}
	return newScale.Spec.Replicas, nil
}

//重启Deployment
func(d *deployment) RestartDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	//通过patch方法实现重启
	//此功能等同于一下kubectl命令
	//kubectl deployment ${service} -p \
	//'{"spec":{"template":{"spec":{"containers":[{"name":"'"${service}"'","env":[{"name":"RESTART_","value":"'$(date +%s)'"}]}]}}}}'

	//使用patchData Map组装数据
	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{"name": deploymentName,
							"env": []map[string]string{{
								"name":  "RESTART_",
								"value": strconv.FormatInt(time.Now().Unix(), 10),
							}},
						},
					},
				},
			},
		},
	}
	//序列化成字符串
	patchByte, err := json.Marshal(patchData)
	if err != nil {
		logger.Error(fmt.Sprintf("序列化失败, %v\n", err))
		return errors.New(fmt.Sprintf("序列化失败, %v\n", err))
	}
	//调用patch方法更新deployment
	_, err = client.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName,
		"application/strategic-merge-patch+json", patchByte, metav1.PatchOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("重启Deployment失败, %v\n", err))
		return errors.New(fmt.Sprintf("重启Deployment失败, %v\n", err))
	}
	return nil
}

//创建Deployment
func(d *deployment) CreateDeployment(client *kubernetes.Clientset, data *DeployCreate) (err error) {
	//将data中的属性组装成appsv1.Deployment对象
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:                       data.Name,
			Namespace:                  data.Namespace,
			Labels:                     data.Label,
		},
		Spec:       appsv1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels:      data.Label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: data.Name,
					Labels: data.Label,
				},
				Spec:       corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name: "http",
									Protocol: corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
		Status:     appsv1.DeploymentStatus{},
	}
	//判断是否打开健康检查功能，若打开，则定义ReadinessProbe和LivenessProbe
	if data.HealthCheck {
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler:                  corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path:        data.HealthPath,
					//intstr.IntOrString的作用是端口可以定义为整行，也可以定义为字符串
					//type=0表示整行，使用intVal
					//type=1表示字符串，使用strVal
					Port:        intstr.IntOrString{
						Type: 0,
						IntVal: data.ContainerPort,
					},
				},
			},
			//初始化等待时间
			InitialDelaySeconds:           5,
			//超时时间
			TimeoutSeconds:                15,
			//执行间隔
			PeriodSeconds:                 5,
		}
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler:                  corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path:        data.HealthPath,
					//intstr.IntOrString的作用是端口可以定义为整行，也可以定义为字符串
					//type=0表示整行，使用intVal
					//type=1表示字符串，使用strVal
					Port:        intstr.IntOrString{
						Type: 0,
						IntVal: data.ContainerPort,
					},
				},
			},
			//初始化等待时间
			InitialDelaySeconds:           15,
			//超时时间
			TimeoutSeconds:                15,
			//执行间隔
			PeriodSeconds:                 5,
		}
	}
	//定义容器的limit和request资源
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits =
	map[corev1.ResourceName]resource.Quantity {
		corev1.ResourceCPU: resource.MustParse(data.Cpu),
		corev1.ResourceMemory: resource.MustParse(data.Memory),
	}
	deployment.Spec.Template.Spec.Containers[0].Resources.Requests =
	map[corev1.ResourceName]resource.Quantity {
		corev1.ResourceCPU: resource.MustParse(data.Cpu),
		corev1.ResourceMemory: resource.MustParse(data.Memory),
	}
	//创建deployment
	_, err = client.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("创建Deployment失败, %v\n", err))
		return errors.New(fmt.Sprintf("创建Deployment失败, %v\n", err))
	}
	return nil
}

//定义DataCell到Pod类型转换的方法
func(d *deployment) toCells(std []appsv1.Deployment) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = deploymentCell(std[i])
	}
	return cells
}

func(d *deployment) fromCells(cells []DataCell) []appsv1.Deployment {
	pods := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		pods[i] = appsv1.Deployment(cells[i].(deploymentCell))
	}
	return pods
}