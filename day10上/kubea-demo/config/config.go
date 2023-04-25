package config

import "time"

const (
	ListenAddr = "0.0.0.0:9090"
	Kubeconfigs = `{"TST-1":"C:\\custom\\project\\go\\config2","TST-2":"C:\\custom\\project\\go\\config2"}`
	//查看容器日志时，显示的tail行数 tail -n 5000
	PodLogTailLine = 5000
	//数据库配置
	DbType = "mysql"
	DbHost = "localhost"
	DbPort = 3306
	DbName = "kubea_demo"
	DbUser = "root"
	DbPwd = "123456"
	//打印mysql debug sql日志
	LogMode = false
	//连接池配置
	MaxIdleConns = 10 //最大空闲连接
	MaxOpenConns = 100 //最大连接数
	MaxLifeTime = 30 * time.Second //最大生存时间
	//helm上传路径
	UploadPath = "C:\\custom\\project\\123\\实战\\"
)
