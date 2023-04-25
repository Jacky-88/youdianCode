package config

import "time"

const (
	//监听地址
	ListenAddr = "0.0.0.0:9090"
	WsAddr = "0.0.0.0:8081"
	//kubeconfig路径
	Kubeconfig = `{"TST-1":"D:\\4.data\\golang\\k8s\\config"}`  //windows路径
	//linux 路径
	//Kubeconfigs = `{"TST-1":"/Users/adoo/.kube/config","TST-2":"/Users/adoo/.kube/config"}`
	//pod日志tail显示行数
	PodLogTailLine = 2000
	//登录账号密码
	AdminUser = "admin"
	AdminPwd = "123456"

	//数据库配置
	DbType = "mysql"
	DbHost = "192.168.31.55"
	DbPort = 3306
	DbName = "k8s_demo"
	DbUser = "root"
	DbPwd = "luyijian"
	//打印mysql debug sql日志
	LogMode = false
	//连接池配置
	MaxIdleConns = 10 //最大空闲连接
	MaxOpenConns = 100 //最大连接数
	MaxLifeTime = 30 * time.Second //最大生存时间
	//helm配置
	UploadPath = "/Users/adoo/chart"
)
