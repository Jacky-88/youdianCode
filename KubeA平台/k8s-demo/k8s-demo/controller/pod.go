package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-demo/service"
	"net/http"
)

var Pod pod

type pod struct {}

//Controller中的方法入参是gin.Context，用于从上下文中获取请求参数及定义响应内容
//流程：绑定参数->调用service代码->根据调用结果响应具体内容

//获取pod列表，支持过滤、排序、分页
func(p *pod) GetPods(ctx *gin.Context) {
	//匿名结构体，用于声明入参，get请求为form格式，其他请求为json格式
	params := new(struct {
		FilterName  string `form:"filter_name"`
		Namespace   string `form:"namespace"`
		Page        int    `form:"page"`
		Limit       int    `form:"limit"`
		Cluster     string `form:"cluster"`
	})
	//绑定参数，给匿名结构体中的属性赋值，值是入参
	//form格式使用ctx.Bind方法，json格式使用ctx.ShouldBindJSON方法
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		//ctx.JSON方法用于返回响应内容，入参是状态码和响应内容，响应内容放入gin.H的map中
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	//service中的的方法通过 包名.结构体变量名.方法名 使用，serivce.Pod.GetPods()
	data, err := service.Pod.GetPods(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Pod列表成功",
		"data": data,
	})
}
//获取pod详情
func(p *pod) GetPodDetail(ctx *gin.Context) {
	params := new(struct{
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
		Cluster   string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodDetail(client, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Pod详情成功",
		"data": data,
	})
}

//删除pod
func(p *pod) DeletePod(ctx *gin.Context) {
	params := new(struct{
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
		Cluster   string `json:"cluster"`
	})
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	err = service.Pod.DeletePod(client, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除Pod成功",
		"data": nil,
	})
}

//更新pod
func(p *pod) UpdatePod(ctx *gin.Context) {
	params := new(struct{
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
		Cluster   string `json:"cluster"`
	})
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	err = service.Pod.UpdatePod(client, params.PodName, params.Namespace, params.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新Pod成功",
		"data": nil,
	})
}

//获取pod容器
func(p *pod) GetPodContainer(ctx *gin.Context) {
	params := new(struct{
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
		Cluster   string `form:"cluster"`
	})
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodContainer(client, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Pod容器成功",
		"data": data,
	})
}

//获取pod中容器日志
func(p *pod) GetPodLog(ctx *gin.Context) {
	params := new(struct{
		ContainerName string `form:"container_name"`
		PodName       string `form:"pod_name"`
		Namespace     string `form:"namespace"`
		Cluster       string `form:"cluster"`
	})
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodLog(client, params.ContainerName, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取Pod中容器日志成功",
		"data": data,
	})
}

//获取每个namespace的pod数量
func(p *pod) GetPodNumPerNp(ctx *gin.Context) {
	params := new(struct{
		Cluster       string `form:"cluster"`
	})
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPodNumPerNp(client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取每个namespace的pod数量成功",
		"data": data,
	})
}